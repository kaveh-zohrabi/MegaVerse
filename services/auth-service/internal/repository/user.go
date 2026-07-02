package repository

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID           string
	Email        string
	Name         string
	PasswordHash string
	Role         string
	Active       bool
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) FindByEmail(email string) (*User, error) {
	user := &User{}
	err := r.db.QueryRow(
		"SELECT id, email, name, password_hash, role, active, created_at, updated_at FROM users WHERE email = ?",
		email,
	).Scan(&user.ID, &user.Email, &user.Name, &user.PasswordHash, &user.Role, &user.Active, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) FindByID(id string) (*User, error) {
	user := &User{}
	err := r.db.QueryRow(
		"SELECT id, email, name, password_hash, role, active, created_at, updated_at FROM users WHERE id = ?",
		id,
	).Scan(&user.ID, &user.Email, &user.Name, &user.PasswordHash, &user.Role, &user.Active, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) Create(user *User) error {
	_, err := r.db.Exec(
		"INSERT INTO users (id, email, name, password_hash, role, active) VALUES (?, ?, ?, ?, ?, ?)",
		user.ID, user.Email, user.Name, user.PasswordHash, user.Role, user.Active,
	)
	return err
}

func (r *UserRepository) EmailExists(email string) (bool, error) {
	var count int
	err := r.db.QueryRow("SELECT COUNT(*) FROM users WHERE email = ?", email).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
