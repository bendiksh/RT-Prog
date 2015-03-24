package main

import (
	"fmt"
	"math"
)

type Event struct{
	Floor int
	Type int
}

func main() {

	max_elev := 4
	N_elev := 2 // number of elevators connected to the master/system
	onJob := make([]int, N_elev, max_elev)
	distance := make([]int, N_elev, max_elev)
	for i := 0; i < N_elev; i++ {
	 	if onJob[i] == 0 {
	 		distance[i] = math.Abs()
	 	}
	 } 
}