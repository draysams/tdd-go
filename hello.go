package main

import "fmt"


const helloWorldPrefix = "Hello, "
func Hello(name string) string {
	if name == "" {
		name = "world"
	}
	return helloWorldPrefix + name
}

func main() {
	fmt.Println(Hello(""))
}