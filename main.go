package main

import (
	"fmt"
	"time"
)

func main() {
	myChan := make(chan string)

	/*
		go func() {
			myChan <- "Message!"
		}()
	*/
	<-time.After(time.Second * 1)

	select {
	case myChan <- "Message":
		fmt.Println("message sent")
	default:
		fmt.Println("no message sent")
	}

	select {
	case msg := <-myChan:
		fmt.Println(msg)
	default:
		fmt.Println("No message")
	}

	<-time.After(time.Second * 3)

	select {
	case msg := <-myChan:
		fmt.Println(msg)
	default:
		fmt.Println("No message")
	}
}

func main2() {
	theMine := [5]string{"rock", "ore", "ore", "rock", "ore"}

	oreChannel := make(chan string)
	minedOreChan := make(chan string)

	//Finder
	go func(mine [5]string) {
		for _, item := range mine {
			if item == "ore" {
				//send item on oreChannel
				oreChannel <- item
			}
		}
	}(theMine)

	// Ore Breaker
	go func() {
		for foundOre := range oreChannel {
			//read from oreChannel
			///foundOre := <-oreChannel
			fmt.Println("From Finder: ", foundOre)
			//send to minedOreChan
			minedOreChan <- "minedOre"
		}
	}()

	// Smelter
	go func() {
		//for i := 0; i < 3; i++ {
		for minedOre := range minedOreChan {
			//read from minedOreChan
			//minedOre := <-minedOreChan
			fmt.Println("From Miner: ", minedOre)
			fmt.Println("From Smelter: Ore is smelted")
		}
	}()

	<-time.After(time.Second * 5) // Again, you can ignore this

}
