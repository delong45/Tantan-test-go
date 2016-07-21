package user

import (
	"fmt"

	"gopkg.in/pg.v4"
)

type User struct {
	Id   int64
	Name string `json:"name"`
	Type string
}

var db *pg.DB

func createSchema() error {
	queries := []string{
		`CREATE TABLE users (id serial primary key, name text, type text)`,
		`CREATE TABLE relationships (id serial primary key, liked integer[], matched integer[], disliked integer[])`,
	}
	for _, q := range queries {
		_, err := db.Exec(q)
		if err != nil {
			return err
		}
	}
	return nil
}

func Init() error {
	db = pg.Connect(&pg.Options{
		User: "postgres",
	})

	return nil
}

func (u *User) String() string {
	return fmt.Sprintf("User<%s %s %s>", u.Id, u.Name, u.Type)
}

func (u *User) Create() error {
	if u.Name == "" {
		return fmt.Errorf("'name' is required")
	}
	u.Type = "user"

	err := db.Create(u)
	if err != nil {
		return err
	}

	return nil
}

func GetUsers() ([]User, error) {
	var users []User
	_, err := db.Query(&users, `SELECT * FROM users`)
	return users, err
}
