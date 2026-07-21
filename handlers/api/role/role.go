package role

import (
	"encoding/json"
	"fmt"
	"net/http"

	"mustaqel/config"
	"mustaqel/models"
)

type RoleRequest struct {
	Name   string
	Type   string
	Status int16
}

func GetRoles(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var roles []models.Role

	if err := config.DB.Find(&roles).Error; err != nil {
		http.Error(w, "", http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]any{
			"message": err.Error(),
			"status":  http.StatusNotFound,
		})
		return
	}

	json.NewEncoder(w).Encode(map[string]any{
		"message": "Role fetched successfully",
		"roles":   roles,
		"status":  200,
	})
}

func GetRole(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	id := r.PathValue("id")

	var role models.Role

	if err := config.DB.First(&role, id).Error; err != nil {
		http.Error(w, "", http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]any{
			"message": err.Error(),
			"status":  http.StatusNotFound,
		})
		return
	}

	json.NewEncoder(w).Encode(map[string]any{
		"message": "Role fetched successfully",
		"role":    role,
		"status":  200,
	})
}

func CreateRole(w http.ResponseWriter, r *http.Request) {

	var role models.Role

	if err := json.NewDecoder(r.Body).Decode(&role); err != nil {
		http.Error(w, "", http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]any{
			"message": err.Error(),
			"status":  http.StatusBadRequest,
		})
		return
	}

	if err := config.DB.Create(&role).Error; err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]any{
			"message": err.Error(),
			"status":  http.StatusInternalServerError,
		})
		return
	}

	json.NewEncoder(w).Encode(map[string]any{
		"message": "Role created successfully",
		"role":    role,
		"status":  200,
	})
}

func UpdateRole(w http.ResponseWriter, r *http.Request) {

	id := r.PathValue("id")

	var role models.Role

	if err := config.DB.First(&role, id).Error; err != nil {
		http.Error(w, "", http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]any{
			"message": err.Error(),
			"status":  http.StatusNotFound,
		})
		return
	}

	var input models.Role

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "", http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]any{
			"message": err.Error(),
			"status":  http.StatusBadRequest,
		})
		return
	}

	if err := config.DB.Model(&role).Updates(input).Error; err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]any{
			"message": err.Error(),
			"status":  http.StatusInternalServerError,
		})
		return
	}

	json.NewEncoder(w).Encode(map[string]any{
		"message": "Role updated successfully",
		"role":    role,
		"status":  200,
	})
}

func DeleteRole(w http.ResponseWriter, r *http.Request) {

	id := r.PathValue("id")

	var role models.Role

	if err := config.DB.First(&role, id).Error; err != nil {
		http.Error(w, "", http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]any{
			"message": err.Error(),
			"status":  http.StatusNotFound,
		})
		return
	}

	roleId := fmt.Sprintf("id=%d", role.ID)

	if err := config.DB.Delete(&role, roleId).Error; err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]any{
			"message": err.Error(),
			"status":  http.StatusInternalServerError,
		})
		return
	}

	json.NewEncoder(w).Encode(map[string]any{
		"message": "Role deleted successfully",
		"status":  200,
	})

}
