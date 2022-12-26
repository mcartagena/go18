package main

import "fmt"

const var1 = 123
const varStr = "Hi there"

func main() {
	const var2 = 456
	fmt.Printf("variable 1 %d", var1)
	fmt.Printf(" variable 2 %d", var2)
	fmt.Printf(" He say: %s", varStr)
}
