package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4, 2 ,1,2,3,4,2,1,2,3,1,2,3,1,2,3,1,2,3,1,2,3}

	Unique_count(&a)
	fmt.Println(Unique_count(&a))
}





func Unique_count (a*[] int) int{
	var sum_uniq int = 1
	bubble_Sort(a)
	for i:=1; i<len(*a); i++{
		if (*a)[i-1] != (*a)[i]{
			sum_uniq++
		}
	}
	return sum_uniq

}


func bubble_Sort (a *[]int) {

	for i := 0; i < len(*a); i++ {
		for j := i; j < len(*a); j++ {
			if (*a)[i] > (*a)[j] {

				swap := (*a)[i]
				(*a)[i] = (*a)[j]
				(*a)[j] = swap
			}
		}

	}
}


