package main

import (

	. "fmt" // Using '.' to avoid prefixing functions with their package names
// This is probably not a good idea for large projects...
	"runtime"
)

var i = 0


func Goroutine1(channel chan int, done chan bool) {
	
	var j int
	for j=0;j<1000000;j++{
		i = <- channel
		i++
		channel <- i
	}
	done <- true
}
func Goroutine2(channel chan int, done chan bool) {
	var j int
	for j=0;j<1000000;j++{
		i = <- channel
		i--
		channel <- i
	}
	done <- true
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // I guess this is a hint to what GOMAXPROCS does...
	// Try doing the exercise both with and without it!
	channel := make(chan int, 1)
	done := make(chan bool, 2)
	channel <- i
	go Goroutine1(channel, done) // This spawns someGoroutine() as a goroutine
	go Goroutine2(channel, done)
	//synchronization / joining
	<- done
	<- done
	Printf("%d\n", i)
}
