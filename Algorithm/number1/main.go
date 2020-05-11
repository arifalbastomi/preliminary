package main

import (
	"fmt"
	"sort"
)

func main(){

	var n int = 9
	var numbers = []int{10,20,20,10,10,30,50,10,20}
	sort.Ints(numbers)

	var count int =0

	for i := 0 ; i< n-1 ; i++{
		if(i < n && numbers[i] == numbers[i+1]){
			count++
			i = i + 1
		}

	}

	fmt.Printf("total %d \n", count)

}