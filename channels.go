package main

import(
	"fmt"
	// "time"
)

func main(){

	// channel creation 
	// In channel only the intial value can only be taken like in this for loop only 1 will be storing and will print it's value in the channel.

	// channel := make(chan int)
	// go func(){
	// 	for i := 1; i < 10; i++ {
	// 		channel <- i
	// 	}
	// }()

	// value := <- channel

	// fmt.Println(value)

	//channel buffering and defer
	ch:= make(chan int)
	go func(){
		// defer close(ch)
		for i:= 1; i<10; i++{
			fmt.Println("This is number thread",i)
			ch <- i;
		}
	}()

	// time.Sleep(time.Second*5)
	fmt.Println(<-ch)

	// for value:= range ch{
	// 	fmt.Println(value)
	// }

	// fmt.Println(<-ch)
	// value, ok := <-ch

	// fmt.Println("Value of channel = ", value,"|", "Ok status =", ok)

}