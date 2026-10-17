package categories

import (
	"encoding/json"
	"net/http"
	"sprig/internal/common"
	"sprig/internal/model"
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
		common.WriteJSON(w, http.StatusBadRequest, common.ErrorResponse{
			Message: "invalid JSON",
		})
		return
	}

	category, err := h.categoryService.CreateCategory(req)
	if err != nil {
		common.WriteJSON(w, http.StatusUnprocessableEntity, common.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	response := CategoryResponse{
		ID:     category.ID,
		UserID: category.UserID,
		Name:   category.Name,
	}

	common.WriteJSON(w, http.StatusCreated, common.SuccessResponse{
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
		common.WriteJSON(w, http.StatusNotFound, common.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	response := CategoryResponse{
		ID:     category.ID,
		UserID: category.UserID,
		Name:   category.Name,
	}

	common.WriteJSON(w, http.StatusOK, common.SuccessResponse{
		Data: response,
	})
}

// FindAll godoc
//
//	@Summary		Ambil daftar kategori
//	@Description	Mengembalikan daftar kategori dengan pagination dan pencarian opsional berdasarkan nama
//	@Tags			categories
//	@Produce		json
//	@Param			page	query		int		false	"Nomor halaman (minimal 1)"			default(1)	minimum(1)
//	@Param			limit	query		int		false	"Jumlah data per halaman (maks 30)"	default(10)	minimum(1)	maximum(30)
//	@Success		200		{object}	common.SuccessResponse
//	@Failure		400		{object}	common.ErrorResponse	"page atau limit bukan angka"
//	@Failure		500		{object}	common.ErrorResponse	"Gagal mengambil data kategori"
//	@Router			/api/v1/categories [get]
func (h *CategoryHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()

	// parsing parameter page & limit
	page, err := common.ParseIntParam(q.Get("page"), 1)
	if err != nil {
		common.WriteJSON(w, http.StatusBadRequest, common.ErrorResponse{
			Message: "Parameter page harus berupa angka",
		})
		return
	}

	limit, err := common.ParseIntParam(q.Get("limit"), 10)
	if err != nil {
		common.WriteJSON(w, http.StatusBadRequest, common.ErrorResponse{
			Message: "Parameter limit harus berupa angka",
		})
		return
	}

	categories, err := h.categoryService.FindAllCategory(model.PaginationParams{Page: page, Limit: limit})
	if err != nil {
		common.WriteJSON(w, http.StatusBadRequest, common.ErrorResponse{
			Message: "Gagal mengambil data kategori",
		})
		return
	}

	responses := make([]CategoryResponse, 0, len(categories.Data))

	for _, category := range categories.Data {
		responses = append(responses, CategoryResponse{
			ID:     category.ID,
			UserID: category.UserID,
			Name:   category.Name,
		})
	}

	common.WriteJSON(w, http.StatusOK, common.SuccessResponse{
		Message: "Data kategori berhasil didapat",
		Data:    responses,
		Meta: &common.PaginationMeta{
			Page:       categories.Page,
			Limit:      categories.Limit,
			TotalItems: categories.TotalItems,
			TotalPages: categories.TotalPages,
		},
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

	category, err := h.categoryService.UpdateCategory(id, req)
	if err != nil {
		common.WriteJSON(w, http.StatusUnprocessableEntity, common.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	response := CategoryResponse{
		ID:     category.ID,
		UserID: category.UserID,
		Name:   category.Name,
	}

	common.WriteJSON(w, http.StatusOK, common.SuccessResponse{
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
		common.WriteJSON(w, http.StatusBadRequest, common.ErrorResponse{
			Message: "invalid JSON",
		})
		return
	}

	err = h.categoryService.DeleteCategory(id)
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
