package main

import(
	"fmt"
)


func sum(a int, b int)int{
	return a+b
}

func multiplier(factor int) func(int) int {
	return func(a int) int {
		return a * factor
	}
}

func main(){
	fmt.Println(sum(1,2))

	var sum = func(a int , b int)int{
		return a+b
	}(2,3)

	fmt.Println(sum)

	double := multiplier(2)
	tripple := multiplier(3)

	fmt.Println(double(3))
	fmt.Println(tripple(3))
}