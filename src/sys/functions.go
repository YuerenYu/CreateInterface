package sys

import (
	"CreateInterface/src/database"
	"CreateInterface/src/utils"
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
)

type String string

func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if _, err := fmt.Fprintln(w, s); err != nil {
		return
	}
}

func Scanner(ch chan database.InterfaceInfo) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		switch scanner.Text() {
		case "add":
			database.AddInterface(ch)
			break
		case "list":
			database.ListInterface()
			break
		case "delete":
			database.DeleteInterface()
			break
		case "":
			break
		default:
			utils.PrintHelp()
		}
	}

}

func RegisterHandle() {
	var interfaceInfo database.InterfaceInfo
	db, err := database.OpenDB()
	utils.CheckError(err)

	rows, err := db.Query("select * from interface_info")
	utils.CheckError(err)

	for rows.Next() {
		err = rows.Scan(&interfaceInfo.Id, &interfaceInfo.Pattern, &interfaceInfo.Content, &interfaceInfo.InsertTime)
		http.Handle(interfaceInfo.Pattern, String(interfaceInfo.Content))
	}

	err = rows.Close()
	utils.CheckError(err)

	err = db.Close()
	utils.CheckError(err)
}

func RegisterHandleSingle(server chan *http.Server, ch chan database.InterfaceInfo) {
	for info := range ch {
		http.Handle(info.Pattern, String(info.Content))
		RestartServer(server)
	}
}

func StartServer(server chan *http.Server) {
	srv := &http.Server{Addr: ":4000"}
	server <- srv
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			//restart server, err is normal.
		}
		wg.Done()
	}()
	wg.Wait()
}

func stopServer(server chan *http.Server) {
	srv := <-server
	if err := srv.Shutdown(nil); err != nil {
		log.Println(err)
	}
}

func RestartServer(server chan *http.Server) {
	stopServer(server)
	StartServer(server)
}
