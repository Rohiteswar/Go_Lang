package main

import(
	"fmt"
	"time"
)

func expensiveTask(str string){
	for i:=1;i<100;i++{
		fmt.Println(i,str)
	}
}

func main(){
	go expensiveTask("First")

	go expensiveTask("Second")

	time.Sleep(time.Second*2)

	fmt.Println("Done!")


}