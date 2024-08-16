package helper

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

// Fungsi untuk memformat pesan kesalahan validasi
func FormatValidationError(err error) map[string]string {
	errors := make(map[string]string)
	for _, err := range err.(validator.ValidationErrors) {
		var errorMsg string
		field := err.Field()

		switch err.Tag() {
		case "required":
			errorMsg = fmt.Sprintf("%s is required", field)
		default:
			errorMsg = fmt.Sprintf("%s is invalid", field)
		}

		errors[field] = errorMsg
	}
	return errors
}

// Fungsi untuk melakukan validasi struct
func ValidateStruct(s interface{}) error {
	return validate.Struct(s)
}
