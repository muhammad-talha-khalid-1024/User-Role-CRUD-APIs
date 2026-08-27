package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"

	"mustaqel/internal/domain"
	"mustaqel/internal/service"
	"mustaqel/pkg/utils"
)

type UserHandler struct {
	service  *service.UserService
	validate *validator.Validate
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{
		service:  service,
		validate: validator.New(),
	}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req domain.CreateUserRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if err := h.validate.Struct(req); err != nil {
		utils.RespondWithValidationError(w, err)
		return
	}

	user, err := h.service.CreateUser(&req)
	if err != nil {
		switch err {
		case domain.ErrUserAlreadyExists:
			utils.RespondWithError(w, http.StatusConflict, "User with this email already exists")
		default:
			utils.RespondWithError(w, http.StatusInternalServerError, "Failed to create user")
		}
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, user)
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		utils.RespondWithError(w, http.StatusBadRequest, "User ID is required")
		return
	}

	user, err := h.service.GetUser(id)
	if err != nil {
		switch err {
		case domain.ErrUserNotFound:
			utils.RespondWithError(w, http.StatusNotFound, "User not found")
		default:
			utils.RespondWithError(w, http.StatusInternalServerError, "Failed to get user")
		}
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, user)
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		utils.RespondWithError(w, http.StatusBadRequest, "User ID is required")
		return
	}

	var req domain.UpdateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if err := h.validate.Struct(req); err != nil {
		utils.RespondWithValidationError(w, err)
		return
	}

	user, err := h.service.UpdateUser(id, &req)
	if err != nil {
		switch err {
		case domain.ErrUserNotFound:
			utils.RespondWithError(w, http.StatusNotFound, "User not found")
		default:
			utils.RespondWithError(w, http.StatusInternalServerError, "Failed to update user")
		}
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, user)
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		utils.RespondWithError(w, http.StatusBadRequest, "User ID is required")
		return
	}

	err := h.service.DeleteUser(id)
	if err != nil {
		switch err {
		case domain.ErrUserNotFound:
			utils.RespondWithError(w, http.StatusNotFound, "User not found")
		default:
			utils.RespondWithError(w, http.StatusInternalServerError, "Failed to delete user")
		}
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "User deleted successfully"})
}

func (h *UserHandler) ListUsers(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	pageSize, _ := strconv.Atoi(r.URL.Query().Get("page_size"))

	users, total, totalPages, err := h.service.ListUsers(page, pageSize)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to list users")
		return
	}

	response := map[string]interface{}{
		"data": users,
		"pagination": map[string]interface{}{
			"current_page": page,
			"page_size":    pageSize,
			"total":        total,
			"total_pages":  totalPages,
		},
	}

	utils.RespondWithJSON(w, http.StatusOK, response)
}
