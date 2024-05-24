package domain

import (
	"context"
	"database/sql"

	// "time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type USER struct {
	ID        uuid.UUID `json:"id" db:"id"`
	FirstName string    `json:"first_name" db:"first_name"`
	LastName  string    `json:"last_name" db:"last_name"`
	Email     string    `json:"email" db:"email"`
	// CreatedAt time.Time `json:"created_at" db:"created_at"`
}

func InsertUser(ctx context.Context, pool *pgxpool.Pool, user USER) error {
	query := `INSERT INTO tb_user (id, first_name, last_name, email) VALUES ($1, $2, $3, $4)`
	_, err := pool.Exec(ctx, query, user.ID, user.FirstName, user.LastName, user.Email)
	return err
}

func CreateUser(db *sql.DB, user USER) error {
	query := `INSERT INTO tb_user (id, first_name, last_name, email) VALUES ($1, $2, $3, $4)`
	_, err := db.Exec(query, user.ID, user.FirstName, user.LastName, user.Email)
	return err
}

func GetUser(db *sql.DB, id uuid.UUID) (USER, error) {
	var user USER
	query := `SELECT id, first_name, last_name, email FROM tb_user WHERE id = $1`
	err := db.QueryRow(query, id).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email)
	return user, err
}

func UpdateUser(db *sql.DB, user USER) error {
	query := `UPDATE tb_user SET first_name = $1, last_name = $2, email = $3 WHERE id = $4`
	_, err := db.Exec(query, user.FirstName, user.LastName, user.Email, user.ID)
	return err
}
