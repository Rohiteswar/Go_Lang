package main

import(
	"fmt"
)

func main(){

	fmt.Println("Hello There!")
	// Output: Hello There!

	//variable
	var name string = "Rohit"
	fmt.Println("Hello There"+" "+name)

	println("======")

	//Loops

	for i:=1;i<10;i++{
		fmt.Println(i);
	}
	//Output: 1 2 3 4 5 6 7 8 9
	fmt.Println("=====")

	//range_loops

	for i := range 5{
		fmt.Println(i);
	}

	



}
