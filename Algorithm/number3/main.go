package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main(){

	var input int = 1345679
    var conv  = strconv.Itoa(input)
	var len = len(conv)
	newArr:= strings.Split(conv, "")
	var counter2  =0
	for len >= 0 {
		var counter=0
		for j := 1; j <=len; j++ {
			if (counter==0){
				fmt.Print(newArr[counter2])
			}else{
				fmt.Print(0)
			}
			counter++
		}
		counter2++
		fmt.Println();
		len--
	}

}
