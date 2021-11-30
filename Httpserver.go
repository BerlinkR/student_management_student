package main

import (
	"fmt"
	"log"
	"net/http"
)

func main(){

	db, err := readFile("student set.json")
	if err != nil{
		fmt.Println("db init fail")
		return
	}
	mux := http.NewServeMux()
	mux.Handle("/show", http.HandlerFunc(db.server_show_list))
	mux.Handle("/order", http.HandlerFunc(db.server_show_order))
	log.Fatal(http.ListenAndServe("localhost:8000", mux))

}

func (db stulist) server_show_list(w http.ResponseWriter, req *http.Request) {
	for index, stu := range db {
		fmt.Fprintf(w, "%s: %s\n", index, stu)
	}
}

func (db stulist) server_show_order(w http.ResponseWriter, req *http.Request) {
	db = db.order()
	for index, stu := range db {
		fmt.Fprintf(w, "%s: %s\n", index, stu)
	}
}





