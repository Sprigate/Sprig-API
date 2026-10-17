package users

import (
	"encoding/json"
	"net/http"
	"sprig/internal/common"
	"strconv"
)

type UserHandler struct {
	userService *UserService
}

func NewUserHandler(userService *UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// Create godoc
//
//	@Summary	Tambah user baru
//	@Tags		users
//	@Accept		json
//	@Produce	json
//	@Param		user	body		CreateUserRequest	true	"Data user baru"
//	@Success	201		{object}	common.SuccessResponse
//	@Failure	400		{object}	common.ErrorResponse
//	@Failure	422		{object}	common.ErrorResponse
//	@Router		/api/v1/users [post]
func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req CreateUserRequest // tampung dulu di DTO

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		common.WriteJSON(w, http.StatusBadRequest, common.ErrorResponse{
			Message: "invalid JSON",
		})
		return
	}

	user, err := h.userService.CreateUser(req)
	if err != nil {
		common.WriteJSON(w, http.StatusUnprocessableEntity, common.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	response := UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}

	common.WriteJSON(w, http.StatusCreated, common.SuccessResponse{
		Message: "Berhasil membuat user",
		Data:    response,
	})
}

// UpdateUser godoc
//
//	@Summary	Update sebagian data user
//	@Tags		users
//	@Accept		json
//	@Produce	json
//	@Param		id		path		int						true	"User ID"
//	@Param		user	body		CreateUserUpdateRequest	true	"Field yang ingin diupdate"
//	@Success	200		{object}	common.SuccessResponse
//	@Failure	404		{object}	common.ErrorResponse
//	@Failure	422		{object}	common.ErrorResponse
//	@Router		/api/v1/users/{id} [patch]
func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")

	var req CreateUserUpdateRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		common.WriteJSON(w, http.StatusBadRequest, common.ErrorResponse{
			Message: "invalid JSON",
		})
		return
	}

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		common.WriteJSON(w, http.StatusBadRequest, common.ErrorResponse{
			Message: "invalid JSON",
		})
		return
	}

	user, err := h.userService.UpdateUser(id, req)
	if err != nil {
		common.WriteJSON(w, http.StatusUnprocessableEntity, common.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	response := UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}

	common.WriteJSON(w, http.StatusOK, common.SuccessResponse{
		Message: "User berhasil diperbarui",
		Data:    response,
	})
}

// Delete godoc
//
//	@Summary	Hapus user
//	@Tags		users
//	@Produce	json
//	@Param		id	path		int	true	"User ID"
//	@Success	200	{object}	common.SuccessResponse
//	@Failure	404	{object}	common.ErrorResponse
//	@Failure	422	{object}	common.ErrorResponse
//	@Router		/api/v1/users/{id} [delete]
func (h *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		common.WriteJSON(w, http.StatusBadRequest, common.ErrorResponse{
			Message: "invalid user id",
		})
		return
	}

	err = h.userService.DeleteUser(id)
	if err != nil {
		common.WriteJSON(w, http.StatusUnprocessableEntity, common.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	common.WriteJSON(w, http.StatusOK, common.SuccessResponse{
		Message: "Data berhasil dihapus",
	})
}
