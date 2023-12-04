package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	// no inheritance in golang; No super or parent

	tanishqa := User{"Tanishqa", "tanishqa@gmail.com", true, 16}
	fmt.Println(tanishqa)
	fmt.Printf("hitesh details are: %+v\n", tanishqa)
	fmt.Printf("Name is %v and email is %v.", tanishqa.Name, tanishqa.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
