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

/*
	//reading response body stream
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error in response Body :%v", err)
	}
	fmt.Println("Content :", string(body))

	//unmarshalling json to struct
	var myTodo Todo
	json.Unmarshal(body, &myTodo)
	fmt.Println("our data Struct :", myTodo)
*/

// this can be achieved directly by newDecoder Fuction

/*newDecoder provides more flexibility and is suitable for decoding large JSON streams,
while Unmarshal is a simpler option for decoding JSON data from memory.

In fewer steps You can get data in our struct.
avoids reading whole stream of body and then unmarshalling it to our structs


it directly reads response body and decodes into our structs
*/
