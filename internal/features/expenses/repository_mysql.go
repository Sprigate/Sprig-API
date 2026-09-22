package expenses

import (
	"database/sql"
	"errors"
	"sprig/internal/model"
	"time"
)

type expenseRepository struct {
	db *sql.DB
}

func NewExpenseRepository(db *sql.DB) ExpensesRepository {
	return &expenseRepository{db: db}
}

// CREATE
func (r *expenseRepository) AddExpense(expense *model.Expense) error {
	if expense == nil {
		return errors.New("Data tidak boleh kosong!")
	}

	query := "INSERT INTO expenses (user_id, category_id, amount, description, expense_date, type, updated_at, created_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"

	now := time.Now()
	result, err := r.db.Exec(
		query,
		expense.UserID,
		expense.CategoryID,
		expense.Amount,
		expense.Description,
		expense.ExpenseDate,
		expense.Type,
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

	expense.ID = uint64(id)
	expense.UpdatedAt = now
	expense.CreatedAt = now

	return nil
}

// UPDATE
func (r *expenseRepository) Save(expense *model.Expense) error {
	query := "UPDATE expenses SET category_id = ?, amount = ?, description = ?, type = ?, updated_at = ? WHERE id = ?"

	now := time.Now()
	_, err := r.db.Exec(
		query,
		expense.CategoryID,
		expense.Amount,
		expense.Description,
		expense.Type,
		now,
		expense.ID,
	)

	expense.UpdatedAt = now

	return err
}

// DELETE
func (r *expenseRepository) DeleteByID(id uint64) error {
	if id <= 0 {
		return errors.New("ID tidak boleh kosong")
	}

	query := "DELETE FROM expenses WHERE id = ?"
	if _, err := r.db.Exec(query, id); err != nil {
		return err
	}

	return nil
}

// FINDALL
func (r *expenseRepository) FindAll() ([]model.Expense, error) {
	query := "SELECT id, user_id, category_id, amount, description, type, updated_at FROM expenses"

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var expenses []model.Expense

	for rows.Next() {
		var expense model.Expense

		err := rows.Scan(
			&expense.ID,
			&expense.UserID,
			&expense.CategoryID,
			&expense.Amount,
			&expense.Description,
			&expense.Type,
			&expense.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		expenses = append(expenses, expense)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return expenses, nil
}

// FINDBYID
func (r *expenseRepository) FindByID(id uint64) (*model.Expense, error) {
	if id <= 0 {
		return nil, errors.New("ID tidak boleh kosong")
	}

	query := "SELECT id, user_id, category_id, amount, description, type, updated_at FROM expenses WHERE id = ?"

	row := r.db.QueryRow(query, id)

	var expense model.Expense

	err := row.Scan(
		&expense.ID,
		&expense.UserID,
		&expense.CategoryID,
		&expense.Amount,
		&expense.Description,
		&expense.Type,
		&expense.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("Expense tidak dapat ditemukan")
		}

		return nil, err
	}

	return &expense, nil
}

// FINDBYFILTER
func (r *expenseRepository) FindByFilter(filter ExpenseFilterService) ([]model.Expense, error) {
	query, args := buildExpenseQuery(filter)

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var expenses []model.Expense

	for rows.Next() {
		var expense model.Expense

		err := rows.Scan(
			&expense.ID,
			&expense.UserID,
			&expense.CategoryID,
			&expense.Amount,
			&expense.Description,
			&expense.Type,
			&expense.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		expenses = append(expenses, expense)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return expenses, nil
}

// (private) QUERY BUILDER
func buildExpenseQuery(f ExpenseFilterService) (string, []any) {
	query := "SELECT id, user_id, category_id, amount, description, type, updated_at FROM expenses WHERE 1=1"

	args := []any{}
	if !f.StartDate.IsZero() {
		query += " AND expense_date >= ?"
		args = append(args, f.StartDate)
	}

	if !f.EndDate.IsZero() {
		query += " AND expense_date >= ?"
		args = append(args, f.EndDate)
	}

	if f.CategoryID != 0 {
		query += " AND category_id = ?"
		args = append(args, f.CategoryID)
	}

	if f.Type != "" {
		query += " AND type = ?"
		args = append(args, f.Type)
	}

	return query, args
}
