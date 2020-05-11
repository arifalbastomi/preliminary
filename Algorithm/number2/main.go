package main

import("fmt")

func main(){

	var n int = 8
	var steps  = []string{"U","D","D","D","U","D","U","U"}

	var down bool = false
	var position int = 0
	var valleys int =0

	for i:=0 ; i<n ; i++ {

		if(steps[i]=="U"){
			position++
		}else{
			position--
		}

		if(!down && position<0){
			valleys++
			down=true
		}

		if(position>=0){
			down=false
		}

	}

	fmt.Printf("valley:  %d \n", valleys)

}

