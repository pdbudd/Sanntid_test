package movement

import (
	"fmt"
	ele "go/elevator"
)

func init(){
	fmt.Println("Movement initialised")
}

func Stop(el ele.Elevator) {

}

func Start(el ele.Elevator) {
	fmt.Println("Start elevator going to: ", el.Goal)
}

func Read_State(el ele.Elevator) {

}

func Save_States(el ele.Elevator) {

}

func Retrieve_States(el ele.Elevator) {
	
}