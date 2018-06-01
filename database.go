package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"fmt"
)

type PostDB struct {
	Id int
	Content string
	Author	string
}

//
var Db *sql.DB

//
func connectDB() {
	var err error

	Db, err = sql.Open("postgres", "user=postgres dbname=postgres password=postgres sslmode=disable")
	if err != nil {
		panic(err)
	}
}

//
func Posts(limit int)(posts []PostDB, err error) {
	rows, err := Db.Query("SELECT id, content, author FROM posts LIMIT $1", limit)
	if err != nil {
		return
	}

	for rows.Next() {
		post := PostDB{}
		err = rows.Scan(&post.Id, &post.Content, &post.Author)
		if err != nil {
			return
		}
		posts = append(posts, post)
	}
	rows.Close()
	return
}

func (post *PostDB) Create() (err error) {
	statement := "INSERT INTO posts (content, author) values ($1, $2) RETURNING id"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(post.Content, post.Author).Scan(&post.Id)
	return
}

func (post *PostDB)Delete() (err error) {
	_, err = Db.Exec("DELETE FROM posts where id = $1", post.Id)
	return
}

//
func GetPost(id int) (post PostDB, err error) {
	post = PostDB{}
	err = Db.QueryRow("SELECT id, content, author FROM posts WHERE id = $1", id).Scan(&post.Id, &post.Content, &post.Author)
	return
}

func (post *PostDB) Update() (err error) {
	_, err = Db.Exec("UPDATE posts SET content = $2, author = $3 WHERE id = $1", post.Id, post.Content, post.Author)
	return
}

func databaseExample() {
	post := PostDB{
		Content:"Hello World!",
		Author:"Sau Sheong",
	}

	fmt.Println(post)
	err := post.Create()
	if err != nil {
		fmt.Println("Create Error:", err)
		return
	}
	fmt.Println(post)

	readPost, err:= GetPost(post.Id)
	if err != nil {
		fmt.Println("Get Error:", err)
		return
	}
	fmt.Println(readPost)

	readPost.Content = "Bonjour Monde!"
	readPost.Author = "Pierre"
	err = readPost.Update()
	if err != nil {
		fmt.Println("Update Error:", err)
		return
	}

	posts, err := Posts(post.Id)
	fmt.Println("A", &err)
	if err != nil {
		fmt.Println("Posts Error:", err)
		return
	}else {
		fmt.Println("B:", &err)
	}
	fmt.Println(posts)

	err = readPost.Delete()
	if err != nil {
		fmt.Println("Delete Error:", err)
		return
	}
}