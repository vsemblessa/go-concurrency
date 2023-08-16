package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	BUFFER_SIZE       = 10
	CONSUMER_COUNT    = 3
	PRODUCER_COUNT    = 6
	MAX_PRODUCT_COUNT = 20
)

var (
	wg                  sync.WaitGroup
	converyorBelt       = make(chan string, BUFFER_SIZE)
	mu_Snickers         sync.Mutex
	mu_SnickersConsumed sync.Mutex
	snickersMade        = 0
	snickersConsumed    = 0
)

func producer(id int) {
	defer wg.Done()

	for {
		mu_Snickers.Lock()
		snickersMade++
		if snickersMade > MAX_PRODUCT_COUNT {
			mu_Snickers.Unlock()
			return
		}
		snicker := fmt.Sprintf("Snickers: %d\n", snickersMade)
		mu_Snickers.Unlock()

		converyorBelt <- snicker
		fmt.Printf("* Producer made %s\n", snicker)

		time.Sleep(time.Duration(rand.Intn(1000)+1) * time.Millisecond)
	}
}

func consumer(id int) {
	defer wg.Done()
	for {
		mu_SnickersConsumed.Lock()
		snickersConsumed++
		if snickersConsumed > MAX_PRODUCT_COUNT {
			mu_SnickersConsumed.Unlock()
			return
		}
		mu_SnickersConsumed.Unlock()

		snicker := <-converyorBelt
		fmt.Printf("- Consuming %s\n", snicker)

		time.Sleep(time.Duration(rand.Intn(10000)+1) * time.Millisecond)
	}
}

func main() {

	for p := 1; p <= PRODUCER_COUNT; p++ {
		wg.Add(1)
		go producer(p)
	}

	for c := 1; c <= CONSUMER_COUNT; c++ {
		wg.Add(1)
		go consumer(c)
	}

	time.Sleep(5 * time.Second)

	wg.Wait()
	close(converyorBelt)
}
