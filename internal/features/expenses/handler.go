package expenses

import (
	"encoding/json"
	"log"
	"net/http"
	"sprig/internal/common"
	"sprig/internal/model"
	"strconv"
	"time"
)

type ExpenseHandler struct {
	expenseService *ExpenseService
}

func NewExpenseHandler(expenseService *ExpenseService) *ExpenseHandler {
	return &ExpenseHandler{
		expenseService: expenseService,
	}
}

// Create godoc
//
//	@Summary	Tambah expense / pengeluaran baru
//	@Tags		expenses
//	@Accept		json
//	@Produce	json
//	@Param		expense	body		CreateExpenseRequest	true	"Data expense / pengeluaran baru"
//	@Success	201		{object}	common.SuccessResponse
//	@Failure	400		{object}	common.ErrorResponse
//	@Failure	422		{object}	common.ErrorResponse
//	@Router		/api/v1/expenses [post]
func (h *ExpenseHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req CreateExpenseRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, common.ErrorResponse{
			Message: "invalid JSON",
		})

		log.Printf("[LOG] %v", err)
		return
	}

	expense, err := h.expenseService.CreateExpense(req)
	if err != nil {
		writeJSON(w, http.StatusUnprocessableEntity, common.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	response := ExpenseResponse{
		ID:          expense.ID,
		UserID:      expense.UserID,
		CategoryID:  expense.CategoryID,
		Amount:      expense.Amount,
		Description: *expense.Description,
		ExpenseDate: expense.ExpenseDate,
		Type:        expense.Type,
		UpdatedAt:   expense.UpdatedAt,
	}

	writeJSON(w, http.StatusCreated, common.SuccessResponse{
		Message: "Berhasil membuat expense",
		Data:    response,
	})
}

// Update godoc
//
//	@Summary	Update sebagian data expense / pengeluaran
//	@Tags		expenses
//	@Accept		json
//	@Produce	json
//	@Param		id		path		int							true	"Expense ID"
//	@Param		expense	body		CreateUpdateExpenseRequest	false	"Field yang ingin diupdate"
//	@Success	201		{object}	common.SuccessResponse
//	@Failure	400		{object}	common.ErrorResponse
//	@Failure	422		{object}	common.ErrorResponse
//	@Router		/api/v1/expenses/{id} [patch]
func (h *ExpenseHandler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")

	var req CreateUpdateExpenseRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
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

	expense, err := h.expenseService.UpdateExpense(id, req)
	if err != nil {
		writeJSON(w, http.StatusUnprocessableEntity, common.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	response := ExpenseResponse{
		ID:          expense.ID,
		UserID:      expense.UserID,
		CategoryID:  expense.CategoryID,
		Amount:      expense.Amount,
		Description: *expense.Description,
		ExpenseDate: expense.ExpenseDate,
		Type:        expense.Type,
		UpdatedAt:   expense.UpdatedAt,
	}

	writeJSON(w, http.StatusCreated, common.SuccessResponse{
		Message: "Expense berhasil diperbarui",
		Data:    response,
	})
}

// Delete godoc
//
//	@Summary	Hapus expense / pengeluaran
//	@Tags		expenses
//	@Produce	json
//	@Param		id	path		int	true	"Expense ID"
//	@Success	200	{object}	common.SuccessResponse
//	@Failure	400	{object}	common.ErrorResponse
//	@Failure	422	{object}	common.ErrorResponse
//	@Router		/api/v1/expenses/{id} [delete]
func (h *ExpenseHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, common.ErrorResponse{
			Message: "invalid JSON",
		})
		return
	}

	if err := h.expenseService.DeleteExpense(id); err != nil {
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
//	@Summary		Ambil semua expense / pengeluaran
//	@Description	Mengembalikan daftar seluruh expense / pengeluaran
//	@Tags			expenses
//	@Produce		json
//	@Param			start_date	query		string	false	"Expense start_date (Format = YYYY-MM-DD)"
//	@Param			end_date	query		string	false	"Expense end_date (Format = YYYY-MM-DD)"
//	@Param			category_id	query		int		false	"Expense category id"
//	@Param			type		query		string	false	"Expense type"
//	@Success		200			{object}	common.SuccessResponse
//	@Failure		400			{object}	common.ErrorResponse
//	@Router			/api/v1/expenses [get]
func (h *ExpenseHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	filter := ExpenseFilterRequest{
		StartDate:  query.Get("start_date"),
		EndDate:    query.Get("end_date"),
		CategoryID: query.Get("category_id"),
		Type:       query.Get("type"),
	}

	var (
		expenses []model.Expense
		err      error
	)

	if filter.StartDate == "" || filter.EndDate == "" || filter.CategoryID == "" || filter.Type == "" {
		expenses, err = h.expenseService.FindAllExpense()
		if err != nil {
			writeJSON(w, http.StatusBadRequest, common.ErrorResponse{
				Message: "Gagal mendapatkan data Expense",
			})

			log.Printf("[LOG] %v", err)
			return
		}
	} else {
		// parsing startdate, enddate, & categoryid
		start_date := parseDate(filter.StartDate)
		end_date := parseDate(filter.EndDate)
		category_id, _ := strconv.ParseUint(filter.CategoryID, 10, 64)

		f := ExpenseFilterService{
			StartDate:  start_date,
			EndDate:    end_date,
			CategoryID: category_id,
			Type:       model.ExpenseType(filter.Type),
		}

		expenses, err = h.expenseService.FilteredData(f)
		if err != nil {
			writeJSON(w, http.StatusUnprocessableEntity, common.ErrorResponse{
				Message: err.Error(),
			})
			return
		}
	}

	responses := make([]ExpenseResponse, 0, len(expenses))

	for _, expense := range expenses {
		responses = append(responses, ExpenseResponse{
			ID:          expense.ID,
			UserID:      expense.UserID,
			CategoryID:  expense.CategoryID,
			Amount:      expense.Amount,
			Description: *expense.Description,
			ExpenseDate: expense.ExpenseDate,
			Type:        expense.Type,
			UpdatedAt:   expense.UpdatedAt,
		})
	}

	writeJSON(w, http.StatusOK, common.SuccessResponse{
		Data: responses,
	})
}

func parseDate(date string) time.Time {
	t, _ := time.Parse(time.DateOnly, date)
	return t
}

// helper (nanti kalao udah banyak yang make pindahin aja)
func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(data)
}
