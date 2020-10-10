package main

import "fmt"

func main() {
	var1 := "A-B-"
	mynum := 13
	var2 := fmt.Sprintf("I like cookies and my birthday falls on %dth August\n", mynum)
	fmt.Printf("%x AND %s", var1, var2)
	fmt.Println(var1)
}
