package ordercontroller

import (
	"encoding/json"
	"net/http"
	"online-store/helper"
	"online-store/models"
	"strconv"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {

	var orders []models.Order
	err := models.GetAllOrders(&orders)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	response := map[string]interface{}{
		"status":  "success",
		"message": "success",
		"data":    orders,
	}
	helper.ResponseJSON(w, http.StatusOK, response)

}

func Create(w http.ResponseWriter, r *http.Request) {
	var order models.Order
	decode := json.NewDecoder(r.Body)
	if err := decode.Decode(&order); err != nil {
		response := map[string]interface{}{
			"status":  "fail",
			"message": "JSON tidak valid",
		}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
	}
	err := helper.ValidateStruct(&order)
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

	err = order.CreateOrder()
	if err != nil {
		response := map[string]interface{}{
			"status":  "fail",
			"message": "gagal menambahkan order",
		}
		helper.ResponseJSON(w, http.StatusOK, response)
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
	}

	var order models.Order
	err = models.GetOrderByID(uint(id), &order)
	if err != nil {
		response := map[string]interface{}{
			"status":  "fail",
			"message": "produk tidak ditemukan",
		}
		helper.ResponseJSON(w, http.StatusNotFound, response)
		return
	}
	helper.ResponseJSON(w, http.StatusOK, order)
}
func UpdateStatus(w http.ResponseWriter, r *http.Request) {
	// Ambil parameter order ID dari URL
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

	var updatedData struct {
		Status string `json:"status" validate:"required"`
	}

	if err := json.NewDecoder(r.Body).Decode(&updatedData); err != nil {
		response := map[string]interface{}{
			"status":  "fail",
			"message": "JSON tidak valid",
		}
		helper.ResponseJSON(w, http.StatusBadGateway, response)
		return
	}
	err = helper.ValidateStruct(&updatedData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		response := map[string]interface{}{
			"status":  "fail",
			"message": errors,
		}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	var order models.Order
	if err := models.DB.First(&order, id).Error; err != nil {
		response := map[string]interface{}{
			"status":  "fail",
			"message": "order tidak ditemukan",
		}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	order.Status = updatedData.Status
	if err := models.UpdateOrder(&order); err != nil {
		response := map[string]interface{}{
			"status":  "fail",
			"message": "gagal mengubah status",
		}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := map[string]interface{}{
		"status":  "success",
		"message": "berhasil mengubah status menjadi " + order.Status,
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

	err = models.DeleteOrder(uint(id))
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
