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

type Relation struct {
	User_id string `json:"user_id"`
	State   string `json:"state"`
	Type    string `json:"type"`
}

type Relationships struct {
	Id       string
	Other_id string
	State    string
}

type State struct {
	State string `json:"state"`
}

var db *pg.DB

func createSchema() error {
	queries := []string{
		`CREATE TABLE users (id serial primary key, name text, "type" text)`,
		`CREATE TABLE relationships (id text, other_id text, state text)`,
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

func GetRelations(id string) ([]Relation, error) {
	var relations []Relation
	var relationships []Relationships
	_, err := db.Query(&relationships, `SELECT * FROM relationships WHERE id = ?`, id)
	if err != nil {
		return relations, err
	}

	for _, r := range relationships {
		rel := Relation{
			User_id: r.Other_id,
			State:   r.State,
			Type:    "relationship",
		}
		relations = append(relations, rel)
	}
	return relations, err
}

func GreateRelation(id, other_id, state string) (Relation, error) {
	r := Relation{
		User_id: other_id,
		State:   state,
		Type:    "relationship",
	}

	if state != "liked" && state != "disliked" {
		return r, fmt.Errorf("no such relation state")
	}

	relationships := Relationships{
		Id:       id,
		Other_id: other_id,
		State:    state,
	}

	var preRelation Relationships
	err := db.Model(&preRelation).Where("id = ?", id).Where("other_id = ?", other_id).Select()
	if err != nil {
		return r, err
	}
	if preRelation.State == state {
		return r, nil
	}
	if preRelation.State == "" {
		db.Create(&relationships)
	}

	var corRelation Relationships
	err = db.Model(&corRelation).Where("id = ?", other_id).Where("other_id = ?", id).Select()
	if err != nil {
		return r, err
	}
	corState := corRelation.State

	switch state {
	case "liked":
		if corState == "matched" {
			r.State = "matched"
		} else if corState == "liked" {
			r.State = "matched"
			relationships.State = "matched"
			corRelation.State = "matched"
			db.Model(&relationships).Column("state").Where("id = ?", id).Where("other_id = ?", other_id).Update()
			db.Model(&corRelation).Column("state").Where("id = ?", other_id).Where("other_id = ?", id).Update()
		} else {
			db.Model(&relationships).Column("state").Where("id = ?", id).Where("other_id = ?", other_id).Update()
		}
	case "disliked":
		if corState == "matched" {
			corRelation.State = "liked"
			db.Model(&corRelation).Column("state").Where("id = ?", other_id).Where("other_id = ?", id).Update()
		}
		db.Model(&relationships).Column("state").Where("id = ?", id).Where("other_id = ?", other_id).Update()
	}

	return r, nil
}
