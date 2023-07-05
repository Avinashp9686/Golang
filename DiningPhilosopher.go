package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	numPhilosophers = 3
	maxDiningCycles = 2
)

type Philosopher struct {
	id                  int
	LeftFork, rightFork *sync.Mutex
	diningCycles        int
}
type DiningTable struct {
	philosophers []*Philosopher
	waiter       *sync.Mutex
}

func (p *Philosopher) think() {
	fmt.Printf("Philosopher %d is thinking\n", p.id)
	time.Sleep(time.Second)
}
func (p *Philosopher) eat() {
	p.LeftFork.Lock()
	p.rightFork.Lock()
	fmt.Printf("Philosopher %d is eating\n", p.id)
	time.Sleep(time.Second)
	p.rightFork.Unlock()
	p.LeftFork.Unlock()
	p.diningCycles++
}
func (p *Philosopher) dine(table *DiningTable) {
	for p.diningCycles < maxDiningCycles {
		p.think()
		table.waiter.Lock()
		p.eat()
		table.waiter.Unlock()
	}
}
func main() {
	table := &DiningTable{
		philosophers: make([]*Philosopher, numPhilosophers),
		waiter:       &sync.Mutex{},
	}
	//create forks
	forks := make([]*sync.Mutex, numPhilosophers)
	for i := 0; i < numPhilosophers; i++ {
		forks[i] = &sync.Mutex{}
	}
	//Create philosophers and assign forks
	for i := 0; i < numPhilosophers; i++ {
		table.philosophers[i] = &Philosopher{
			id:        i,
			LeftFork:  forks[i],
			rightFork: forks[(i+1)%numPhilosophers],
		}
	}
	//start dining
	var wg sync.WaitGroup
	wg.Add(numPhilosophers)
	for i := 0; i < numPhilosophers; i++ {
		go func(p *Philosopher) {
			defer wg.Done()
			p.dine(table)
		}(table.philosophers[i])
	}
	wg.Wait()
}
