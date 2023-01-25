package main

import . "fmt"

type student struct{
	name string
	gender string
	data map[string]int
}

func newStudent() *student{
	d := student{}
	d.data = map[string]int{}
	return &d
}

func main() {
	var studentNumber, subjectNumber, score int
	var name, gender, subject string

	Scanln(&studentNumber,&subjectNumber)
	
	s := make([]student,studentNumber)	
	
	for i:=0; i<studentNumber; i++ {	
		obj := newStudent()
		Scanln(&name,&gender)
		obj.name = name;
		obj.gender = gender
		for j:=0; j<subjectNumber; j++ {
			Scanln(&subject,&score)
			obj.data[subject] = score
			
		}
		s[i] = *obj
	}
	
	for i:=0; i<studentNumber; i++ {
		Println("----------")
		Println(s[i].name,s[i].gender)
		
		for index, val := range s[i].data {
			Println(index, val)
		}
		
	}
	Println("----------")
}