package gopg

import (
	"fmt"
	"github.com/go-pg/pg/orm"
	"log"
	"time"
)

type Author struct {
	Id       int
	Name     string
	Books    []*Book
	AvatarId int
}

type Book struct {
	Id        int
	Title     string
	AuthorID  int
	Author    Author // has one relation
	EditorID  int
	Editor    *Author   // has one relation
	CreatedAt time.Time `sql:"default:now()"`
	UpdatedAt time.Time
}

func QueryExmaple() {
	//example2()
	example1()
}

func example2() {
	db := NewDb()
	defer db.Close()
	ShowQueryLine(db)

	filter2 := func(q *orm.Query) (*orm.Query, error) {

		return q, nil
	}

	filter1 := func(q *orm.Query) (*orm.Query, error) {
		q.Apply(filter2)
		return q, nil
	}

	var bs []*Book
	q := db.Model(&bs)
	q.Apply(filter1)
	err := q.Select()
	if err != nil {
		panic(err)
	}

	for _, b := range bs {
		fmt.Println(b)
	}
}

func example1() {
	db := NewDb()
	defer db.Close()
	ShowQueryLine(db)

	//filter1 := func(q *orm.Query) (*orm.Query, error) {
	//	q = q.Where(fmt.Sprintf(`"%s"."id" = ?`, "author"), 1)
	//	return q, nil
	//}
	//
	//filter2 := func(q *orm.Query) (*orm.Query, error) {
	//	q.Column("author.*", "Books")
	//	return q, nil
	//}

	//var a Author
	//q := db.Model(&a)
	////q.Apply(filter1)
	////q.Apply(filter2)
	//q.Column("id").Column("Books")
	//q.Where(`"author"."id" = ?`, 5)
	//err := q.Select()
	//if err != nil {
	//	panic(err)
	//}
	//
	//fmt.Println(a)
	//if len(a.Books) > 0 {
	//	fmt.Println(a.Books[0])
	//}

	var authors []*Author
	var author Author
	err := db.Model(&author).
		Column("id").
		Column("name").
		Where("id = ?", 10).
		Select()

	if err != nil {
		log.Println(err)
		panic(err)
	}

	for _, a := range authors {
		fmt.Println(a)
	}

	fmt.Println(author)
}
