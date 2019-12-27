package main

import (
	"fmt"
)
func main(){
	var a = 8
	var b = 5
	var r int
	r = multiply(a,b)
	fmt.Println(r)
}
func multiply(a int, b int) (int){
	var rez int

	var i = 1
	for i<=b{
		rez = rez+a
		i++
	}

	return rez

}