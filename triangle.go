package main

import "fmt"

func main() {
	a := 0.0
	b := 0.0
	c := 0.0
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)

	if (c < a+b) && (b < a+c) && (a < b+c) {
		fmt.Println("YES")

	} else {
		fmt.Println("NO")

	}
}
