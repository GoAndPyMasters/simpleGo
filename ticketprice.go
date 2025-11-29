package main

import  "fmt"

func main() {
	// var int ticketPrice = 50
	var name string
	var age uint 
	fmt.Println("Please enter your name and age ")
	_, err := fmt.Scanln(&name, &age)
	if err != nil{
		fmt.Printf("Error reading input %v ", err)
		return
	}

	if age < 18{
		fmt.Printf("Sorry %s, you are too young \n", name)

	}else {
		// var totalticket unit  = 200/50
		fmt.Printf("you can buy %d tickets you changes are %d \n", 200/50, 200%50 )
	}
}