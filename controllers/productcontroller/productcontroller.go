package productcontroller

import (
	"encoding/json"
	"net/http"
	"online-store/helper"
	"online-store/models"
	"strconv"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	product_ch := make(chan []models.Product)

	go func() {
		var products []models.Product
		err := models.GetAllProducts(&products)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		product_ch <- products

	}()
	response := map[string]interface{}{
		"status":  "success",
		"message": "success",
		"data":    <-product_ch,
	}
	helper.ResponseJSON(w, http.StatusOK, response)

}

func Create(w http.ResponseWriter, r *http.Request) {

	var product models.Product
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&product); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := helper.ValidateStruct(&product)
	if err != nil {
		errors := helper.FormatValidationError(err)
		response := map[string]interface{}{
			"status":  "fail",
			"message": errors,
		}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	defer r.Body.Close()

	// inser ke db
	err = product.CreateProduct()
	if err != nil {
		response := map[string]interface{}{
			"status":  "fail",
			"message": "gagal menambahkan produk",
		}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := map[string]interface{}{
		"status":  "success",
		"message": "data berhasil disimpan",
	}
	helper.ResponseJSON(w, http.StatusOK, response)

}
func Show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		response := map[string]interface{}{
			"status":  "fail",
			"message": "ID tidak valid",
		}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	var product models.Product
	err = models.GetProductByID(uint(id), &product)
	if err != nil {
		response := map[string]interface{}{
			"status":  "fail",
			"message": "produk tidak ditemukan",
		}
		helper.ResponseJSON(w, http.StatusNotFound, response)
		return
	}
	response := map[string]interface{}{
		"status":  "success",
		"message": "success",
		"data":    product,
	}
	helper.ResponseJSON(w, http.StatusOK, response)
}

func Update(w http.ResponseWriter, r *http.Request) {
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
	var product models.Product

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&product); err != nil {
		response := map[string]interface{}{
			"status":  "fail",
			"message": "format JSON tidak valid",
		}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	err = helper.ValidateStruct(&product)
	if err != nil {
		errors := helper.FormatValidationError(err)
		response := map[string]interface{}{
			"status":  "fail",
			"message": errors,
		}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	defer r.Body.Close()

	err = models.UpdateProduct(int64(id), &product)
	if err != nil {
		response := map[string]interface{}{
			"status":  "fail",
			"message": "gagal mengubah data",
		}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}
	response := map[string]interface{}{
		"status":  "success",
		"message": "data berhasil diubah",
	}
	helper.ResponseJSON(w, http.StatusOK, response)

}
func Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		response := map[string]interface{}{
			"status":  "fail",
			"message": "data gagal dihapus",
		}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	err = models.DeleteProduct(uint(id))
	if err != nil {
		response := map[string]interface{}{
			"status":  "fail",
			"message": "data gagal dihapus",
		}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := map[string]interface{}{
		"status":  "success",
		"message": "data berhasil dihapus",
	}
	helper.ResponseJSON(w, http.StatusOK, response)
}
