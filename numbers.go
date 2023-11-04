package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var nums []string
	var str string
	var res []string

	reader := bufio.NewReader(os.Stdin)
	str, _ = reader.ReadString('\n')

	nums = strings.Fields(str)
	res = append(res, nums[0])

	for _, num := range nums {
		a := 1

		for _, length := range res {
			if num == length {
				a *= 0
			}
		}

		if a != 0 {
			res = append(res, num)
		}
	}

	fmt.Print(len(res))
}
