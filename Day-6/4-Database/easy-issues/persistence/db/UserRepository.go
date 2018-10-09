package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"github.com/PacktPublishing/Advanced-Go-Programming-in-7-Days/Day-6/4-Database/easy-issues/domain"
)

const (
	usersDriver = "sqlite3"
	querySelectAllUsers = "SELECT * FROM users"
	querySelectUser = "SELECT * FROM users WHERE user_id=?"
	queryInsertUser = "INSERT INTO users (user_name, user_surname, user_email) VALUES (?, ?, ?)"
	queryDeleteUser = "DELETE FROM users WHERE user_id=? LIMIT 1"
	)

const userSchema = `CREATE TABLE IF NOT EXISTS users (
	user_id integer primary key autoincrement,
    user_name text,
	user_surname text,
    user_email text);`

// UserRepository concrete implementation of in-memory db
type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository() *UserRepository {
	db, err := sqlx.Open(usersDriver, ":memory:")
	if err != nil {
		log.Fatal("Failed to connect to database.")
	}
	r := &UserRepository{
		db: db,
	}
	r.Init(userSchema)
	return r
}

func (r *UserRepository)Init(schema string) error  {
	// execute a query on the server
	_, err := r.db.Exec(schema)
	return err
}

func (r *UserRepository)All() ([]*domain.User, error)  {
	rows, err := r.db.Queryx(querySelectAllUsers)
	if err != nil {
		return nil, err
	}
	// iterate over each row
	users := make([]*domain.User, 0, 0)
	for rows.Next() {
		var u domain.User
		err = rows.StructScan(&u)
		if err != nil {
			return nil, err
		}
		users = append(users, &u)
	}

	return users, nil
}

func (r *UserRepository)GetById(id int64) (*domain.User, error)  {
	stmt, err := r.db.Preparex(querySelectUser)
	if err != nil {
		return nil, err
	}

	var u domain.User
	err = stmt.Get(&u, id)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (r *UserRepository)Create(u *domain.User) error {
	stmt, err := r.db.Preparex(queryInsertUser)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(u.Name, u.Surname, u.Email)
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

func (r *UserRepository)Delete(id int64) error {
	stmt, err := r.db.Preparex(queryDeleteUser)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}
