package main

import (
	"fmt"
)

func main() {
	a := []float64{10.2, 4.2, 2.3, 22.2, 2.1, -5.2, -1.1, 0.2}

	Bubble_Sort(&a)
	fmt.Println(a)
}


func Bubble_Sort (a *[]float64){

	for i  := 0; i < len(*a); i++ {
		for j := i; j < len(*a); j++ {
			if (*a)[i] > (*a)[j] {

				swap := (*a)[i]
				(*a)[i] = (*a)[j]
				(*a)[j] = swap
			}
		}
	}
}

