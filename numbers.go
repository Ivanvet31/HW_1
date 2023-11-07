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
	var res = make(map[string]int)
	reader := bufio.NewReader(os.Stdin)
	str, _ = reader.ReadString('\n')

	nums = strings.Fields(str)

	for i, num := range nums {
		res[num] = i
	}

	fmt.Print(len(res))
}
