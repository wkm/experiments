package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const PhilosopherCount = 5

var Chopsticks = make([]sync.Mutex, PhilosopherCount)
var Philosophers = make([]chan int, PhilosopherCount)

// the left chopstick is in the index before this philosopher
func leftChop(position int) *sync.Mutex {
	if position == 0 {
		return &Chopsticks[PhilosopherCount-1]
	}

	return &Chopsticks[position-1]
}

// the right chopstick is in the index of this philosopher
func rightChop(position int) *sync.Mutex {
	return &Chopsticks[position]
}

// naive implementation; causes deadlocks
func philosophate1(position int, ch chan int) {
	eating := false
	for {
		// wait until we're poked
		<-ch

		// change our state
		if eating {
			fmt.Printf("%d: putting down left\n", position)
			leftChop(position).Unlock()
			fmt.Printf("%d: putting down right\n", position)
			rightChop(position).Unlock()
			eating = false
		} else {
			fmt.Printf("%d: picking up left\n", position)
			leftChop(position).Lock()
			fmt.Printf("%d: picking up right\n", position)
			rightChop(position).Lock()
			eating = true
		}
	}
}

// implements a dining philospher simulation
func main() {
	for i := range Philosophers {
		Philosophers[i] = make(chan int)
	}

	for id, philo := range Philosophers {
		fmt.Printf("created philosopher %d\n", id)
		go philosophate1(id, philo)
	}

	// randomly poke philosphers
	for {
		fmt.Printf("%s\n", Philosophers)

		time.Sleep(500 * time.Millisecond)
		index := rand.Int() % PhilosopherCount

		// poke this philosopher
		fmt.Printf("# philosopher %d\n", index)
		Philosophers[index] <- 1
	}
}
