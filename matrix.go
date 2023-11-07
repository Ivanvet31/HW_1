package main

import "fmt"

func main() {
	n := 0
	fmt.Scan(&n)

	matrix := make([][]int, 0)

	for i := 0; i < n; i++ {

		arr := make([]int, 0)

		for j := 0; j < n; j++ {
			elem := 0
			fmt.Scan(&elem)
			arr = append(arr, elem)
		}

		matrix = append(matrix, arr)
	}

	ans := "yes"

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] != matrix[j][i] {
				ans = "no"
				break
			}
		}
	}
	fmt.Print(ans)
}
