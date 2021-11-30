package main

import (
	"fmt"
	"sort"
)

type student struct {
	Name 		string  `json:"Name"`
	Id 			string  `json:"Id"`
	Gender   	string  `json:"Gender"`
	MarkMath    int		`json:"MarkMath"`
	MarkEnglish int		`json:"MarkEnglish"`
}

type stulist []student

func (list stulist) insert(stu student) {
	list = append(list, stu)
}

func (list stulist) delete(studentId string) (stulist){

	for index, stu := range list{
		if stu.Id == studentId{
			copy(list[index:], list[index+1:])
			return list
		}
	}
	return list
}

func (list stulist) order() (stulist){
	fmt.Println("After order")
	sort.Sort(list)
	fmt.Println(list)
	return list
}

func (list stulist) Len() (int){
	return len(list)
}
func (list stulist) Less(i, j int) (bool){
	return list[i].MarkMath < list[j].MarkMath
}
func (list stulist) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}