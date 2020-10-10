package main

import "fmt"

func main() {
	a, b := "First", "Second"
	//var c, d string
	fmt.Println(swap(a, b))
}

func swap(a, b string) (c, d string) {
	c = b
	d = a
	return // Naked Return
}
