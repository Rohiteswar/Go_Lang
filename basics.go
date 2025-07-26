package main

import(
	"fmt"
	"math/rand"
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

	//If/else

	var x = rand.Int31n(10)
	if x > 5{
		fmt.Println("x is greater than 5")
	} else{
		fmt.Println("x is less than 5")
		fmt.Println("This is the actual Number = ",x)
	}

	//Switch Case

	Gender := "Men"

	switch Gender{
	case "Men":
		fmt.Println("You are a Men")
	case "Women":
		fmt.Println("You are a `Women")
	}



}
