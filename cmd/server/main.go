package main

import "fmt"

func Run() error {
	fmt.Println("hello world")
	return nil
}

func main() {
	fmt.Println("Points Challenge API")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
