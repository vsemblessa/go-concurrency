package main

import (
	"fmt"
	"sync"
	"time"
)

// Dining philosophers

var EAT_TIME = time.Second * 1
var THINK_TIME = time.Second * 2

var HUNGER = 1

type Philosopher struct {
	leftFork, rightFork int
	name                string
}

type OrderLeave struct {
	mu              *sync.Mutex
	philosophersIdx []int
}

var Philosophers = []Philosopher{
	{name: "Plato", leftFork: 4, rightFork: 0},
	{name: "Socrates", leftFork: 0, rightFork: 1},
	{name: "Aristotle", leftFork: 1, rightFork: 2},
	{name: "Pascal", leftFork: 2, rightFork: 3},
	{name: "Locke", leftFork: 3, rightFork: 4},
}

var Stats = OrderLeave{
	mu:              &sync.Mutex{},
	philosophersIdx: []int{},
}

var Forks []*sync.Mutex

func main() {

	// Used to leave the table
	var table sync.WaitGroup
	table.Add(len(Philosophers))

	var seated sync.WaitGroup
	seated.Add(len(Philosophers))

	// Allocate forks
	for range Philosophers {
		Forks = append(Forks, &sync.Mutex{})
	}

	for idx, philosopher := range Philosophers {
		go func(table *sync.WaitGroup, seated *sync.WaitGroup, philosopher Philosopher, forks []*sync.Mutex, idx int) {
			defer table.Done()

			fmt.Printf("%s is seated. \n", philosopher.name)
			seated.Done()
			seated.Wait() // Wait for everyone to get seated

			for i := HUNGER; i > 0; i-- {
				if philosopher.leftFork > philosopher.rightFork {
					forks[philosopher.rightFork].Lock()
					forks[philosopher.leftFork].Lock()
				} else {
					forks[philosopher.leftFork].Lock()
					forks[philosopher.rightFork].Lock()
				}

				fmt.Printf("\t %s took both forks. \n", philosopher.name)

				fmt.Printf("\t\t %s is eating. \n", philosopher.name)
				time.Sleep(EAT_TIME)

				fmt.Printf("\t\t %s is thinking. \n", philosopher.name)
				time.Sleep(THINK_TIME)

				forks[philosopher.leftFork].Unlock()
				forks[philosopher.rightFork].Unlock()
				fmt.Printf("\t %s has dropped both forks. \n", philosopher.name)
			}

			fmt.Printf("%s is leaving the table. \n", philosopher.name)

			// Record table leave
			Stats.mu.Lock()
			Stats.philosophersIdx = append(Stats.philosophersIdx, idx)
			Stats.mu.Unlock()

		}(&table, &seated, philosopher, Forks, idx)
	}

	table.Wait() // Wait for everyone to leave

	for idx, pIdx := range Stats.philosophersIdx {
		fmt.Printf("\n%d: %s. \n", idx+1, Philosophers[pIdx].name)
	}
}
