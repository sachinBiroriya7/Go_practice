package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	fmt.Println()

	myUrl := "https://jsonplaceholder.typicode.com/todos/1"

	resp, err := http.Get(myUrl)
	if err != nil {
		log.Fatalf("Get req err: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Error in Getting response :%v", err)
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

	var myTodo Todo
	DecodeErr := json.NewDecoder(resp.Body).Decode(&myTodo)
	if DecodeErr != nil {
		log.Fatalf("Error in Unmarshalling Body :%v", DecodeErr)
	}
	fmt.Println(myTodo)
}

/*newDecoder provides more flexibility and is suitable for decoding large JSON streams,
while Unmarshal is a simpler option for decoding JSON data from memory.

In fewer steps You can get data in our struct.
avoids reading whole stream of body and then unmarshalling it to our structs


it directly reads response body and decodes into our structs
*/
