package elevator

import (
	"fmt"
)

func init() {
	fmt.Println("Elevator package initialized")
}

type State struct{
	Location int
	Door bool
	Moving bool
}

func (s *State) Set(loc int, dr bool, mov bool){
	s.Location = loc
	s.Door = dr
	s.Moving = mov
}

type Elevator struct{
	Stat State
	Goal int
}

func (e *Elevator) Set(st State, gl int){
	e.Stat = st
	e.Goal = gl
}