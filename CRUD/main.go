package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

type Todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	fmt.Println()

	myUrl := "https://jsonplaceholder.typicode.com/todos"

	makeGetRequest(myUrl + "/1")
	makePostRequest(myUrl)
	makePutRequest(myUrl + "/2")
	makeDeleteRequest(myUrl + "/1")
}

func makeGetRequest(req string) {

	fmt.Println()
	fmt.Println("In Get req")
	res, err := http.Get(req)
	if err != nil {
		log.Fatalf("error in get request :%v", err)
	}
	defer res.Body.Close()

	var myTodo Todo
	DecodeErr := json.NewDecoder(res.Body).Decode(&myTodo)
	if DecodeErr != nil {
		log.Fatalf("Unmarshall error :%v", DecodeErr)
	}
	fmt.Println("Unmarshalled json :", myTodo)
}

func makePostRequest(url string) {

	fmt.Println()
	fmt.Println("In Post req")
	myTodo := Todo{
		UserId:    209,
		Id:        121,
		Title:     "myTodo",
		Completed: true,
	}

	postedData, err := json.Marshal(myTodo)
	if err != nil {
		log.Fatalf("Marshalled/Encoded error :%v", err)
	}

	fmt.Println("Marshalled/Encoded data :", postedData)

	//converting []byte after marshalling to  *strings.Reader type to pass in POST req
	jsonReader := strings.NewReader(string(postedData))

	res, err := http.Post(url, "application/json", jsonReader)
	if err != nil {
		log.Fatalf("POST req error :%v", err)
	}
	defer res.Body.Close()

	fmt.Println("Posted data :", string(postedData))
	//fmt.Println("Posted data :", jsonReader)

}

func makePutRequest(url string) {

	fmt.Println()
	fmt.Println("In PUT req")
	myTodo := Todo{
		UserId:    21223,
		Id:        2,
		Title:     "updated title",
		Completed: false,
	}
	marshalledJson, err := json.Marshal(myTodo)
	if err != nil {
		log.Fatalf("error in marshalling :%v", err)
	}
	fmt.Println("encoded json :", string(marshalledJson))

	jsonReader := strings.NewReader(string(marshalledJson))
	req, err := http.NewRequest("PUT", url, jsonReader)
	if err != nil {
		log.Fatalf("error in making req :%v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("error in response :%v", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("error in reading response Body :%v", err)
	}
	fmt.Println(res.Status)
	fmt.Println("Sent Json :", string(body))
}

func makeDeleteRequest(url string) {

	fmt.Println()
	fmt.Println("In Delete req")

	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		log.Fatalf("Request error :%v", err)
	}

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("Response error :%v", err)
	}
	res.Body.Close()

	body, _ := io.ReadAll(res.Body)
	fmt.Println("Response :", body)
	fmt.Println("Status :", res.Status)

}
