package gopg

import (
	"fmt"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"log"
	"time"
)

type User struct {
	Id     int64
	Name   string
	Emails []string
}

func (u User) String() string {
	return fmt.Sprintf("User<%d %s %v>", u.Id, u.Name, u.Emails)
}

type Story struct {
	Id       int64
	Title    string
	AuthorId int64
	Author   *User
}

func NewDb() *pg.DB {
	return pg.Connect(&pg.Options{
		User: "postgres",
	})
}

func ShowQueryLine(db *pg.DB) {
	db.OnQueryProcessed(func(event *pg.QueryProcessedEvent) {
		query, err := event.FormattedQuery()
		if err != nil {
			panic(err)
		}

		log.Printf("[%s] %s", time.Since(event.StartTime), query)
	})
}

func (s Story) String() string {
	return fmt.Sprintf("Story<%d %s %s>", s.Id, s.Title, s.Author)
}

func PgExample() {
	u := &User{Id: 10086, Name: "Tom", Emails: []string{"a@a.com", "b@b.com"}}
	fmt.Println(u.String())

	s := &Story{Id: 10010, Title: "Golang In Action", AuthorId: u.Id, Author: u}
	fmt.Println(s.String())

	db := NewDb()
	defer db.Close()

	err := createSchema(db)
	if err != nil {
		panic(err)
	}

	err = db.Insert(u)
	if err != nil {
		panic(err)
	}

	u2 := &User{
		Name:   "admin",
		Emails: []string{"admin1@a.com", "admin2@a.com"},
	}
	err = db.Insert(u2)
	if err != nil {
		panic(err)
	}

	err = db.Insert(&User{
		Name:   "root",
		Emails: []string{"root1@a.com", "root2@a.com"},
	})
	if err != nil {
		panic(err)
	}

	err = db.Insert(s)
	if err != nil {
		panic(err)
	}

	/////////

	// Select user by primary key
	u3 := &User{Id: u.Id}
	err = db.Select(u3)
	if err != nil {
		panic(err)
	}
	fmt.Println(u3.String())

	// Select all users
	var users []User
	err = db.Model(&users).Select()
	if err != nil {
		panic(err)
	}
	for _, u := range users {
		fmt.Println(u.String())
	}

	// Select story and associate author in one query
	s2 := new(Story)
	err = db.Model(s2).
		Relation("Author").
		Where("story.id = ?", s.Id).
		Select()
	if err != nil {
		panic(err)
	}
	fmt.Println(s2.String())
}

func createSchema(db *pg.DB) error {
	for _, model := range []interface{}{(*User)(nil), (*Story)(nil)} {
		err := db.CreateTable(model, &orm.CreateTableOptions{
			Temp: true,
		})

		if err != nil {
			return err
		}
	}
	return nil
}
