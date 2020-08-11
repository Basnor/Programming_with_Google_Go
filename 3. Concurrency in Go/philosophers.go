package main

import (
	"fmt"
	"sync"
)

type ChopS struct{
	sync.Mutex
}

type Philo struct{
	leftCS, rightCS *ChopS
}

func (p Philo) eat(pnum int, wg *sync.WaitGroup) {
	for i := 0; i < 3; i++ {
		p.leftCS.Lock()
		p.rightCS.Lock()

		fmt.Println("starting to eat ", pnum+1)
		fmt.Println("finishing eating ", pnum+1)

		p.rightCS.Unlock()
		p.leftCS.Unlock()
	}

	wg.Done()
}

func main() {
	CSticks := make([]*ChopS, 5)
	var wg sync.WaitGroup

	for  i := 0; i < 5; i++ {
		CSticks[i] = new(ChopS)
	}

	philos := make([]*Philo, 5)
	for i := 0; i < 5; i++ {
		philos[i] = &Philo{CSticks[i], CSticks[(i+1)%5]}
	}

	wg.Add(5)
	for i := 0; i < 5; i++ {
		go philos[i].eat((i+2)%5, &wg)
	}
	wg.Wait()
}