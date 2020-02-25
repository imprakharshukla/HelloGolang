package main

import (
	"fmt"
	"strconv"
)

var num int

func main() {
	input()

}

func input() {

	fmt.Println("Enter the number")
	fmt.Scan(&num)
	fmt.Println(num)
	fmt.Println("The number of factors are  " + strconv.Itoa(getFactors(num)) + "   " + strconv.Itoa(getReverse()))
	if isPrime(num) {
		if isEmirp() {
			fmt.Println("The number is an emirp number")
		} else {
			fmt.Println("The number is not an emirp number")
		}
	}else {
		fmt.Println("The number is not an emirp number")
	}

}

func getFactors(y int) int {
	c := 0
	for i := 1; i <= y; i++ {
		if y%i == 0 {
			c++
		}
	}
	return c
}

func isPrime(u int) bool {
	o := getFactors(u)
	return o == 2
}

func getReverse() int {
	var r int
	for i := num; i != 0; i = i / 10 {
		r = r * 10
		r = r + i%10
	}
	return r
}

func isEmirp() bool {
	return isPrime(getReverse())
}
