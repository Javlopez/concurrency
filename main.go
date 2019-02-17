package main

import (
	"fmt"
	"time"
)

func main() {
	//theMine := [5]string{"ore1", "ore2", "ore3"}
	//oreChan := make(chan string, 3)
	bufferedChan := make(chan string, 3)

	go func() {
		bufferedChan <- "first"
		fmt.Println("Sent 1st")
		bufferedChan <- "second"
		fmt.Println("Sent 2nd")
		bufferedChan <- "third"
		fmt.Println("Sent 3rd")
		bufferedChan <- "four"
		fmt.Println("Sent 4th")
	}()

	go func() {
		firstRead := <-bufferedChan
		fmt.Println("Receiving..")
		fmt.Println(firstRead)
		secondRead := <-bufferedChan
		fmt.Println(secondRead)
		thirdRead := <-bufferedChan
		fmt.Println(thirdRead)
		fourRead := <-bufferedChan
		fmt.Println(fourRead)
	}()

	<-time.After(time.Second * 2)
	/*
		//Finder
		go func(mine [5]string) {
			for _, item := range mine {
				oreChan <- item
			}
		}(theMine)

		// Ore Breaker
		go func() {
			for i := 0; i < 3; i++ {
				foundOre := <-oreChan
				fmt.Println("Miner: Received " + foundOre + " from finder")
			}
		}()
		<-time.After(time.Second * 5)
	*/
}

/*
func main() {
	theMine := [5]string{"rock", "ore", "ore", "rock", "ore"}
	//foundOre := finder(theMine)
	//minedOre := miner(foundOre)
	//smelter(minedOre)
	myFirstChannel := make(chan string)

	myFirstChannel <- "hello"

	myVariable := <-myFirstChannel

	go finder1(theMine)
	go finder2(theMine)
	<-time.After(time.Second * 5) //you can ignore this for now
	fmt.Println("Done")
}

func finder1(param [5]string) {
	fmt.Println("Finder 1 found ore!")
}

func finder2(param [5]string) {
	fmt.Println("Finder 2 found ore!")
}
*/

/*

func finder(param [5]string) [5]string {
	fmt.Println(param)
	x := [5]string{"rock", "ore", "ore", "rock", "ore"}
	return x
}

func miner(param [5]string) [5]string {
	fmt.Println(param)
	x := [5]string{"rock", "ore", "ore", "rock", "ore"}
	return x
}

func smelter(param [5]string) {
	fmt.Println(param)
}*/
