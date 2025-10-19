package main

import (
	"fmt"
	"math/rand"
	"time"
)

var MAX_CHICKEN_PRICE = 5
var MAX_TOFU_PRICE = 3

func main() {
	var chickenChannel = make(chan string)
	var tofuChannel = make(chan string)
	var websites = []string{"walmart.com", "google.com", "scienmanas.dev"}
	for i:= range websites {
		go checkChickenPrices(websites[i], chickenChannel)
		go checkTofuPrices(websites[i], tofuChannel)
	}
	sendMessage(chickenChannel, tofuChannel)
}

func checkChickenPrices(website string, chickenChannel chan string) {
	for {
		time.Sleep(time.Second*1)
		var chickenPrice = rand.Float32()*20
		if chickenPrice <= float32(MAX_CHICKEN_PRICE) {
			chickenChannel <- website
			break
		}
	}
}

func checkTofuPrices(website string, tofuChannel chan string) {
	for {
		time.Sleep(time.Second*1)
		var tofu_price = rand.Float32()*20
		if tofu_price <= float32(MAX_TOFU_PRICE) {
			tofuChannel <- website
			break
		}
	}
}

func sendMessage(chickenChannel chan string, tofuChannel chan string) {
	// fmt.Printf("\nFound a deal on chicken at: %s\n", <- chickenChannel)
	select {
	case website:= <- chickenChannel :
		fmt.Printf("\nFound a deal on chicken at: %s\n", website)
	case website:= <- tofuChannel :
		fmt.Printf("\nFound a deal on tofu at: %s\n", website)
	}
}