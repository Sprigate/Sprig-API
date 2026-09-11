package categories

import (
	"encoding/json"
	"net/http"
	"sprig/internal/common"
	"strconv"
)

type CategoryHandler struct {
	categoryService *CategoryService
}

func NewCategoryHandler(categoryService *CategoryService) *CategoryHandler {
	return &CategoryHandler{
		categoryService: categoryService,
	}
}

// Create godoc
//
//	@Summary	Tambah category baru
//	@Tags		categories
//	@Accept		json
//	@Produce	json
//	@Param		category	body		CreateCategoryRequest	true	"Data kategori baru"
//	@Success	201			{object}	common.SuccessResponse
//	@Failure	400			{object}	common.ErrorResponse
//	@Failure	422			{object}	common.ErrorResponse
//	@Router		/api/v1/categories [post]
func (h *CategoryHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req CreateCategoryRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, common.ErrorResponse{
			Message: "invalid JSON",
		})
		return
	}

	category, err := h.categoryService.CreateCategory(req)
	if err != nil {
		writeJSON(w, http.StatusUnprocessableEntity, common.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	response := CategoryResponse{
		ID:     category.ID,
		UserID: category.UserID,
		Name:   category.Name,
	}

	writeJSON(w, http.StatusCreated, common.SuccessResponse{
		Message: "Berhasil membuat kategori",
		Data:    response,
	})
}

// FindByName godoc
//
//	@Summary	Ambil detail kategori berdasarkan Nama
//	@Tags		categories
//	@Produce	json
//	@Param		name	path		string	true	"Nama Kategori"
//	@Success	200		{object}	common.SuccessResponse
//	@Failure	404		{object}	common.ErrorResponse
//	@Router		/api/v1/categories/{name} [get]
func (h *CategoryHandler) FindByName(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")

	category, err := h.categoryService.FindCategory(name)
	if err != nil {
		writeJSON(w, http.StatusNotFound, common.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	response := CategoryResponse{
		ID:     category.ID,
		UserID: category.UserID,
		Name:   category.Name,
	}

	writeJSON(w, http.StatusOK, common.SuccessResponse{
		Data: response,
	})
}

// FindAll godoc
//
//	@Summary		Ambil semua kategori
//	@Description	Mengembalikan daftar seluruh kategori
//	@Tags			categories
//	@Produce		json
//	@Success		200	{object}	common.SuccessResponse
//	@Failure		400	{object}	common.ErrorResponse
//	@Router			/api/v1/categories [get]
func (h *CategoryHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	categories, err := h.categoryService.FindAllCategory()
	if err != nil {
		writeJSON(w, http.StatusBadRequest, common.ErrorResponse{
			Message: "Gagal mengambil data kategori",
		})
		return
	}

	responses := make([]CategoryResponse, 0, len(categories))

	for _, category := range categories {
		responses = append(responses, CategoryResponse{
			ID:     category.ID,
			UserID: category.UserID,
			Name:   category.Name,
		})
	}

	writeJSON(w, http.StatusOK, common.SuccessResponse{
		Data: responses,
	})
}

// UpdateCategory godoc
//
//	@Summary	Update sebagian data Kategori
//	@Tags		categories
//	@Accept		json
//	@Produce	json
//	@Param		id			path		int							true	"ID Kategori"
//	@Param		category	body		CreateCategoryUpdateRequest	true	"Field yang ingin diupdate"
//	@Success	200			{object}	common.SuccessResponse
//	@Failure	400			{object}	common.ErrorResponse
//	@Failure	422			{object}	common.ErrorResponse
//	@Router		/api/v1/categories/{id} [patch]
func (h *CategoryHandler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")

	var req CreateCategoryUpdateRequest

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

	category, err := h.categoryService.UpdateCategory(id, req)
	if err != nil {
		writeJSON(w, http.StatusUnprocessableEntity, common.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	response := CategoryResponse{
		ID:     category.ID,
		UserID: category.UserID,
		Name:   category.Name,
	}

	writeJSON(w, http.StatusOK, common.SuccessResponse{
		Message: "Kategori berhasil di perbarui",
		Data:    response,
	})
}

// Delete godoc
//
//	@Summary	Hapus kategori
//	@Tags		categories
//	@Produce	json
//	@Param		id	path		int	true	"ID Kategori"
//	@Success	200	{object}	common.SuccessResponse
//	@Failure	422	{object}	common.ErrorResponse
//	@Router		/api/v1/categories/{id} [delete]
func (h *CategoryHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, common.ErrorResponse{
			Message: "invalid JSON",
		})
		return
	}

	err = h.categoryService.DeleteCategory(id)
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

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(data)
}
