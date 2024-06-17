package main

import "fmt"

func Generator(c chan<- int) {
	for i := 1; i <= 10; i++ {
		c <- i
	}
	close(c)
}

func Double(in <-chan int, out chan<- int) {
	for value := range in {
		out <- 2 * value
	}

	close(out)
}

func PrintValues(c chan int) {
	for value := range c {
		fmt.Println(value)
	}
}

func main(){
	c := make(chan int)
	d := make(chan int)

	go Generator(c)
	go Double(c, d)
	PrintValues(d)
	
}
