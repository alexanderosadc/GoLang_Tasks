package main

import (
	"fmt"
	"sync"
	"time"
)

type ChopS struct{ sync.Mutex }

type Philo struct {
	id              int
	leftCS, rightCS *ChopS
}

// Used for giving access to the philosophers to the data
// when we send true to the channel then permission is given
// and philpsophers starts eating
// when channel is free than philosopher finished eating
var hostPermission = make(chan bool, 2)
var wg sync.WaitGroup

// Philo object method which works on the locking and unlocking chopsticks
// when the philosopher ate numberOfTimesToEat then the goroutine stops for
// this philospher.
func (p Philo) eat(numberOfTimesToEat int) {
	for index := 0; index < numberOfTimesToEat; index++ {

		hostPermission <- true
		p.leftCS.Lock()
		p.rightCS.Lock()

		fmt.Println("eating ", p.id)
		time.Sleep(1 * time.Second)
		fmt.Println("finished eating ", p.id)

		p.rightCS.Unlock()
		p.leftCS.Unlock()
		<-hostPermission
	}
	wg.Done()
}

func main() {
	CSticks := make([]*ChopS, 5)
	for i := 0; i < 5; i++ {
		CSticks[i] = new(ChopS)
	}

	philos := make([]*Philo, 5)
	for i := 0; i < 5; i++ {
		philos[i] = &Philo{id: i, leftCS: CSticks[i], rightCS: CSticks[(i+1)%5]}
	}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go philos[i].eat(3)
	}

	wg.Wait()
}
