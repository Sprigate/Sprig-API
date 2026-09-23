package users

import (
	"database/sql"
	"errors"
	"sprig/internal/model"
	"time"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) AddUser(user *model.User) error {
	if user == nil {
		return errors.New("Data tidak boleh kosong!")
	}

	query := "INSERT INTO users (name, email, password, created_at, updated_at) VALUES (?, ?, ?, ?, ?)"

	now := time.Now()
	result, err := r.db.Exec(
		query,
		user.Name,
		user.Email,
		user.Password,
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

	user.ID = uint64(id)
	user.CreatedAt = now
	user.UpdatedAt = now

	return nil
}

func (r *userRepository) FindByID(id uint64) (*model.User, error) {
	if id <= 0 {
		return nil, errors.New("ID tidak boleh kosong")
	}

	query := "SELECT id, name, email, password, created_at FROM users WHERE id = ?"
	row := r.db.QueryRow(query, id)

	var user model.User

	err := row.Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("User tidak dapat ditemukan")
		}

		return nil, err
	}

	return &user, nil
}

func (r *userRepository) FindAll() ([]model.User, error) {
	query := "SELECT id, name, email, created_at FROM users"

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.User

	for rows.Next() {
		var user model.User

		err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *userRepository) Save(user *model.User) error {
	query := "UPDATE users SET name = ?, email = ?, password = ?, updated_at = ? WHERE id = ?"

	now := time.Now()
	_, err := r.db.Exec(
		query,
		user.Name,
		user.Email,
		user.Password,
		now,
		user.ID,
	)

	user.UpdatedAt = now

	return err
}

func (r *userRepository) DeleteByID(id uint64) error {
	if id <= 0 {
		return errors.New("ID tidak boleh kosong")
	}

	query := "DELETE FROM users WHERE id = ?"
	_, err := r.db.Exec(query, id)

	if err != nil {
		return err
	}

	return nil
}
