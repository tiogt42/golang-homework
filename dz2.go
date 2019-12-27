package main

import (
	"fmt"
)
func main(){
	var a = 1

	var r int
	r = fibonacci(a)
	fmt.Println(r)
}
func fibonacci(a int) (int){
	var rez = 0
	var i = 1
	var y = 0

	for a>1 {
		rez = y + i
		y = i
		i = rez
		a--
	}

	return rez

}