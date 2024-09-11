package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type Page struct {
	Page string `json :"page"`
}

type Words struct {
	Input string   `json :"input"`
	Words []string `json :"words"`
}

type Occur struct {
	Words map[string]int `json: "words"`
}

func main() {

	args := os.Args
	fmt.Println(args)

	if len(args) > 2 {
		fmt.Println("No args passed")
	}

	response, err := http.Get(args[1])
	anyError(err)

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	anyError(err)

	if response.StatusCode != 200 {
		fmt.Printf("Invalid output (HTTP Code %d): %s\n", response.StatusCode, string(body))
		os.Exit(1)
	}

	fmt.Println(string(body))

	fmt.Println("************Parsed using structs *****************")

	var page Page
	err = json.Unmarshal(body, &page)
	anyError(err)

	switch page.Page {
	case "words":
		var words Words
		err := json.Unmarshal(body, &words)
		anyError(err)
		fmt.Println("page : ", page.Page)
		fmt.Println("Words : ", words.Words)

	case "occurrence":
		var occur Occur
		err := json.Unmarshal(body, &occur)
		anyError(err)

		fmt.Println("Page : ", page.Page)
		fmt.Println("Words : ", occur.Words)

		fmt.Println("In map")
		for k, v := range occur.Words {
			fmt.Println(k, " : ", v)
		}
	default:
		fmt.Println("Page not found")
	}

}

func anyError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
