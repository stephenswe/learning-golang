package main


import (
	"fmt"
	"strings"
)


func main(){

	firstName := "Stephen"

	lastName := "Michael"

	fullName := firstName + " " + lastName

	fmt.Println(fullName)

	fmt.Println(strings.ToUpper(fullName))

}