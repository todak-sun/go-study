package main

import "fmt"

func main() {
	amount := 6
	double(&amount)
	fmt.Println(amount)
}

func double(value *int) {
	*value *= 2
}
