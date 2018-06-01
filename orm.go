package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/jmoiron/sqlx"
	"time"
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

/////////
type PostGorm struct {
	Id       int
	Content  string
	Author   string `sql:"not null"`
	Comments []CommentGorm
	CreateAt time.Time
}

type CommentGorm struct {
	Id       int
	Content  string
	Author   string `sql:"not null"`
	PostId   int
	CreateAt time.Time
}

var DbGrom *gorm.DB

func initGorm() {
	var err error
	DbGrom, err = gorm.Open("postgres", "user=postgres dbname=postgres password=postgres sslmode=disable")
	if err != nil {
		panic(err)
	}
	DbGrom.AutoMigrate(&PostGorm{}, &CommentGorm{})
}

func ormExample2() {
	post := PostGorm{
		Content: "Hello World!",
		Author:  "Sau Sheong",
	}
	fmt.Println(post)
	DbGrom.Create(&post)
	fmt.Println(post)

	comment := CommentGorm{
		Content: "God post!",
		Author:  "Joe",
	}
	DbGrom.Model(&post).Association("Comments").Append(comment)

	var readPost PostGorm
	DbGrom.Where("author = $1", "Sau Sheong").First(&readPost)
	var comments []CommentGorm
	DbGrom.Model(&readPost).Related(&comments)
	fmt.Println(comments)
}
