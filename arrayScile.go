package main
import "fmt"

func main(){
//  creating arry and slince and use it 
	names := make([]string, 0, 5)
	names = append(names, "Alice", "Bob", "Charlie", "Diana", "Eve", "Frank")
	age := make([]int, 0,3)
	age = append(age, 23,23,45,65,34,234,23,23,43,45)

	fmt.Println(names)
	fmt.Println(age)

// using loop in golanb
	for i:=0; i<5; i++{
		fmt.Println(i)
	}
	count :=0
	for count<12{
		count++
		// fmt.Println("count value is ", count)
	}
	// Style C: The "For Range" (Python style)

	for index , value := range names{
		fmt.Printf("%d : %s \n", index, value)
	}
}
