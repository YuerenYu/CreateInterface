package main

import (
	"CreateInterface/src/database"
	"CreateInterface/src/sys"
	"CreateInterface/src/utils"
	_ "github.com/mattn/go-sqlite3"
	"net/http"
	"sync"
)

func main() {
	utils.PrintWelcome()
	SINGLE := make(chan *http.Server, 1)
	dataChan := make(chan database.InterfaceInfo)
	wg := sync.WaitGroup{}

	go sys.Scanner(dataChan)

	sys.RegisterHandle()

	wg.Add(1)
	go func() {
		sys.RegisterHandleSingle(SINGLE, dataChan)
		wg.Done()
	}()
	sys.StartServer(SINGLE)
	wg.Wait()

}
