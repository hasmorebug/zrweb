package main

import (
	"bytes"
	"encoding/csv"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type PostMemory struct {
	Id      int
	Content string
	Author  string
}

var PostById map[int]*PostMemory
var PostsByAuthor map[string][]*PostMemory

func store(post PostMemory) {
	PostById[post.Id] = &post
	PostsByAuthor[post.Author] = append(PostsByAuthor[post.Author], &post)
}

func storeInMemory() {
	PostById = make(map[int]*PostMemory)
	PostsByAuthor = make(map[string][]*PostMemory)

	post1 := PostMemory{Id: 1, Content: "Hello World!", Author: "Sau Sheong"}
	post2 := PostMemory{Id: 2, Content: "Bonjour Monde!", Author: "Pierre"}
	post3 := PostMemory{Id: 3, Content: "HHola Mundo!", Author: "Pedro"}
	post4 := PostMemory{Id: 4, Content: "Greeting Earthlings!", Author: "Sau Sheong"}

	store(post1)
	store(post2)
	store(post3)
	store(post4)

	fmt.Println(PostById[1])
	fmt.Println(PostById[2])

	for _, post := range PostsByAuthor["Sau Sheong"] {
		fmt.Println(post)
	}

	for _, post := range PostsByAuthor["Pedro"] {
		fmt.Println(post)
	}
}

/////////
func storeInFile() {
	data := []byte("Hello World!")
	err := ioutil.WriteFile("template/data1", data, 0644)
	if err != nil {
		fmt.Println("Error WriteFile:", err)
		return
	}

	read, err := ioutil.ReadFile("template/data1")
	if err != nil {
		fmt.Println("Error ReadFile:", err)
		return
	}
	fmt.Println(string(read))

	file, err := os.Create("template/data2")
	if err != nil {
		fmt.Println("Create Error:", err)
		return
	}
	defer file.Close()

	bytes, err := file.Write(data)
	if err != nil {
		fmt.Println("WriteFile Error:", err)
		return
	}
	fmt.Printf("Wrote %d bytes to file\n", bytes)

	file2, err := os.Open("template/data2")
	if err != nil {
		fmt.Println("Error Open:", err)
		return
	}
	defer file2.Close()

	read2 := make([]byte, len(data))
	bytes, err = file2.Read(read2)
	if err != nil {
		fmt.Println("Read Error:", err)
		return
	}
	fmt.Printf("Read % bytes from file\n", bytes)
	fmt.Println(string(read2))

}

//
func storeInCSV() {
	csvFile, err := os.Create("template/post.csv")
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()

	allPosts := []PostMemory{
		{Id: 1, Content: "Hello World!", Author: "Sau Sheong"},
		{Id: 2, Content: "Bonjour Monde!", Author: "Pierre"},
		{Id: 3, Content: "HHola Mundo!", Author: "Pedro"},
		{Id: 4, Content: "Greeting Earthlings!", Author: "Sau Sheong"},
	}

	writer := csv.NewWriter(csvFile)
	for _, post := range allPosts {
		line := []string{strconv.Itoa(post.Id), post.Content, post.Author}
		err := writer.Write(line)
		if err != nil {
			panic(err)
		}
	}
	writer.Flush()

	file2, err := os.Open("template/post.csv")
	if err != nil {
		panic(err)
	}
	defer file2.Close()

	reader := csv.NewReader(file2)
	reader.FieldsPerRecord = -1
	record, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	var posts []PostMemory
	for _, item := range record {
		id, err := strconv.Atoi(item[0])
		if err != nil {
			fmt.Println("ParseInt Error:", err)
		}

		post := PostMemory{Id: id, Content: item[1], Author: item[2]}
		posts = append(posts, post)
	}
	fmt.Println(posts[0].Id)
	fmt.Println(posts[0].Content)
	fmt.Println(posts[0].Author)
}

func storeGob(data interface{}, filename string) {
	buffer := new(bytes.Buffer)
	encoder := gob.NewEncoder(buffer)
	err := encoder.Encode(data)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(filename, buffer.Bytes(), 0600)
	if err != nil {
		panic(err)
	}
}

func loadGob(data interface{}, filename string) {
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	buffer := bytes.NewBuffer(raw)
	decoder := gob.NewDecoder(buffer)
	err = decoder.Decode(data)
	if err != nil {
		panic(err)
	}

}

func storeInGob() {
	post := PostMemory{Id: 1, Content: "Hello World!", Author: "Sau Sheong"}
	storeGob(post, "template/post.gob")
	var postRead PostMemory
	loadGob(&postRead, "template/post.gob")
	fmt.Println(postRead)
}

//
func storeDataExample() {
	storeInGob()
	//storeInCSV()
	//storeInFile()
	//storeInMemory()
}
