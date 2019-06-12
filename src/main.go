package main

import (
	"CreateInterface/src/database"
	"CreateInterface/src/utils"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
)

type String string

func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if _, err := fmt.Fprintln(w, s); err != nil {
		return
	}
}

func main() {
	utils.PrintTips()
	var interfaceInfo database.InterfaceInfo
	//_, err := fmt.Scanf("%s", &interfaceInfo.Pattern)
	//utils.CheckError(err)
	//
	//fmt.Println(interfaceInfo.Pattern)
	//
	//_, err = fmt.Scanf("%s", &interfaceInfo.Content)
	//utils.CheckError(err)
	//
	//fmt.Println(interfaceInfo.Content)
	//
	db, err := database.OpenDB()
	utils.CheckError(err)
	//
	//database.InsertInfo(db, interfaceInfo)
	//
	rows, err := db.Query("select * from interface_info")
	utils.CheckError(err)

	for rows.Next() {
		err = rows.Scan(&interfaceInfo.Id, &interfaceInfo.Pattern, &interfaceInfo.Content, &interfaceInfo.InsertTime)
		fmt.Println(interfaceInfo)
		http.Handle(interfaceInfo.Pattern, String(interfaceInfo.Content))
	}
	rows.Close()

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
