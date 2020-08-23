package main

import (
	"fmt"
	"time"
	"github.com/go-vgo/robotgo"
)


//Global Variables
var keys_start [3]string = [3]string{"x", "shift","ctrl"}
var keys_stop [3]string = [3]string{"a", "shift","ctrl"}
var clicks_per_second int = 60


//function checks if keys pressing
func clicker_core() {
	
	//creating channel
	var stop chan bool = make(chan bool)

	//start keys
	start_keys := robotgo.AddEvents(keys_start[0], keys_start[1], keys_start[2])
		
	if start_keys { 		

		fmt.Println("Clicks started")

		//running clicking function
		go clicking(stop) 
	}

	//stop keys
	stop_keys := robotgo.AddEvents(keys_stop[0], keys_stop[1], keys_stop[2])

	if stop_keys { 

		fmt.Println("Clicks stopped")

		//send true to channel stop
		stop <-true 
	}

	//closing channel
	close(stop)
}


//function that clicking =)
func clicking(stop chan bool) {

	//Creating time duration beetween clicks
	time_for_sleeping := time.Duration(1/float64(clicks_per_second)*10e8)

	for {

		select{

		//stop clicks
		case <-stop :
			return  

		//start clicks
		default :

			//left click
			robotgo.MouseClick("left", false)

			//sleeping beetween clicks
			time.Sleep(time_for_sleeping)

		}

	}

}

// terminal introduction
func print_help() {
	fmt.Println("Autoclicker-go by Gick")
	fmt.Println("Autoclicker speed is 60 clicks per second")
	fmt.Println("To start press ctrl + shift + x")
	fmt.Println("To stop press ctrl + shift + a")
}


//main function
func main() {
	print_help()
	clicker_core()
}