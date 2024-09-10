package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

// {"page":"words","input":"","words":[]}
type Words struct {
	Page  string   `json:"page"`
	Input string   `json:"input"`
	Words []string `json: "Words"`
}

func main() {

	ParsedArgs := parseUrl()

	/****************************Making http.Get()*********************************/

	respCode, Respbody := httpGetReq(ParsedArgs)

	/****************************Parse Json *********************************/

	if respCode != 200 {
		fmt.Printf("Response code is not %v and Body is : %s", respCode, Respbody)
		os.Exit(1)
	}

	var words Words
	err := json.Unmarshal(Respbody, &words)
	anyError(err)

	fmt.Printf("JSON parsed\n Page:%v \n Words:%v\n", words.Page, words.Words)

}

// *********************************************************************************************//
func httpGetReq(ParsedArgs string) (int, []byte) {
	resp, err := http.Get(ParsedArgs) //get takes a string a parameter and returns a response
	anyError(err)

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	anyError(err)

	fmt.Printf("status code : %v, and the content is %v\n", resp.StatusCode, string(body))

	return resp.StatusCode, body
}

func parseUrl() string {
	fmt.Println("PASSING URL THROUGH CMD")

	args := os.Args
	if len(args) > 2 {
		fmt.Println("no args passed")
		os.Exit(1)
	}

	fmt.Println("passed url is :", args[1])

	//parsing url
	parse, err := url.ParseRequestURI(args[1])

	if err != nil {
		log.Fatal(err) // this is equivalent to -> fmt.Println(err), os.Exit(1)
	}
	fmt.Println(parse)
	return args[1]

}

func anyError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
