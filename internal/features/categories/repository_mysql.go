package categories

import (
	"database/sql"
	"errors"
	"sprig/internal/model"
	"strings"
	"time"
)

type categoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) CategoryRepository {
	return &categoryRepository{db: db}
}

func (r *categoryRepository) AddCategory(category *model.Category) error {
	if category == nil {
		return errors.New("Data tidak boleh kosong!")
	}

	query := "INSERT INTO categories (user_id, name) VALUES (?, ?)"

	_, err := r.db.Exec(
		query,
		category.UserID,
		category.Name,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *categoryRepository) FindCategoryByName(name string) (*model.Category, error) {
	if strings.TrimSpace(name) == "" {
		// coba nanti ubah pake sentinel errors aja. males ngetik ulang
		return nil, errors.New("Data tidak boleh kosong")
	}

	query := "SELECT id, user_id, name FROM categories WHERE name = ?"

	row := r.db.QueryRow(query, name)

	var category model.Category
	err := row.Scan(
		&category.ID,
		&category.UserID,
		&category.Name,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("Kategori tidak dapat ditemukan")
		}

		return nil, err
	}

	return &category, nil
}

func (r *categoryRepository) FindByID(id uint64) (*model.Category, error) {
	if id <= 0 {
		return nil, errors.New("ID tidak boleh kurang dari 0")
	}

	query := "SELECT id, user_id, name FROM categories WHERE id = ?"

	row := r.db.QueryRow(query, id)

	var category model.Category
	if err := row.Scan(
		&category.ID,
		&category.UserID,
		&category.Name,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("Kategori tidak dapat ditemukan")
		}

		return nil, err
	}

	return &category, nil
}

func (r *categoryRepository) FindAll() ([]model.Category, error) {
	query := "SELECT id, user_id, name FROM categories"

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []model.Category

	for rows.Next() {
		var category model.Category

		err := rows.Scan(
			&category.ID,
			&category.UserID,
			&category.Name,
		)
		if err != nil {
			return nil, err
		}

		categories = append(categories, category)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return categories, nil
}

func (r *categoryRepository) Save(category *model.Category) error {
	query := "UPDATE categories SET name = ?, updated_at = ? WHERE id = ?"

	now := time.Now()
	_, err := r.db.Exec(
		query,
		category.Name,
		now,
		category.ID,
	)

	category.UpdatedAt = now

	return err
}

func (r *categoryRepository) DeleteByID(id uint64) error {

	query := "DELETE FROM categories WHERE id = ?"
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
