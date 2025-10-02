package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var mu sync.Mutex

func main() {
	data := make(map[int]struct{})
	workers := 1_000

	for range workers {
		go updateData(data)
		go readData(data)
	}
}

func updateData(data map[int]struct{}) {
	for {
		// Generate random number
		randomNumber := rand.Intn(1_000_000_000)
		mu.Lock()
		data[randomNumber] = struct{}{}
		mu.Unlock()
	}
}

func readData(data map[int]struct{}) {
	for {
		// Generate random number
		randomNumber := rand.Intn(1_000_000_000)
		// mu.Lock() <- with this, the race condition is fixed
		_, ok := data[randomNumber]
		// mu.Unlock()
		if !ok {
			fmt.Println("Number not found")
		}
	}
}
