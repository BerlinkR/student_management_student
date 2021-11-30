package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
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

func (list stulist) server_show_list(w http.ResponseWriter, req *http.Request) {
	for _, stu := range list {
		fmt.Fprintf(w, "%s\n", stu.stutostr())
	}
}

func (list stulist) server_show_order(w http.ResponseWriter, req *http.Request) {
	list = list.order()
	for _, stu := range list {
		fmt.Fprintf(w, "%s\n", stu.stutostr())
	}
}

func (list stulist) server_show_insert(w http.ResponseWriter, req *http.Request) {
	list = list.order()
	for _, stu := range list {
		fmt.Fprintf(w, "%s\n", stu.stutostr())
	}
}

func (stu student) stutostr() (string){
	var s string
	s = stu.Id + " " + stu.Name + " " + stu.Gender + " " + strconv.Itoa(stu.MarkMath) + " " + strconv.Itoa(stu.MarkEnglish)
	return s
}





