package main

import "fmt"

func main() {
	fmt.Println(Hello("Brendan"))
}

func Hello(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}