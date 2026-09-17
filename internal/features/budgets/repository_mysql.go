package budgets

import (
	"database/sql"
	"errors"
	"sprig/internal/model"
	"time"
)

type budgetRepository struct {
	db *sql.DB
}

func NewBudgetRepository(db *sql.DB) BudgetsRepository {
	return &budgetRepository{db: db}
}

func (r *budgetRepository) AddBudget(budget *model.Budget) error {
	if budget == nil {
		return errors.New("Data tidak boleh kosong!")
	}

	query := "INSERT INTO budgets (user_id, total_income, needs_percentage, wants_percentage, savings_percentage, year, month, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)"

	now := time.Now()
	result, err := r.db.Exec(
		query,
		budget.UserID,
		budget.TotalIncome,
		budget.NeedsPercentage,
		budget.WantsPercentage,
		budget.SavingsPercentage,
		budget.Year,
		budget.Month,
		now,
		now,
	)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	budget.ID = uint64(id)
	budget.CreatedAt = now
	budget.UpdatedAt = now

	return nil
}

func (r *budgetRepository) Save(budget *model.Budget) error {
	query := "UPDATE budgets SET total_income = ?, needs_percentage = ?, wants_percentage = ?, savings_percentage = ?, month = ?, year = ?, updated_at = ?"

	now := time.Now()
	_, err := r.db.Exec(
		query,
		budget.TotalIncome,
		budget.NeedsPercentage,
		budget.WantsPercentage,
		budget.SavingsPercentage,
		budget.Month,
		budget.Year,
		now,
	)

	budget.UpdatedAt = now

	return err
}

func (r *budgetRepository) DeleteByID(id uint64) error {
	if id <= 0 {
		return errors.New("ID tidak boleh kosong")
	}

	query := "DELETE FROM budgets WHERE id = ?"
	if _, err := r.db.Exec(query, id); err != nil {
		return err
	}

	return nil
}

func (r *budgetRepository) FindByID(id uint64) (*model.Budget, error) {
	if id <= 0 {
		return nil, errors.New("ID tidak boleh kosong")
	}

	query := "SELECT id, user_id, total_income, needs_percentage, wants_percentage, savings_percentage, month, year, updated_at FROM budgets WHERE id = ?"
	row := r.db.QueryRow(query, id)

	var budget model.Budget

	err := row.Scan(
		&budget.ID,
		&budget.UserID,
		&budget.TotalIncome,
		&budget.NeedsPercentage,
		&budget.WantsPercentage,
		&budget.SavingsPercentage,
		&budget.Month,
		&budget.Year,
		&budget.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("Budget tidak dapat ditemukan")
		}

		return nil, err
	}

	return &budget, nil
}

func (r *budgetRepository) FindAll() ([]model.Budget, error) {
	query := "SELECT id, user_id, total_income, needs_percentage, wants_percentage, savings_percentage, month, year, updated_at FROM budgets"

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var budgets []model.Budget

	for rows.Next() {
		var budget model.Budget

		// NOTES : buat bagian scan, harus sesuai posisi column yang dimintanya sama query
		// kalau nggak, nanti bakal error. karna ketidak sesuaian
		// tipe data yang diminta, atau apapun
		err := rows.Scan(
			&budget.ID,
			&budget.UserID,
			&budget.TotalIncome,
			&budget.NeedsPercentage,
			&budget.WantsPercentage,
			&budget.SavingsPercentage,
			&budget.Month,
			&budget.Year,
			&budget.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		budgets = append(budgets, budget)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return budgets, nil
}

// NOTES : Next time ubah responsenya jadi 1 data saja. jangan array
func (r *budgetRepository) FilterByMonthAndYear(month uint8, year uint16) ([]model.Budget, error) {
	if month <= 0 && year <= 0 {
		return nil, errors.New("Data bulan atau tahun tidak boleh kosong")
	}

	query := "SELECT id, user_id, total_income, needs_percentage, wants_percentage, savings_percentage, month, year, updated_at FROM budgets WHERE month = ? AND year = ?"

	rows, err := r.db.Query(query, month, year)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var budgets []model.Budget

	for rows.Next() {
		var budget model.Budget

		err := rows.Scan(
			&budget.ID,
			&budget.UserID,
			&budget.TotalIncome,
			&budget.NeedsPercentage,
			&budget.WantsPercentage,
			&budget.SavingsPercentage,
			&budget.Month,
			&budget.Year,
			&budget.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		budgets = append(budgets, budget)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return budgets, nil
}
