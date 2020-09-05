package main

import (
	"flag"
	"fmt"
	"time"
	"strconv"
	"strings"
	"github.com/go-vgo/robotgo"
)


//Global variables
var clicks_per_second int = 60
var keys_start []string = []string{"x","shift","ctrl"}
var keys_stop []string = []string{"a","shift","ctrl"}


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


//special sort (fix bag in robotgo.AddEvents) (example: we have f11,shift,x . After special_sort we will have -> x,f11,shift)
func special_sort(array_to_sort []string) []string {
	
	var sorted_array []string
	
	stop:=len(array_to_sort)

	for len(sorted_array)!=stop{

		var min int = 1000
		var word string
		var num int

		for i, x := range array_to_sort {

			if len(x)<=min {
				min=len(x)
				word=x
				num=i
			}
		}

		sorted_array = append(sorted_array, word)
		array_to_sort = append(array_to_sort[:num], array_to_sort[num+1:]...)
	}

	return sorted_array
}



//main function
func main() {


	//Getting command line flags
	set_clicks := flag.Int("clicking_speed", 60, "clicks per second")
	set_start_keys := flag.String("start_keys", "x+shift+ctrl", "keys for starting clicking")
	set_stop_keys := flag.String("stop_keys", "a+shift+ctrl", "keys for stopping clicking")
	flag.Parse()


	//Set speed of clicks
	clicks_per_second = *set_clicks


	//Set start_keys
	if len(strings.Split(*set_start_keys, "+")) == 3 {
		keys_start = special_sort(strings.Split(*set_start_keys, "+"))
	} else {
		fmt.Println("Flag set_start_keys need 3 keys. Now keys are set as default")
	}

	
	//Set stop_keys
	if len(strings.Split(*set_stop_keys, "+")) == 3 {
		keys_stop = special_sort(strings.Split(*set_stop_keys, "+"))
	} else {
		fmt.Println("Flag set_stop_keys need 3 keys. Now keys are set as default")
	}


	//terminal introduction
	fmt.Println("Autoclicker-go by Gick")
	fmt.Println("######################")
	fmt.Println("To set parametres that you need -> run in terminal(cmd) with flags: --clicking_speed=60 --start_keys=x+ctrl+shift --stop_keys=a+ctrl+shift")
	fmt.Println("YOU NEED SET ONLY 3 KEYS in set_start_keys and set_stop_keys!!!")
	fmt.Println("All keys you can find on https://github.com/go-vgo/robotgo/blob/master/docs/keys.md")
	fmt.Println("######################")
	fmt.Println("Autoclicker speed is " + strconv.Itoa(clicks_per_second) + " clicks per second")
	fmt.Println("To start press " + keys_start[0] + " + " + keys_start[1] + " + " +  keys_start[2])
	fmt.Println("To stop press " + keys_stop[0] + " + " + keys_stop[1] + " + " +  keys_stop[2])

	//start core
	clicker_core()
}