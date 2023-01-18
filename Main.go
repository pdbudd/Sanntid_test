package main

import (
	"fmt"
	ord "go/orders"
	mov "go/movement"
	ele "go/elevator"
)

func main(){
	stat1 := ele.State{
		Location: 2,
		Door: false,
		Moving: true}
	ele1 := ele.Elevator{
		Stat: stat1,
		Goal: 5}
	ord.New_Order(4)
	mov.Start(ele1)
	fmt.Println("numero")
}