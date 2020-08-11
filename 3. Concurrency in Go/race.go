/*
* first  - first goroutine, printed "First."
* second - second goroutine, printed "Second."
* 
* Both goroutines are executed independently, without any synchronization.
* So the relative order of the execution can be different every time. 
* 
* In most cases, the program prints "Second." first, but sometimes it prints "First." first.
* This happens because of race condition. It's a situation there are
* several independent goroutines whose execution order we cannot predict.
*
*/

package main

import (
	"fmt"
	"time"
)

func first() {
	fmt.Println("First.")
}

func second() {
	fmt.Println("Second.")
}

func main() {
	go first()
	go second()
	time.Sleep(time.Millisecond)
}