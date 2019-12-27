package main

import "fmt"

func fibbonachi(n int) int{

	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibbonachi(n - 1) + fibbonachi(n - 2)
}
func main() {

	fmt.Println(fibbonachi(4))  // 3

}
