package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type NewTodo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	fmt.Println()
	fmt.Println("Using NewDecoder Instead of Unmarshal")

	myUrl := "https://jsonplaceholder.typicode.com/todos/1"

	res, err := http.Get(myUrl)
	if err != nil {
		log.Fatalf("error in Get :%v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Fatalf("error in response :%v", err)
	}

	//decoding with NewDecoder
	var myTodo NewTodo
	DecodeErr := json.NewDecoder(res.Body).Decode(&myTodo)
	if DecodeErr != nil {
		log.Fatalf("error in Decoding/unmarshalling  :%v", DecodeErr)
	}

	fmt.Println("Unmarshalled/Decoded json :", myTodo)
}
