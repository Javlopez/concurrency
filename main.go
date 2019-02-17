package main

import (
	"fmt"
	"time"
)

func main() {
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
		for i := 0; i < 3; i++ {
			//read from oreChannel
			foundOre := <-oreChannel
			fmt.Println("From Finder: ", foundOre)
			//send to minedOreChan
			minedOreChan <- "minedOre"
		}
	}()

	// Smelter
	go func() {
		for i := 0; i < 3; i++ {
			//read from minedOreChan
			minedOre := <-minedOreChan
			fmt.Println("From Miner: ", minedOre)
			fmt.Println("From Smelter: Ore is smelted")
		}
	}()

	<-time.After(time.Second * 5) // Again, you can ignore this

}
