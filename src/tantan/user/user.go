package user

import (
	"crypto/md5"
	"fmt"
	"io"
	"math/rand"
	"strconv"
	"time"

	"gopkg.in/pg.v4"
)

type User struct {
	Id   int64
	Name string `json:"name"`
	Type string
}

type Users []User

var db *pg.DB

func createSchema(db *pg.DB) error {
	queries := []string{
		`CREATE TABLE users (id serial primary key, name text)`,
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

	err := createSchema(db)
	if err != nil {
		return err
	}
	return nil
}

func Md5(text string) string {
	m := md5.New()
	io.WriteString(m, text)
	return fmt.Sprintf("%x", m.Sum(nil))
}

func getUid() string {
	nano := time.Now().UnixNano()
	rand.Seed(nano)
	n := rand.Int63()
	id := Md5(Md5(strconv.FormatInt(nano, 10)) + Md5(strconv.FormatInt(n, 10)))
	return id
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

func (us *Users) List() error {
	return nil
}
