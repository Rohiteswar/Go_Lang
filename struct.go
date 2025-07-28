package main

import(
	"fmt"
)

type rect struct{
	width int
	height int
}

//struct method
// func (r rect) area()int{
// 	return r.width * r.height
// }

//struct pointer method
func (r *rect) area()int{
	return r.width * r.height
}


// func area(r rect)int{
// 	return r.width*r.height
// }

func main(){
	r:= rect{width:20,height:20}
	fmt.Println(r.area())
}