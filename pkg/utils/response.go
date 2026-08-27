package utils

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
	Error   string            `json:"error"`
	Message string            `json:"message,omitempty"`
	Details map[string]string `json:"details,omitempty"`
}

func RespondWithJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

func RespondWithError(w http.ResponseWriter, status int, message string) {
	RespondWithJSON(w, status, ErrorResponse{
		Error:   http.StatusText(status),
		Message: message,
	})
}

func RespondWithValidationError(w http.ResponseWriter, err error) {
	details := make(map[string]string)

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, ve := range validationErrors {
			details[ve.Field()] = getValidationMessage(ve)
		}
	}

	RespondWithJSON(w, http.StatusBadRequest, ErrorResponse{
		Error:   "Validation failed",
		Details: details,
	})
}

func getValidationMessage(ve validator.FieldError) string {
	switch ve.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email format"
	case "min":
		return "Value is too short"
	case "max":
		return "Value is too long"
	default:
		return ve.Error()
	}
}
