package main

import "fmt"

func addstudent(filename string) (error){
	stu := input_student()
	list, err := readFile(filename)
	if err != nil{
		return err
	}

	list.insert(stu)
	err2 := writeFile(list, filename)
	if err2 != nil{
		return err2
	}

	return nil
}

func deletestudent(filename string, studentId string) (error){
	list, err := readFile(filename)
	if err != nil{
		return err
	}

	list.delete(studentId)
	err2 := writeFile(list, filename)
	if err2 != nil{
		return err2
	}

	return nil
}

func showlist(filename string) (error){
	list, err := readFile(filename)
	if err != nil{
		fmt.Println("insert false: ", err.Error())
	}
	fmt.Println(list)
	return err
}