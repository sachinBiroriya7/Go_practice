package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

type Page struct {
	Name string `json:"page"`
}

type Response interface {
	GetResponse() string
}

type Words struct {
	Input string   `json:"input"`
	Words []string `json:"words"`
}

func (w Words) GetResponse() string {
	return fmt.Sprintf("Words: %s", strings.Join(w.Words, ", "))
}

type Occur struct {
	Words map[string]int `json:"words"`
}

func (o Occur) GetResponse() string {
	words := []string{}
	for word, occurrence := range o.Words {
		words = append(words, fmt.Sprintf("%s (%d)", word, occurrence))
	}
	return fmt.Sprintf("Words: %s", strings.Join(words, ", "))
}

func main() {
	fmt.Println("in main")

	args := os.Args
	if len(args) < 2 {
		fmt.Println("no arguments passed")
		os.Exit(1)
	}
	fmt.Println("passed url is :", args[1])

	res, err := doRequest(args[1])
	if err != nil {
		log.Fatal("get request error : \n", err)
	}

	fmt.Printf("Response is: %s\n", res.GetResponse())
}

func doRequest(req string) (Response, error) { //returning Response interface that could be pages/occurence or anything
	res, err := http.Get(req)
	if err != nil {
		return nil, fmt.Errorf("get request error: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("no response: %v", err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("read body error: %v", err)
	}
	fmt.Printf("Response successful %v and the content is %v\n", res.StatusCode, string(body))

	var page Page
	err = json.Unmarshal(body, &page)
	if err != nil {
		return nil, fmt.Errorf("unmarshal error: %v", err)
	}

	switch page.Name {
	case "words":
		var words Words
		err := json.Unmarshal(body, &words)
		if err != nil {
			return nil, fmt.Errorf("Words Unmarshal error :  %v\n", err)
		}
		return words, nil

	case "occurrence":
		var occur Occur
		err := json.Unmarshal(body, &occur)
		if err != nil {
			return nil, fmt.Errorf("Occurrence unmarshal error,  %v/n", err)
		}
		return occur, nil
	}
	return nil, nil

}
