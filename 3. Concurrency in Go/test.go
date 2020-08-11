package main

import (
	"fmt"
	"math"
	"sync"
)

// ChopStick is a
type ChopStick struct{ sync.Mutex }

// Phil is a
type Phil struct {
	number  int
	c       chan int
	rightCS *ChopStick
	leftCS  *ChopStick
}

// Host is a
type Host struct {
	firstP  *Phil
	secondP *Phil
}

func eat(p *Phil, h *Host, wg *sync.WaitGroup, iwg *sync.WaitGroup, c chan int) {
	for i := 0; i < 3; i++ {

		go host(c, p, h, iwg)
		<-p.c

		p.rightCS.Lock()
		p.leftCS.Lock()
		fmt.Println("starting to eat ", p.number)
		fmt.Println("finishing to eat ", p.number)
		p.rightCS.Unlock()
		p.leftCS.Unlock()
	}
	wg.Done()
}

func host(c chan int, p *Phil, h *Host, iwg *sync.WaitGroup) {
	fmt.Printf("Phil %d wants to eat\n", p.number)
	if h.firstP == nil && h.secondP == nil {
		h.firstP = p
		h.firstP.c <- p.number
		fmt.Println(h)
		h.firstP = nil
	} else if h.firstP == nil && math.Abs(float64(p.number)-float64(h.secondP.number)) >= 2 && h.secondP != nil {
		h.firstP = p
		h.firstP.c <- p.number
		fmt.Println(h)
		h.firstP = nil
	} else if h.firstP != nil && math.Abs(float64(p.number)-float64(h.firstP.number)) >= 2 && h.secondP == nil {
		h.secondP = p
		h.secondP.c <- p.number
		fmt.Println(h)
		h.secondP = nil
	} else {
		go host(c, p, h, iwg)
	}

}

func main() {
	chopSticks := make([]*ChopStick, 5)
	phils := make([]*Phil, 5)
	hostChannel := make(chan int)
	h := &Host{}
	wg := &sync.WaitGroup{}
	iwg := &sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		chopSticks[i] = &ChopStick{}
	}

	for i := 0; i < 5; i++ {
		phils[i] = &Phil{
			number:  i + 1,
			c:       make(chan int),
			rightCS: chopSticks[i],
			leftCS:  chopSticks[(i+1)%5],
		}
		wg.Add(1)
		go eat(phils[i], h, wg, iwg, hostChannel)
	}
	wg.Wait()
}
