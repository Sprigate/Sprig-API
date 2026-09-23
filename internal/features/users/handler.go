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
		writeJSON(w, http.StatusBadRequest, common.ErrorResponse{
			Message: "invalid JSON",
		})
		return
	}

	user, err := h.userService.CreateUser(req)
	if err != nil {
		writeJSON(w, http.StatusUnprocessableEntity, common.ErrorResponse{
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

	writeJSON(w, http.StatusCreated, common.SuccessResponse{
		Message: "Berhasil membuat user",
		Data:    response,
	})
}

// FindAll godoc
//
//	@Summary		Ambil semua user
//	@Description	Mengembalikan daftar seluruh user
//	@Tags			users
//	@Produce		json
//	@Success		200	{object}	common.SuccessResponse
//	@Failure		404	{object}	common.ErrorResponse
//	@Router			/api/v1/users [get]
func (h *UserHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	users, err := h.userService.FindAllUser()
	if err != nil {
		writeJSON(w, http.StatusBadRequest, common.ErrorResponse{
			Message: "Gagal mendapatkan data user",
		})
		return
	}

	responses := make([]UserResponse, 0, len(users))

	for _, user := range users {
		responses = append(responses, UserResponse{
			ID:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
		})
	}

	writeJSON(w, http.StatusOK, common.SuccessResponse{
		Data: responses,
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
		writeJSON(w, http.StatusBadRequest, common.ErrorResponse{
			Message: "invalid JSON",
		})
		return
	}

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, common.ErrorResponse{
			Message: "invalid JSON",
		})
		return
	}

	user, err := h.userService.UpdateUser(id, req)
	if err != nil {
		writeJSON(w, http.StatusUnprocessableEntity, common.ErrorResponse{
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

	writeJSON(w, http.StatusOK, common.SuccessResponse{
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
		writeJSON(w, http.StatusBadRequest, common.ErrorResponse{
			Message: "invalid user id",
		})
		return
	}

	err = h.userService.DeleteUser(id)
	if err != nil {
		writeJSON(w, http.StatusUnprocessableEntity, common.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	writeJSON(w, http.StatusOK, common.SuccessResponse{
		Message: "Data berhasil dihapus",
	})
}

// helper (nanti kalao udah banyak yang make pindahin aja)
func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(data)
}
