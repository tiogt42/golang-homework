package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 2, 2, 2}

	unique_count(&a)
	fmt.Println(unique_count(&a))
}


func unique_count (a *[]int) int{
	b := 0
	v := 0
	c := 0
	d := 0
	g := 0
	for i  := 0; i < len(*a); i++ {
		for j := i; j < len(*a); j++ {
			if (*a)[i] < (*a)[j] {


				c = c+1
				d = c
			}


			if (*a)[i] > (*a)[j] {


				v = v+1
				g = v
			}
		}
	}
	b = d+g
	return b
}


