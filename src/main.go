package main

import (
	"fmt"
	"log"
	"net/http"
)

type String string

func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if _, err := fmt.Fprintln(w, s); err != nil {
		return
	}
	fmt.Println(s)
}

func main() {
	path1 := "/string"
	content := `{
    "code": 10000,
    "msg": "成功",
    "friendlyMsg": "成功",
    "data": {
    }
}`
	http.Handle(path1, String(content))
	ip := string("127.0.0.1")
	port := 4000
	server := fmt.Sprintf("%s:%d", ip, port)
	log.Fatal(http.ListenAndServe(server, nil))

}
