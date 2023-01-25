package main
import . "fmt"
var num int

func inputSubNum() (int,error) {		
	Scanln(&num)
	if num >0 {
		return num, nil
	}
	return num, Errorf("잘못된 과목 수입니다.")
	
	
	
}

func average(num int) (float64, error) {
	var score, total float64
	
	for i := 0; i < num; i++ {
		Scanln(&score)
		if score < 0 || score>100{
			return score, Errorf("잘못된 점수입니다.")
		}
		total+=score		
	}	
	
	avg := total/float64(num)
	
	return avg, nil
}

func main() {	
	defer func(){
		r:=recover()
		if r!=nil{
			Println(r)
		}
	}()
	
	num,err1 := inputSubNum()
	
	if err1 != nil{
		panic(err1)
	}
	
	result,err2 := average(num)
	
	if err2 != nil{
		panic(err2)
	}
	
	Println(result)	
}