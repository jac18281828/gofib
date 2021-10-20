package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/jac18281828/gofib/fibonacci"
)

func main() {
	if len(os.Args) > 1 {
		value := os.Args[1]
		n, err := strconv.Atoi(value)
		if err == nil {
			calc := fibonacci.NewCalc()
			n0 := calc.Fib(n)
			fmt.Println("GoFIB n =", n, ",", n0.Text(10))
		} else {
			fmt.Println(value, "not a number")
		}
	} else {
		fmt.Println("Expected f_n")
	}
}
