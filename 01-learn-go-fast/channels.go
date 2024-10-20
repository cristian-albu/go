package main

import (
	"fmt"
	"math/rand"
	"time"
)

const MAX_CHICKEN_PRICE float32 = 5
const MAX_TOFU_PRICE float32 = 3

func Channels() {
	var c = make(chan int, 5)

	go process(c)

	for i := range c {
		fmt.Println(i)
		time.Sleep(time.Second * 1)
	}

	var chickenChannel = make(chan string)
	var tofuChannel = make(chan string)

	var websites = []string{"Walmart", "Costco", "Whole foods"}

	for i := range websites {
		go checkChickenPrices(websites[i], chickenChannel)
		go checkTofuPrices(websites[i], tofuChannel)
	}
	sendMessage(chickenChannel, tofuChannel)

}

func process(c chan int) {
	defer close(c)
	for i := 0; i < 5; i++ {
		c <- i
	}

	fmt.Println("Exiting process")
}

func checkChickenPrices(website string, chickenChannel chan string) {
	for {
		time.Sleep(time.Second * 1)
		var chickenPrice = rand.Float32() * 20
		if chickenPrice <= MAX_CHICKEN_PRICE {
			chickenChannel <- website
			break
		}
	}
}

func checkTofuPrices(website string, tofuChannel chan string) {
	for {
		time.Sleep(time.Second * 1)
		var chickenPrice = rand.Float32() * 20
		if chickenPrice <= MAX_CHICKEN_PRICE {
			tofuChannel <- website
			break
		}
	}
}

func sendMessage(chickenChannel chan string, tofuChannel chan string) {
	select {
	case website := <-chickenChannel:
		fmt.Printf("Text sent: Found deal on chicken at %v \n", website)
	case website := <-tofuChannel:
		fmt.Printf("Email sent: Found deal on tofu at %v \n", website)
	}
}
