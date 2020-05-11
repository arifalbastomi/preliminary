package main

import "fmt"

func main(){

	var n int  = 100
	var lamp int= 1
	var next int = 1
	var on int= 0

	for lamp = 1; lamp <= n; lamp++ {

		var counter int= 0
		for next = 1; next * next <= n; next++ {

			if (lamp % next == 0) {
				counter++

				if (lamp / next != next){
					counter++
				}
			}
		}

		if (counter % 2 == 1){
			on++
		}
	}
	fmt.Printf("the number of lights that are on: %d  \n", on)
}


