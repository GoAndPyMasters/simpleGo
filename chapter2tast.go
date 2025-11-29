// this is chapter 2 task for array and slice


// The Scanln Trap: fmt.Scanln stops reading at the first
// space. If I type "John Doe", name becomes "John", and
// "Doe" gets stuck in the input buffer (and might get 
// read as the email automatically). For real-world CLI tools, 
// we usually use bufio (buffered I/O), but Scanln is fine
//  for learning.

package main

import "fmt"

func main(){

	scileName := make([]string,0,5)
	var name string
	var email string 
	for{
		fmt.Println("you can type 'exit or 'quit' to stop the program ")
		fmt.Println("Enter your name ")
		_, err := fmt.Scanln(&name)

		if err != nil{
			fmt.Printf("Error leading input %v ", err)
			continue
		}

		if name == "exit" || name =="quit"{
			fmt.Println("Exiting program ")
			break
		}

		fmt.Println("Enter your email")
		_, err = fmt.Scanln(&email)


		if err != nil{
			fmt.Printf("Error leading input %v ", err)
			continue
		}
		
		scileName = append(scileName, name + ":" + email)
		fmt.Printf("Remaining capacity is %d \n", cap(scileName)- len(scileName))
	}
	fmt.Println("You have entered the following names and emails:")
	for index, value :=range scileName{
		fmt.Printf("%d : %s \n", index, value)
	}
}