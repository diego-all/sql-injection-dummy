package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"time"
)

const dbTimeout = time.Second * 3

var db *sql.DB

func New(dbPool *sql.DB) Models {
	db = dbPool

	return Models{
		User: User{},
	}
}

type Models struct {
	User User
}

type User struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	FirstName string    `json:"first_name,omitempty"`
	LastName  string    `json:"last_name,omitempty"`
	Password  string    `json:"password"`
	Active    int       `json:"active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *User) DeleteUser(id string) error {

	// Exec uses context.Background internally; to specify the context, use ExecContext.
	// ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	// defer cancel()

	stmt := "DELETE FROM users WHERE id = $1"

	//ExecContext es una y Exec es otro
	_, err := db.Exec(stmt, id)
	if err != nil {
		return err
	}

	return nil
}

func (u *User) DeleteUserSQLi(id string) error {
	//ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	//defer cancel()

	query := fmt.Sprintf("DELETE FROM users WHERE id = '%s'", id)

	fmt.Println(query)

	if _, err := db.Exec(query); err != nil {
		log.Fatalln("Couldn't delete", err)
	}

	return nil
}

func (u *User) GetUserByIDSQLi(id string) (*User, error) {

	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// Sustituir por query vulnerable con sprintf

	query := `select id, email, first_name, last_name, password, user_active, created_at, updated_at from users where email = $1`

	var user User
	row := db.QueryRowContext(ctx, query, id)

	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.Password,
		&user.Active,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
