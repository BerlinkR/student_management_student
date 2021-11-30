package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"testing"
)

//If you want to write a test function:
//You must write "Test" at the prefix and then the next word must use big cap as first character
func TestShowList(t *testing.T){
	fetch("http://localhost:8000/show")
}
func TestShowOrder(t *testing.T){
	fetch("http://localhost:8000/order")
}

func TestRemoteInsert(t *testing.T){
	var stu student
	stu.Id = "2017213119"
	stu.Name = "He Hanyue"
	stu.Gender = "Female"
	stu.MarkMath = 92
	stu.MarkEnglish = 89
	err := writepost("http://localhost:8000/insert",stu)
	if err != nil{
		t.Error("wrong happen in writePost of insert test")
	}
}

func fetch(url string){
	resp, err := http.Get(url)

	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}
	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		os.Exit(1)
	}
	fmt.Printf("%s", b)
}

func generateUrlMap(stu student) (map[string][]string){
	stuMap := make(map[string][]string)
	stuMap["Id"] = []string{1:stu.Id}
	stuMap["Name"] = []string{1:stu.Name}
	stuMap["Gender"] = []string{1:stu.Gender}
	stuMap["MarkMath"] = []string{1:strconv.Itoa(stu.MarkMath)}
	stuMap["MarkEnglish"] = []string{1:strconv.Itoa(stu.MarkEnglish)}
	return stuMap
}

func writepost(url string, stu student) (error){

	stumap := generateUrlMap(stu)
	resp, err := http.PostForm(url,stumap)
	if err != nil{
		fmt.Println("err in witepost")
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		fmt.Println("err in witepost")
		return err
	}
	fmt.Println(body)

	return nil
}