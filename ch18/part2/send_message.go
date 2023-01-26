package main

import (
	. "fmt" 
	"time"
)

func main() {
	ch := make(chan string)
	defer func(){
		close(ch)
	}()
	
	go sendMessage(ch)
	
	for t:=10; t>0; t-- {
		select {
			case msg := <-ch:
				Printf("%s 메시지가 발송되었습니다.",msg)
				return
			default :
				Printf("%d 초 안에 메시지를 입력하시오.\n",t)
		}
		time.Sleep(1000*time.Millisecond)
	}
	Println("메시지 발송에 실패했습니다.")
}

func sendMessage(ch chan string) {
	var s string
	Scanln(&s)
	ch <- s
}