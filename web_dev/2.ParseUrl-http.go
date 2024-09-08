package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

func anyError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	fmt.Println("PASSING URL THROUGH CMD")

	args := os.Args
	if len(args) > 2 {
		fmt.Println("no args passed")
		os.Exit(1)
	}

	fmt.Print("passed url is :", args[1])

	//parsing url
	parse, err := url.ParseRequestURI(args[1])

	if err != nil {
		log.Fatal(err) // this is equivaltent to -> fmt.Println(err), os.Exit(1)
	}
	fmt.Println(parse)

	/****************************Making http.Get()*********************************/

	resp, err := http.Get(args[1]) //get takes a string a parameter and returns a response
	anyError(err)

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	anyError(err)

	fmt.Printf("status code : %v, and is content is %v", resp.StatusCode, string(body))
}

//go run .\2.ParseUrl-http.go http://localhost:8080/words

/*
PASSING URL THROUGH CMD
passed url is :http://localhost:8080/wordshttp://localhost:8080/words
status code : %v, and is content is %v 200 {"page":"words","input":"","words":[]}

*/
/******************************************************************************************/

/*parsing url ->  breaking down a web address into its individual parts so that a computer can understand
and use it to access a specific resource on the internet*/

//takes raw url and return parsed structured url,  Represents the parsed structure of a URL  and error if any

//body is streamed , and taken when needed
