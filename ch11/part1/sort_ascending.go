package main
import . "fmt"

func bubbleSort(nums []int){
	var temp int
	for i:=0; i<len(nums); i++{
		for j:=0; j<(len(nums)-1)-i; j++{
			if nums[j]>nums[j+1]{
				temp = nums[j]
				nums[j] = nums[j+1]
				nums[j+1] = temp
			}
		}
		
	}
}

func inputNums() []int{
	var n int
	Scanln(&n)	
	nums := make([]int, 0, n)
	tmp:=0
	for i:=0; i<n ;i++{
		Scanln(&tmp)
		nums = append(nums,tmp)
	}
	return nums
}

func outputNums(nums []int){
	for i:=0; i<len(nums); i++{
		Printf("%d ",nums[i])
	}
}

func main(){
	nums:=inputNums()
	bubbleSort(nums)
	outputNums(nums)
}