package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x int
	x = 8
	y := 7

	fmt.Println(x)
	fmt.Println(y)

	value, err := strconv.ParseInt("NaN", 0, 64)

	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		fmt.Println(value)
	}

	m := make(map[string]int)
	m["key"] = 6
	fmt.Println(m["key"])

	s := []int{1, 2, 3}
	for index, value := range s {
		fmt.Printf("%v -> %v\n", index, value)
	}

	s = append(s, 16)

	for index, value := range s {
		fmt.Printf("%v -> %v\n", index, value)
	}
}
