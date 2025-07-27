package main

import (
	"fmt"
)

func main(){
	var a [5]int = [5]int{1,2,3,4,5}
	sum:= 0
	for i := range len(a){
		sum+=a[i]
	}
	fmt.Println(sum)
	fmt.Println(a)

	//slicing
	var b = a[1:4]
	fmt.Println(b)

	//string slicing
	var s = "hello"
	var t = s[2:4]
	fmt.Println(t)

	//initialising with make

	// var users []string = make([]string, 3)
	// println(len(users))
	// println(cap(users))
	// fmt.Println(users)
	// fmt.Print(users[0] == "")
	// fmt.Println(users)

	//pass by reference
	// var users []string = []string{"Rohit", "Eswar"}
	// var users2 = users // Copied by reference
	// users2[0] = "harkirat2"
	// fmt.Print(users)
	


	// pass by value clean process
	var users []string = []string{"Rohit", "Eswar"}
	var users2 []string = make([]string, len(users))
	copy(users2,users)
	
	fmt.Println(users)
	fmt.Println(users2)
	
	//pass by referece by slicing
	// var users []string = []string{"Rohit", "Eswar"}
	// var users2 = users[:]
	// users2[0] = "Rohit2"
	// fmt.Println(users,users2)

	//maps

	customers := make(map[string]int)
	customers["Rohit"]=10
	customers["Eswar"]=20
	fmt.Println(customers["Rohit"])
	fmt.Println(customers)
	delete(customers,"Eswar")
	fmt.Println(customers)

	value, exists := customers["Rohit"]
	fmt.Println(value,exists)
}