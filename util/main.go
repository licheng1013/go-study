package main

import (
	"flag"
	"fmt"
)

func main() {
	var port int
	flag.IntVar(&port, "p", 2022, "端口号")
	flag.Parse()
	fmt.Println("param:", port)
	fmt.Println("Hello World")
	study1()
}

// 三个协程 顺序打印 cat，dog，duck，100次
func study1() {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	ch3 := make(chan int, 1)
	chMain := make(chan int, 1)
	go func() {
		for {
			select {
			case _ = <-ch1:
				fmt.Println("cat")
				chMain <- 1
			}
		}
	}()
	go func() {
		for {
			select {
			case _ = <-ch2:
				fmt.Println("dog")
				chMain <- 1
			}
		}
	}()
	go func() {
		for {
			select {
			case _ = <-ch3:
				fmt.Println("duck")
				chMain <- 1
			}
		}
	}()

	i := 1
	for {
		if i >= 100 {
			return
		}

		k := i % 3
		if k == 1 {
			ch1 <- 1
		} else if k == 2 {
			ch2 <- 1
		} else {
			ch3 <- 1
		}
		select {
		case _ = <-chMain:

		}
		i++
	}
}
