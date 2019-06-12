package main

import (
	"CreateInterface/src/database"
	"CreateInterface/src/utils"
	"bufio"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	"os"
)

type String string

func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if _, err := fmt.Fprintln(w, s); err != nil {
		return
	}
}

func main() {
	utils.PrintWelcome()

	go scanner()

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

	ip := string("127.0.0.1")
	port := 4000
	server := fmt.Sprintf("%s:%d", ip, port)
	log.Fatal(http.ListenAndServe(server, nil))

	/*


			path1 := "/string"
			content := `{
		    "code": 10000,
		    "msg": "成功",
		    "friendlyMsg": "成功",
		    "data": {
		    }
		}`
			http.Handle(path1, String(content))


	*/

}

func scanner() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		switch scanner.Text() {
		case "add":
			database.AddInterface()
			break
		case "list":
			database.ListInterface()
			break
		case "delete":
			database.DeleteInterface()
			break
		default:
			utils.PrintHelp()
		}
	}

}
