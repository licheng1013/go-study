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
}
