package domain

import (
	"database/sql"
	// "time"

	"github.com/google/uuid"
)

type USER struct {
	ID        uuid.UUID `json:"id" db:"id"`
	FirstName string    `json:"first_name" db:"first_name"`
	LastName  string    `json:"last_name" db:"last_name"`
	Email     string    `json:"email" db:"email"`
	// CreatedAt time.Time `json:"created_at" db:"created_at"`
}

func InsertUser(db *sql.DB, user USER) error {
	query := `INSERT INTO db_test (id, first_name, last_name, email) VALUES ($1, $2, $3, $4)`
	_, err := db.Exec(query, user.ID, user.FirstName, user.LastName, user.Email)
	return err
}

func CreateUser(db *sql.DB, user USER) error {
	query := `INSERT INTO db_test (id, first_name, last_name, email) VALUES ($1, $2, $3, $4)`
	_, err := db.Exec(query, user.ID, user.FirstName, user.LastName, user.Email)
	return err
}

func GetUser(db *sql.DB, id uuid.UUID) (USER, error) {
	var user USER
	query := `SELECT id, first_name, last_name, email FROM db_test WHERE id = $1`
	err := db.QueryRow(query, id).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email)
	return user, err
}

func UpdateUser(db *sql.DB, user USER) error {
	query := `UPDATE db_test SET first_name = $1, last_name = $2, email = $3 WHERE id = $4`
	_, err := db.Exec(query, user.FirstName, user.LastName, user.Email, user.ID)
	return err
}
