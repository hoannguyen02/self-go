1. Chanel act like a queue: First come first serve
2. Why we need chanel? chanel to receive data from go routine
3. Syntax 
```
ch <- v    // Send v to channel ch.(assign data for channel)
v := <-ch  // Receive from ch, and
           // assign value to v.
```
4. Create channel: ch := make(chan int)
```
package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c // it will be blocked here until channel received data

	fmt.Println(x, y, x+y)
}
```
5. Create chanel with Buffered: ch := make(chan int, 2)
```
package main

import (
	"fmt"
)
func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
    ch <- 3 // It will be blocked here because not enough buffer
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
```
6. Range and close, (only sender close channel)
```
package main

import (
	"fmt"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		fmt.Println("push/assign data to channel", x)
		x, y = y, x+y
	}
	close(c) // close mainly tell the for loop in main should end
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c { // this for loop will end once channel is close
		fmt.Println("read data from channel to print", i)
	}
}

```
7. Select(another way to return instead of close channel): The select statement lets a goroutine wait on multiple communication operations. A select blocks until one of its cases can run, then it executes that case. It chooses one at random if multiple are ready.
```
package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}

```

8. Default (The default case in a select is run if no other case is ready. Use a default case to try a send or receive without blocking)
```
package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}
```