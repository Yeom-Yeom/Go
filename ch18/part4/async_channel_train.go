package main

import . "fmt"

func main() {
	ch := make(chan bool)
	go func(){
		for i:=0; i<20; i++{
			ch<-true
		}
	}()
	
	Println("송신 루틴 완료")

	for i:=0; i<20; i++{
		<-ch
		Printf("수신한 데이터 : %d\n",i)
	}		
}