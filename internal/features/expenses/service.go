package expenses

import (
	"errors"
	"sprig/internal/model"
	"strings"
	"time"
)

type ExpenseService struct {
	repository ExpensesRepository
}

func NewExpenseService(repository ExpensesRepository) *ExpenseService {
	return &ExpenseService{
		repository: repository,
	}
}

func (s *ExpenseService) CreateExpense(req CreateExpenseRequest) (*model.Expense, error) {
	if req.CategoryID <= 0 {
		return nil, errors.New("Kategori pengeluaran tidak boleh kosong")
	}

	if req.Amount <= 500 {
		return nil, errors.New("Nilai pengeluaran tidak boleh kurang dari Rp 500")
	}

	if req.Type == "" {
		return nil, errors.New("Tipe pengeluaran tidak boleh kosong")
	}

	expense_date := parseDate(req.ExpenseDate)

	expense := &model.Expense{
		UserID:      req.UserID,
		CategoryID:  req.CategoryID,
		Amount:      req.Amount,
		Description: req.Description,
		ExpenseDate: expense_date,
		Type:        req.Type,
	}

	if err := s.repository.AddExpense(expense); err != nil {
		return nil, err
	}

	return expense, nil
}

func (s *ExpenseService) UpdateExpense(id uint64, req CreateUpdateExpenseRequest) (*model.Expense, error) {
	existingExpense, err := s.repository.FindByID(id)
	if err != nil {
		return nil, err
	}

	if req.CategoryID != nil && *req.CategoryID != 0 {
		existingExpense.CategoryID = *req.CategoryID
	}

	if req.Amount != nil && *req.Amount != 0 {
		if *req.Amount <= 500 {
			return nil, errors.New("Nilai pengeluaran tidak boleh kurang dari Rp 500")
		}

		existingExpense.Amount = *req.Amount
	}

	description := *req.Description
	if req.Description != nil && strings.TrimSpace(*req.Description) != "" {
		existingExpense.Description = &description
	}

	expense_date := parseDate(*req.ExpenseDate)
	if req.ExpenseDate != nil && !expense_date.IsZero() {
		existingExpense.ExpenseDate = expense_date
	}

	// ---- Debug State ----
	// log.Printf("[DEBUG] Expense isEmpty : %t", expense_date.IsZero())
	// log.Printf("[DEBUG] Expense raw value : %v", expense_date)

	if err = s.repository.Save(existingExpense); err != nil {
		return nil, err
	}

	return existingExpense, nil
}

func (s *ExpenseService) DeleteExpense(id uint64) error {
	if id <= 0 {
		return errors.New("ID tidak boleh kosong")
	}

	if err := s.repository.DeleteByID(id); err != nil {
		return err
	}

	return nil
}

func (s *ExpenseService) FindAllExpense() ([]model.Expense, error) {
	expense, err := s.repository.FindAll()
	if err != nil {
		return nil, err
	}

	return expense, nil
}

type ExpenseFilterService struct {
	StartDate  time.Time
	EndDate    time.Time
	CategoryID uint64
	Type       model.ExpenseType
}

func (s *ExpenseService) FilteredData(filter ExpenseFilterService) ([]model.Expense, error) {
	expenses, err := s.repository.FindByFilter(filter)
	if err != nil {
		return nil, err
	}

	return expenses, nil
}
