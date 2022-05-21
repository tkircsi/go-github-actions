package main

import (
	"fmt"

	"github.com/tkircsi/go-github-actions/math"
)

func main() {
	fmt.Println("Start app")
	defer func() {
		fmt.Println("Exit app")
	}()

	a, b := 20, 30
	c := math.Sum(a, b)
	fmt.Printf("Sum %d and %d result is %d.\n", a, b, c)
}
