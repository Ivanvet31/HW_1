package main

import "fmt"

func main() {
	n := 0
	fmt.Scanf("%d", &n)

	arr := make([]int, n)

	for i := 0; i < n; i++ {

		if i == n-1 {
			fmt.Scanf("%d", &arr[0])
			break
		}

		fmt.Scan(&arr[i+1])
	}

	for i := 0; i < n; i++ {
		fmt.Printf("%d ", arr[i])
	}
}
