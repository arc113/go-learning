package main

import "fmt"

func main() {
	var1 := "Hey, I am archit"
	mynum := 13
	var2 := fmt.Sprintf("I like cookies and my birthday falls on %dth August\n", mynum)
	fmt.Printf("%q AND %q", var1, var2)
}
