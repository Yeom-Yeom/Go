package main
import . "fmt"

func inputNums() []int{
	var nums []int
	nums = make([]int, 0, 5)
	tmp := 0
	for i:=0; i<5; i++{
		Scanf("%d",&tmp)
		nums = append(nums,tmp)
	}
	return nums
}

func calExam(arr []int)(sum int, over_90 int, under_70 int){
	for i:=0; i<len(arr); i++{
		sum += arr[i]
		if arr[i] >= 90{
			over_90++
		}
		if arr[i] < 70{
			under_70++
		}
	}
	return sum,over_90,under_70
}

func printResult(sum int, over_90 int, under_70 int){
	var result bool = true

	if sum<400 {
		result = false
		Printf("총점이 400점 미만입니다.\n")
	}
	if over_90<2 {
		result=false
		Printf("90점 이상 과목 수가 2개 미만입니다.\n")
	} 
	if under_70 !=0 {
		result = false
		Printf("70점 미만 과목이 있습니다.\n")
	}
	
	if result == false{
		Printf("아이패드를 살 수 없습니다.\n")
	}else{
		Printf("아이패드를 살 수 있습니다.\n")
	}
}

func main() {	
	nums:=inputNums()
	printResult(calExam(nums))
}