package main

import (
	"fmt"
	"time"
	"github.com/go-vgo/robotgo"
)


//function checks if keys pressing
func clicker_core(click_per_second int) {
	
	//creating channel
	var stop chan bool = make(chan bool)

	//start keys
	start_keys := robotgo.AddEvents("x", "ctrl", "shift")
		
	if start_keys { 		

		fmt.Println("Clicks started")

		//running clicking function
		go clicking(stop, click_per_second) 
	}

	//stop keys
	stop_keys := robotgo.AddEvents("a", "ctrl", "shift")

	if stop_keys { 

		fmt.Println("Clicks stopped")

		//send true to channel stop
		stop <-true 
	}

	//closing channel
	close(stop)
}


//function that clicking =)
func clicking(stop chan bool, click_per_second int) {

	//Creating time duration beetween clicks
	time_for_sleeping := time.Duration(1/float64(click_per_second)*10e8)

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
	clicker_core(60)
}