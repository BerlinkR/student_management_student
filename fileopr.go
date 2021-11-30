package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func input_student() (student){
	fmt.Println("Please input student message: ")
	var nstu student
	fmt.Scanln(&(nstu.Name),&(nstu.Id), &(nstu.Gender), &(nstu.MarkMath), &(nstu.MarkEnglish))
	return nstu
}

func create_and_reset() (error){
	first := []student{{
		Name: "Rong Bailin",
		Id: "2021111239",
		Gender: "Male",
		MarkMath: 89,
		MarkEnglish: 78,
	}}
	err := writeFile(first, "student set.json")

	return err
}

func writeFile(stus stulist,filename string) (error){
	data, err := json.MarshalIndent(stus, "", "    ")
	if err != nil {
		return err
	}
	err2 := ioutil.WriteFile(filename,data, 644)
	if err2 != nil{
		return err2
	}

	return nil
}

func readFile(filename string) (stulist, error){

	filePtr, err := os.Open(filename)
	if err != nil {
		fmt.Println("Open file failed [Err:%s]", err.Error())
		return nil, err
	}
	defer filePtr.Close()

	var stus stulist
	// 创建json解码器
	decoder := json.NewDecoder(filePtr)
	err = decoder.Decode(&stus)
	if err != nil {
		fmt.Println("Decoder failed: ", err.Error())
		return nil, err
	}
	return stus, nil
}


