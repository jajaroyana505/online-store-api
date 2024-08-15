package authcontroller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"online-store/config"
	"online-store/helper"
	"online-store/models"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var userInput models.Users

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userInput); err != nil {

		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}
	defer r.Body.Close()

	// ambil data user data database
	var user models.Users

	if err := models.DB.Where("username = ?", userInput.Username).First(&user).Error; err != nil {
		response := map[string]string{"mesage": err.Error()}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return

	}

	// cek apakah password valid?
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInput.Password))
	if err != nil {
		response := map[string]string{"mesage": err.Error()}
		helper.ResponseJSON(w, http.StatusUnauthorized, response)
		return
	}

	// proses pembuatan token jwt

	expTime := time.Now().Add(time.Minute * 1)
	claims := &config.JWTClaim{
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "online-stor-app",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	// mendeklarasikan algoritma yg akan digunakan untuk sigin

	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := tokenAlgo.SignedString(config.JWT_KEY)
	if err != nil {
		response := map[string]string{"mesage": err.Error()}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}
	// set token ke cookie

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    token,
		HttpOnly: true,
	})

	response := map[string]string{"mesage": "login success"}
	helper.ResponseJSON(w, http.StatusOK, response)

}

func ChangePassword(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		response := map[string]interface{}{
			"status":  "fail",
			"message": "ID tidak valid",
		}
		helper.ResponseJSON(w, http.StatusBadGateway, response)
		return
	}

	newPasswordCh := make(chan string)
	var userInput struct {
		CurrentPassword string `json:"current_password"`
		NewPassword     string `json:"new_password"`
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userInput); err != nil {
		response := map[string]string{"message": "Json tidak valid"}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close()

	var user models.Users
	result := models.DB.First(&user, id)
	if result.Error != nil {
		response := map[string]string{"message": "failed get data from database"}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	// cek apakah password valid?
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInput.CurrentPassword))
	if err != nil {
		fmt.Println(err.Error())
		response := map[string]string{"message": "Current password incorrect"}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	go func() {
		hashPassword, _ := bcrypt.GenerateFromPassword([]byte(userInput.NewPassword), bcrypt.DefaultCost)
		newPasswordCh <- string(hashPassword)
	}()

	user.Password = <-newPasswordCh
	err = models.DB.Save(user).Error
	if err != nil {
		response := map[string]string{"message": "failed change password"}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := map[string]string{"message": "Success change password"}
	helper.ResponseJSON(w, http.StatusAccepted, response)

}

func Register(w http.ResponseWriter, r *http.Request) {

	var userInput models.Users
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userInput); err != nil {
		response := map[string]string{"message": "fail"}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	defer r.Body.Close()

	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
	userInput.Password = string(hashPassword)

	if userInput.CreateUser(models.DB) == 0 {
		response := map[string]string{"message": "username exist"}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	response := map[string]string{"message": "Success create new user"}
	helper.ResponseJSON(w, http.StatusOK, response)
}

func Logout(w http.ResponseWriter, r *http.Request) {

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    "",
		HttpOnly: true,
		MaxAge:   -1,
	})

	response := map[string]string{"message": "Logout"}

	helper.ResponseJSON(w, http.StatusOK, response)

}
