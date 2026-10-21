package main

import (

	"fmt"
)

func main(){

	var theChannels string = "Juststephen" // long form 

	var channel = "Stephenswe" // infarred to string

	subscribers:= 5000 //short form declaration and assignment

	totalSubscibers:= subscribers + 1000

	likes, comments := 100, 30

	fmt.Println("thechannels:", theChannels)
	fmt.Println("channel:", channel)
	fmt.Println("subscriber:", subscribers)
	fmt.Println("totalSubscribers:", totalSubscibers)
	fmt.Println("likes:", likes)
	fmt.Println("comments:", comments)


}