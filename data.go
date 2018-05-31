package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

type PostJson struct {
	Id       int       `json:"id"`
	Content  string    `json:"content"`
	Author   Author    `json:"author"`
	Comments []Comment `json:"comments"`
}

type Author struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Comment struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func parseJson() {
	jsonFile, err := os.Open("template/post.json")
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return
	}

	defer jsonFile.Close()

	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("Error reading JSON data:", err)
		return
	}

	var post PostJson
	json.Unmarshal(jsonData, &post)
	fmt.Println(post)

}

func decodeJson() {
	jsonFile, err := os.Open("template/post.json")
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return
	}

	defer jsonFile.Close()

	decoder := json.NewDecoder(jsonFile)
	for {
		var post PostJson
		err := decoder.Decode(&post)
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println("Error decoding JSON:", err)
			return
		}
		fmt.Println(post)
	}
}

func createJson() {
	post := PostJson{
		Id:      1,
		Content: "Hello world",
		Author: Author{
			Id:   2,
			Name: "Sau Sheong",
		},
		Comments: []Comment{
			{
				Id:      3,
				Content: "Have a great day!",
				Author:  "Adman",
			},
			{
				Id:      4,
				Content: "How are you today?",
				Author:  "Betty",
			},
		},
	}

	output, err := json.MarshalIndent(&post, "", "\t")
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}

	err = ioutil.WriteFile("template/post2.json", output, 0644)
	if err != nil {
		fmt.Println("Error writting JSON to file:", err)
		return
	}
}

func encodeJson() {
	post := PostJson{
		Id:      1,
		Content: "Hello world",
		Author: Author{
			Id:   2,
			Name: "Sau Sheong",
		},
		Comments: []Comment{
			{
				Id:      3,
				Content: "Have a great day!",
				Author:  "Adman",
			},
			{
				Id:      4,
				Content: "How are you today?",
				Author:  "Betty",
			},
		},
	}

	jsonFile, err := os.Create("template/post3.json")
	if err != nil {
		fmt.Println("Error creating JSON file:", err)
		return
	}

	encoder := json.NewEncoder(jsonFile)
	err = encoder.Encode(&post)
	if err != nil {
		fmt.Println("Error encoding JSON to file:", err)
	}
}

//
type PostXml struct {
	XMLName  xml.Name     `xml:"post"`
	Id       string       `xml:"id,attr"`
	Content  string       `xml:"content"`
	Author   AuthorXml    `xml:"author"`
	Xml      string       `xml:",innerxml"`
	Comments []CommentXml `xml:"comments>comment"`
}

type AuthorXml struct {
	Id   string `xml:"id,attr"`
	Name string `xml:",chardata"`
}

type CommentXml struct {
	Id      string    `xml:"id,attr"`
	Content string    `xml:"content"`
	Author  AuthorXml `xml:"author"`
}

func parseXml() {
	xmlFile, err := os.Open("template/post.xml")
	if err != nil {
		fmt.Println("Error open XML file:", err)
		return
	}

	defer xmlFile.Close()

	xmlData, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		fmt.Println("Error reading XML file:", err)
		return
	}

	var post PostXml
	xml.Unmarshal(xmlData, &post)
	fmt.Println(post)
}

func decodeXml() {
	xmlFile, err := os.Open("template/post.xml")
	if err != nil {
		fmt.Println("Error opening XML file:", err)
		return
	}
	defer xmlFile.Close()

	decoder := xml.NewDecoder(xmlFile)
	for {
		t, err := decoder.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error decoding XML into tokens:", err)
			return
		}

		switch se := t.(type) {
		case xml.StartElement:
			if se.Name.Local == "comment" {
				var comment CommentXml
				decoder.DecodeElement(&comment, &se)
				fmt.Println(comment)
			}
		}
	}
}

//
func createXml() {
	post := PostXml{
		Id:      "1",
		Content: "Hello World!",
		Author: AuthorXml{
			Id:   "2",
			Name: "Sau Shenong",
		},
		Comments: []CommentXml{
			{
				Id: "1",
				Author: AuthorXml{
					Id:   "1",
					Name: "Betty",
				},
				Content: "Have a grate day!",
			},
		},
	}

	output, err := xml.MarshalIndent(&post, "", "\t")
	if err != nil {
		fmt.Println("Error marshalling to XML:", err)
		return
	}

	err = ioutil.WriteFile("template/post2.xml", []byte(xml.Header+string(output)), 0644)
	if err != nil {
		fmt.Println("Error Writing XML to file:", err)
		return
	}
}

func encodeXml() {
	post := PostXml{
		Id:      "1",
		Content: "Hello World!",
		Author: AuthorXml{
			Id:   "2",
			Name: "Sau Shenong",
		},
		Comments: []CommentXml{
			{
				Id: "1",
				Author: AuthorXml{
					Id:   "1",
					Name: "Betty",
				},
				Content: "Have a grate day!",
			},
		},
	}

	xmlFile, err := os.Create("template/post3.xml")
	if err != nil {
		fmt.Println("Error creating XML file:", err)
		return
	}

	encoder := xml.NewEncoder(xmlFile)
	encoder.Indent("", "\t")
	err = encoder.Encode(&post)
	if err != nil {
		fmt.Println("Error encoding XML to file:", err)
		return
	}
}
