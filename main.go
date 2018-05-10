package main

import "fmt"
import "github.com/igorfg/kwic-go/kwic"

func main() {
	var fsm kwic.DataStorageManager = &kwic.FileBasedStorageManager{}

	fsm.Init()

	for i := 0; i < fsm.Length(); i++ {
		fmt.Println(fsm.Line(i))
	}

	fmt.Println("Fim da Execução")
}
