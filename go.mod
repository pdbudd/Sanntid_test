module main

go 1.18

replace go/orders => ./orders

require go/orders v0.0.0-00010101000000-000000000000

require (
	go/elevator v0.0.0-00010101000000-000000000000 // indirect
	go/movement v0.0.0-00010101000000-000000000000 // indirect
)

replace go/movement => ./movement

replace go/elevator => ./elevator
