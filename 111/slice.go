package main

import "fmt"

func main() {
	//var s = make([]int, 0, 5)
	//s := make([]int, 3, 5)
	s := []int{1, 2}
	fmt.Println(len(s), cap(s), s)
	s = append(s, 1)
	fmt.Println(len(s), cap(s), s)
}
