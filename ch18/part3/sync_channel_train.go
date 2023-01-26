package main

import (
	. "fmt"
	"time"	
)

func main() {
	ch := make(chan bool)

	go func(){
		for i:=0; i<20; i++{
			ch<-true
		}
		Print("송신 루틴 완료")
	}()

	for i:= 0; i<20 ;i++{
		if <-ch{
			Printf("수신한 데이터 : %d\n", i)
		}else{
			break
		}
	}	
	time.Sleep(time.Second*3)
	
}