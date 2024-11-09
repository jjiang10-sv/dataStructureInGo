
package binaryTree

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Philosopher struct {
	Id        int
	LeftFork  chan bool
	RightFork chan bool
}

const (
	RandSecond      = 1e9
	NOfPhilosophers = 8
	Phil            = "Phil"
)

func main() {
	counter := make(chan int, 1)
	counter <- 0

	forks := make([]chan bool, NOfPhilosophers)
	for i := 0; i < len(forks); i++ {
		forks[i] = make(chan bool, 1)
	}

	philosophers := make([]*Philosopher, NOfPhilosophers)
	for i, _ := range forks {
		philosophers[i] = makePhilosopher(i+1, forks[i], forks[(i+1)%NOfPhilosophers])
	}

	wg := sync.WaitGroup{}
	wg.Add(NOfPhilosophers)

	fmt.Printf("There are %v philosophers sitting at a table\n", NOfPhilosophers)
	for _, phil := range philosophers {
		go func(syncer *sync.WaitGroup, ph *Philosopher) {
			defer syncer.Done()
			ph.dining(counter)
			fmt.Printf("%s %v - is done dining\n", Phil, ph.Id)
		}(&wg, phil)
	}
	wg.Wait()
	c := <-counter
	fmt.Printf("%v philosophers finished eating!\n", c)
}

// makePhilosopher makes a new Philosopher struct
func makePhilosopher(index int, leftFork, rightFork chan bool) *Philosopher {
	return &Philosopher{
		Id:        index,
		LeftFork:  leftFork,
		RightFork: rightFork,
	}
}

// dining the process is get the forks, eating, returning the forks and increase the counter by 1
func (phil *Philosopher) dining(counter chan int) {
	phil.getForks()
	phil.eating()
	phil.returnForks()
	c := <-counter
	c += 1
	counter <- c
}

// getForks the process of get the forks is thinking, left fork and right fork
func (phil *Philosopher) getForks() {
	phil.thinking()
	phil.getLeftFork()
	phil.getRightFork()
}

// getLeftFork try to get left fork, if you can't, think and try to get left fork again
func (phil *Philosopher) getLeftFork() {
	select {
	case phil.LeftFork <- true: // Get left fork
		fmt.Printf("%s %v - got the left fork\n", Phil, phil.Id)
		return
	default: // Think and try to get left fork again
		fmt.Printf("%s %v - can't get left fork\n", Phil, phil.Id)
		phil.thinking()
		phil.getLeftFork()
	}
}

// getRightFork try to get right fork, if you can't, put the left fork down, let others use and try to get forks again
func (phil *Philosopher) getRightFork() {
	select {
	case phil.RightFork <- true: // Get right fork
		fmt.Printf("%s %v - got the right fork\n", Phil, phil.Id)
		return
	default: // Put the left fork down, let others use and try to get forks again
		fmt.Printf("%s %v - put the left fork down\n", Phil, phil.Id)
		<-phil.LeftFork
		phil.getForks()
	}
}

// returnForks returns the forks after eating
func (phil *Philosopher) returnForks() {
	<-phil.LeftFork
	<-phil.RightFork
	fmt.Printf("%s %v - return forks\n", Phil, phil.Id)
}

// thinking sleeps for a random period of time
func (phil *Philosopher) thinking() {
	t := time.Duration(rand.Int63n(RandSecond))
	fmt.Printf("%s %v - is thinks for %v\n", Phil, phil.Id, t)
	time.Sleep(t)
}

// eating sleeps for a random period of time
func (phil *Philosopher) eating() {
	t := time.Duration(rand.Int63n(RandSecond))
	fmt.Printf("%s %v - is eats for %v\n", Phil, phil.Id, t)
	time.Sleep(t)
}
