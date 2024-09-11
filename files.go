package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("creating file")
	file, err := os.Create("myFile.txt")
	if err != nil {
		fmt.Println("err in creating file", err)
		return
	}
	defer file.Close()

	// writing in file

	content := "this string is going to be written in the file"

	bytes_written, err := io.WriteString(file, content)
	if err != nil {
		fmt.Println("err in Writing file :", err)
		return
	}
	fmt.Println("bytes written :", bytes_written)

	//reading a file

	//1. using buffer
	buffer := make([]byte, 1024)

	file, err = os.Open("myFile.txt") //os.OpenFile with os.O_WRONLY | os.O_CREATE
	if err != nil {
		fmt.Println("err in Opening file", err)
		return
	}
	defer file.Close()

	for {
		n, err := file.Read(buffer)
		if err == io.EOF {
			fmt.Println("file end reached")
			break
		}
		if err != nil {
			fmt.Println("err in Reading file", err)
			return
		}

		fmt.Println(string(buffer[:n]))

	}

	//using readAll func -> loads the complete content in memory, not suitable for large files.

	n, err := os.ReadFile("myFile.txt")
	if err != nil {
		fmt.Println("err in Reading file", err)
		return
	}
	fmt.Println(string(n))
}
