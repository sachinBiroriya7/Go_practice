package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	fmt.Println("encoding")
	Myurl := "https://jsonplaceholder.typicode.com/todos/1"

	parsedUrl, err := url.Parse(Myurl)
	if err != nil {
		fmt.Println("error in parsing ")
		return
	}
	fmt.Println("parsed url :", parsedUrl)

	res, err := http.Get(Myurl)
	if err != nil {
		fmt.Println("error in Request :", err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("error in Body  :", err)
		return
	}
	fmt.Println("content :", string(body))

	//un-marshall
	var todo1 Todo
	UnMarErr := json.Unmarshal(body, &todo1)
	if UnMarErr != nil {
		fmt.Println("UnMarshal error :", UnMarErr)
		return
	}
	fmt.Println("UnMarshalled Form :", todo1)

	//marshal
	var todo2 = Todo{
		UserId:    2,
		Id:        2,
		Title:     "task 2",
		Completed: false,
	}

	data, MarErr := json.Marshal(todo2)
	if MarErr != nil {
		fmt.Println("Marshal error :", MarErr)
		return
	}
	fmt.Println("Marshalled Form :", string(data))
}
