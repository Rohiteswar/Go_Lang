package main

import(
	"fmt"
)


func sum(a int, b int)int{
	return a+b
}

func main(){
	fmt.Println(sum(1,2))

	var sum = func(a int , b int)int{
		return a+b
	}(2,3)

	fmt.Println(sum)
}