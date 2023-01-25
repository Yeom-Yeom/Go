package main
import . "fmt"


func main(){
	s := make([]string,0,50)
	var name string

	for {
		Scanln(&name)
		if name == "0"{
			break
		}else{
			s = append(s,name)
		}
	}

	for i:=0; i<len(s) ; i++{
		defer Println(s[i])
	}
}