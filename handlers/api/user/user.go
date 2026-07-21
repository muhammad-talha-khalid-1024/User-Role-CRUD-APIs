package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"mustaqel/config"
	"mustaqel/models"
	"mustaqel/utils/password"
)

type UserRequest struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
	Status    int16
}

func GetUsers(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var users []models.User

	if err := config.DB.Find(&users).Error; err != nil {
		http.Error(w, "", http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]any{
			"message": err.Error(),
			"status":  http.StatusNotFound,
		})
		return
	}

	json.NewEncoder(w).Encode(map[string]any{
		"message": "User fetched successfully",
		"users":   users,
		"status":  200,
	})
}

func GetUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	id := r.PathValue("id")

	var user models.User

	if err := config.DB.First(&user, id).Error; err != nil {
		http.Error(w, "", http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]any{
			"message": err.Error(),
			"status":  http.StatusNotFound,
		})
		return
	}

	json.NewEncoder(w).Encode(map[string]any{
		"message": "User fetched successfully",
		"user":    user,
		"status":  200,
	})
}

func CreateUser(w http.ResponseWriter, r *http.Request) {

	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "", http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]any{
			"message": err.Error(),
			"status":  http.StatusBadRequest,
		})
		return
	}

	pswd, err := password.HashPassword(user.Password)
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]any{
			"message": err.Error(),
			"status":  http.StatusBadRequest,
		})
		return
	}
	user.Password = pswd

	if err := config.DB.Create(&user).Error; err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]any{
			"message": err.Error(),
			"status":  http.StatusInternalServerError,
		})
		return
	}

	json.NewEncoder(w).Encode(map[string]any{
		"message": "User created successfully",
		"user":    user,
		"status":  200,
	})
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

	id := r.PathValue("id")

	var user models.User

	if err := config.DB.First(&user, id).Error; err != nil {
		http.Error(w, "", http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]any{
			"message": err.Error(),
			"status":  http.StatusNotFound,
		})
		return
	}

	var input UserRequest

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "", http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]any{
			"message": err.Error(),
			"status":  http.StatusBadRequest,
		})
		return
	}

	if input.Password != "" {
		pswd, err := password.HashPassword(input.Password)
		if err != nil {
			http.Error(w, "", http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]any{
				"message": err.Error(),
				"status":  http.StatusBadRequest,
			})
			return
		}
		input.Password = pswd
	}

	if err := config.DB.Model(&user).Updates(input).Error; err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]any{
			"message": err.Error(),
			"status":  http.StatusInternalServerError,
		})
		return
	}

	json.NewEncoder(w).Encode(map[string]any{
		"message": "User updated successfully",
		"user":    user,
		"status":  200,
	})
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	id := r.PathValue("id")

	var user models.User

	if err := config.DB.First(&user, id).Error; err != nil {
		http.Error(w, "", http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]any{
			"message": err.Error(),
			"status":  http.StatusNotFound,
		})
		return
	}

	userId := fmt.Sprintf("id=%d", user.ID)

	if err := config.DB.Delete(&user, userId).Error; err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]any{
			"message": err.Error(),
			"status":  http.StatusInternalServerError,
		})
		return
	}

	json.NewEncoder(w).Encode(map[string]any{
		"message": "User deleted successfully",
		"status":  200,
	})

}
