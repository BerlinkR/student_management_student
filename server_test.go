package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
)

func Testshowlist(t *testing.T){
	fetch("http://localhost:8000/show")
}
func Testshoworder(t *testing.T){
	fetch("http://localhost:8000/order")
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