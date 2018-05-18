package main

import "github.com/igorfg/kwic-go/kwic"

func main() {
	var (
		storageManager kwic.DataStorageManager
		indexManager   kwic.IndexManager
		outPutManager  kwic.OutputManager
	)

	cliInterface := new(kwic.CommandLineInterface)
	cliInterface.Init(storageManager, indexManager, outPutManager)
}
