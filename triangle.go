package main

import "fmt"

func main() {
	var a int
	var b int
	var c int
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)

	if (c < a+b) && (b < a+c) && (a < b+c) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
