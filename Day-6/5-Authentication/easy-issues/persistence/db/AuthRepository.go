package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"github.com/PacktPublishing/Advanced-Go-Programming-in-7-Days/Day-6/5-Authentication/easy-issues/domain"
)

const (
	authDriver = "sqlite3"
	querySelectUserByEmail = "SELECT * FROM auth WHERE user_email=?"
	queryInsertNewUser = "INSERT INTO auth (user_uuid, user_email, user_status, user_password) VALUES (?, ?, ?, ?)"
	)

const authSchema = `CREATE TABLE IF NOT EXISTS auth (
	user_id integer primary key autoincrement,
	user_uuid text,
    user_email text,
    user_status text,
	user_password text);`

// UserRepository concrete implementation of in-memory db
type AuthRepository struct {
	db *sqlx.DB
}

func NewAuthRepository() *AuthRepository {
	db, err := sqlx.Open(authDriver, "test.db")
	if err != nil {
		log.Fatal("Failed to connect to database.")
	}
	r := &AuthRepository{
		db: db,
	}
	r.Init(authSchema)
	return r
}

func (r *AuthRepository)Init(schema string) error  {
	// execute a query on the server
	_, err := r.db.Exec(schema)
	return err
}

func (r *AuthRepository)GetRegistrationByEmail(email string) (*domain.UserRegistration, error)  {
	stmt, err := r.db.Preparex(querySelectUserByEmail)
	if err != nil {
		return nil, err
	}

	var u domain.UserRegistration
	err = stmt.Get(&u, email)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (r *AuthRepository)Create(u *domain.UserRegistration) error {
	stmt, err := r.db.Preparex(queryInsertNewUser)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(u.Uuid, u.Email,u.Status, u.PasswordHash)
	if err != nil {
		return err
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	u.Id = lastId
	return nil
}