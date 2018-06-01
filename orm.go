package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type PostOrm struct {
	Id         int
	Content    string
	AuthorName string `db:"author"`
}

var DbOrm *sqlx.DB

func initSqlx() {
	var err error
	DbOrm, err = sqlx.Open("postgres", "user=postgres dbname=postgres password=postgres sslmode=disable")
	if err != nil {
		panic(err)
	}
}

func GetPostOrm(id int) (post PostOrm, err error) {
	post = PostOrm{}

	err = DbOrm.QueryRowx("SELECT id, content, author FROM posts WHERE id = $1", id).StructScan(&post)
	if err != nil {
		return
	}

	return
}

func (post *PostOrm) Create() (err error) {
	err = DbOrm.QueryRow("INSERT INTO posts (content, author) VALUES ($1, $2) RETURNING id",
		post.Content, post.AuthorName).Scan(&post.Id)
	return
}

func ormExample() {
	post := PostOrm{
		Content:    "Hello World!",
		AuthorName: "Sau Sheong",
	}

	err := post.Create()
	if err != nil {
		fmt.Println("Create post Error:", err)
		return
	}
	fmt.Println(post)

	readPost, err := GetPostOrm(post.Id)
	if err != nil {
		fmt.Println("GetPostOrm Error:", err)
		return
	}
	fmt.Println(readPost)
}
