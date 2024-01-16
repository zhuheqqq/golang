package main

import "fmt"

func main() {
	const (
		n1 = iota
		n2 = 100
		n3
		n4
	)

	fmt.Println(n1, n2, n3, n4)
}
