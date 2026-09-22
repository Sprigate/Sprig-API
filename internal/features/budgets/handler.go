package budgets

import (
	"encoding/json"
	"log"
	"net/http"
	"sprig/internal/common"
	"sprig/internal/model"
	"strconv"
)

type BudgetHandler struct {
	budgetService *BudgetService
}

func NewBudgetHandler(budgetService *BudgetService) *BudgetHandler {
	return &BudgetHandler{
		budgetService: budgetService,
	}
}

// Create godoc
//
//	@Summary	Tambah budget baru
//	@Tags		budgets
//	@Accept		json
//	@Produce	json
//	@Param		budget	body		CreateBudgetRequest	true	"Data budget baru"
//	@Success	201		{object}	common.SuccessResponse
//	@Failure	400		{object}	common.ErrorResponse
//	@Failure	422		{object}	common.ErrorResponse
//	@Router		/api/v1/budgets [post]
func (h *BudgetHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req CreateBudgetRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, common.ErrorResponse{
			Message: "invalid JSON",
		})
		return
	}

	budget, err := h.budgetService.CreateBudget(req)
	if err != nil {
		writeJSON(w, http.StatusUnprocessableEntity, common.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	response := BudgetResponse{
		ID:                budget.ID,
		UserID:            budget.UserID,
		TotalIncome:       budget.TotalIncome,
		NeedsPercentage:   budget.NeedsPercentage,
		WantsPercentage:   budget.WantsPercentage,
		SavingsPercentage: budget.SavingsPercentage,
		Month:             budget.Month,
		Year:              budget.Year,
		UpdatedAt:         budget.UpdatedAt,
	}

	writeJSON(w, http.StatusCreated, common.SuccessResponse{
		Message: "Berhasil membuat budget",
		Data:    response,
	})
}

// Update godoc
//
//	@Summary	Update sebagian data budget
//	@Tags		budgets
//	@Accept		json
//	@Produce	json
//	@Param		id		path		int							true	"Budget ID"
//	@Param		budget	body		CreateBudgetUpdateRequest	false	"Field yang ingin diupdate"
//	@Success	200		{object}	common.SuccessResponse
//	@Failure	400		{object}	common.ErrorResponse
//	@Failure	422		{object}	common.ErrorResponse
//	@Router		/api/v1/budgets/{id} [patch]
func (h *BudgetHandler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")

	var req CreateBudgetUpdateRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, common.ErrorResponse{
			Message: "invalid JSON",
		})
		return
	}

	// log.Printf("%#v", &req)

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, common.ErrorResponse{
			Message: "invalid JSON",
		})
		return
	}

	budget, err := h.budgetService.UpdateBudget(id, req)
	if err != nil {
		writeJSON(w, http.StatusUnprocessableEntity, common.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	response := BudgetResponse{
		ID:                budget.ID,
		UserID:            budget.UserID,
		TotalIncome:       budget.TotalIncome,
		NeedsPercentage:   budget.NeedsPercentage,
		WantsPercentage:   budget.WantsPercentage,
		SavingsPercentage: budget.SavingsPercentage,
		Month:             budget.Month,
		Year:              budget.Year,
		UpdatedAt:         budget.UpdatedAt,
	}

	writeJSON(w, http.StatusOK, common.SuccessResponse{
		Message: "Budget berhasil diperbarui",
		Data:    response,
	})
}

// Delete godoc
//
//	@Summary	Hapus budget
//	@Tags		budgets
//	@Produce	json
//	@Param		id	path		int	true	"Budget ID"
//	@Success	200	{object}	common.SuccessResponse
//	@Failure	400	{object}	common.ErrorResponse
//	@Failure	422	{object}	common.ErrorResponse
//	@Router		/api/v1/budgets/{id} [delete]
func (h *BudgetHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, common.ErrorResponse{
			Message: "invalid JSON",
		})
		return
	}

	if err := h.budgetService.DeleteBudget(id); err != nil {
		writeJSON(w, http.StatusUnprocessableEntity, common.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	writeJSON(w, http.StatusOK, common.SuccessResponse{
		Message: "Data berhasil dihapus",
	})
}

// FindAll godoc
//
//	@Summary		Ambil semua budget
//	@Description	Mengembalikan daftar seluruh budget
//	@Tags			budgets
//	@Produce		json
//	@Param			month	query		int	false	"Budget Month"
//	@Param			year	query		int	false	"Budget Year"
//	@Success		200		{object}	common.SuccessResponse
//	@Failure		400		{object}	common.ErrorResponse
//	@Router			/api/v1/budgets [get]
func (h *BudgetHandler) FindAll(w http.ResponseWriter, r *http.Request) {

	// NOTES : Karna ini parameternya seperti contoh dibawah
	// api/v1/budgets?month=1&year=2026
	// yang dimana ini merupakan query parameter, bukan path parameter
	// maka pengambilan data nya berbeda
	monthStr := r.URL.Query().Get("month")
	yearStr := r.URL.Query().Get("year")

	log.Printf("month\t: %s", monthStr)
	log.Printf("year\t: %s", yearStr)

	var (
		budgets []model.Budget
		err     error
	)

	// tanpa query parameter
	if monthStr == "" || yearStr == "" {
		budgets, err = h.budgetService.FindAllBudget()
		if err != nil {
			writeJSON(w, http.StatusBadRequest, common.ErrorResponse{
				Message: "Gagal mendapatkan data budget",
			})

			log.Printf("Error message : %v", err)
			return
		}
	} else {
		// dengan parameter

		month, err := strconv.ParseUint(monthStr, 10, 8)
		if err != nil {
			writeJSON(w, http.StatusBadRequest, common.ErrorResponse{
				Message: "invalid JSON",
			})
			return
		}

		year, err := strconv.ParseUint(yearStr, 10, 16)
		if err != nil {
			writeJSON(w, http.StatusBadRequest, common.ErrorResponse{
				Message: "invalid JSON",
			})
			return
		}

		// NOTE : Karna hasil parse dari strconv.ParseUint masih
		// bertipe uint64, jadi konversi lagi aja.
		// example : uint8(month). uint64 -> uint8
		budgets, err = h.budgetService.FilterByMonthAndYear(uint8(month), uint16(year))
		if err != nil {
			writeJSON(w, http.StatusUnprocessableEntity, common.ErrorResponse{
				Message: err.Error(),
			})
			return
		}

		// DEBUG
		// log.Printf("%v, (month) data type\t: %T", uint8(month), uint8(month))
		// log.Printf("%v, (year) data type\t: %T", uint16(year), uint16(year))
	}

	responses := make([]BudgetResponse, 0, len(budgets))

	for _, budget := range budgets {
		responses = append(responses, BudgetResponse{
			ID:                budget.ID,
			UserID:            budget.UserID,
			TotalIncome:       budget.TotalIncome,
			NeedsPercentage:   budget.NeedsPercentage,
			WantsPercentage:   budget.WantsPercentage,
			SavingsPercentage: budget.SavingsPercentage,
			Month:             budget.Month,
			Year:              budget.Year,
			UpdatedAt:         budget.UpdatedAt,
		})
	}

	writeJSON(w, http.StatusOK, common.SuccessResponse{
		Data: responses,
	})
}

// helper (nanti kalao udah banyak yang make pindahin aja)
func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(data)
}
