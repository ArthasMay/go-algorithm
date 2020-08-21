package advanced

import (
	"fmt"
	"testing"
)

func Test_Channel(t *testing.T) {
	theMine := [5]string{"rock", "ore", "ore", "rock", "ore"}

	done := make(chan bool)

	oreChan := make(chan string)
	minedOreChan := make(chan string)

	// Finder
	go func(mine [5]string) {
		for _, item := range mine {
			if item == "ore" {
				oreChan <- item
			}
		}
	}(theMine)

	// Breaker
	go func() {
		for i := 0; i < 3; i++ {
			foundOre := <-oreChan //read from oreChannel
			fmt.Println("From Finder:", foundOre)
			minedOreChan <- "minedOre" //send to minedOreChan
		}
		// 不close的话，会阻塞在下面的 range 那边，close()可以显式的通知channel range结束
		close(minedOreChan)
	}()

	// Smelter
	go func() {
		// for i := 0; i < 3; i++ {
		// 	minedOre := <-minedOreChan //read from minedOreChan
		// 	fmt.Println("From Miner:", minedOre)
		// 	fmt.Println("From Smelter: Ore is smelted")
		// }
		for minedOre := range minedOreChan {
			fmt.Println("From Miner:", minedOre)
			fmt.Println("From Smelter: Ore is smelted")
		}
		done <- true
	}()

	<-done
}

func Fabonacci(c chan int, quit chan bool) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x + y
		case <- quit: 
			fmt.Println("quit")
			return
		}
	}
}

func Test_ChannelSelected(t *testing.T) {
	c := make(chan int)
	quit := make(chan bool)
	
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- true
	}()
	
	Fabonacci(c, quit)
}
