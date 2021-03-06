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
	mux.Handle("/insert", http.HandlerFunc(db.server_show_insert))

	log.Fatal(http.ListenAndServe("localhost:8000", mux))

}

func (db stulist) server_show_list(w http.ResponseWriter, req *http.Request) {
	for _, stu := range db {
		fmt.Fprintf(w, "%s\n", stu.stutostr())
	}
}

func (db stulist) server_show_order(w http.ResponseWriter, req *http.Request) {
	db = db.order()
	fmt.Fprintf(w,"after order:")
	for _, stu := range db {
		fmt.Fprintf(w, "%s\n", stu.stutostr())
	}
}

func (db stulist) server_show_insert(w http.ResponseWriter, req *http.Request) {

	req.ParseForm()
	stu := reqfromToStu(req)
	db = db.insert(stu)
	writeFile(db, "student set.json")
	fmt.Println("after insert")
	for _, stu := range db {
		fmt.Println(stu.stutostr())
	}

}
func (stu student) stutostr() (string){
	var s string
	s = stu.Id + " " + stu.Name + " " + stu.Gender + " " + strconv.Itoa(stu.MarkMath) + " " + strconv.Itoa(stu.MarkEnglish)
	return s
}

func reqfromToStu(req *http.Request) (student){
	var stu student
	stu.Id = req.Form["Id"][1]
	stu.Name = req.Form["Name"][1]
	stu.Gender = req.Form["Gender"][1]
	stu.MarkEnglish,_ = strconv.Atoi(req.Form["MarkEnglish"][1])
	stu.MarkMath,_ = strconv.Atoi(req.Form["MarkMath"][1])
	return stu
}




