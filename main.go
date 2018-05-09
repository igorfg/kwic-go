package main

import "fmt"
import "github.com/igorfg/kwic-go/kwic"
import "log"

func main() {
	file_name := "papers.txt"
	lines, err := kwic.Input(file_name)

	if err != nil {
		log.Fatalf("Error in input(%v): %v", file_name, err)
	}

	for i := 0; i < len(lines); i++ {
		fmt.Println(lines[i])
	}

	fmt.Println("chegou")
}
