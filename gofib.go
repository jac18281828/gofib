package main

import (
	"fmt"
	"math/big"
	"os"
	"strconv"
)

func fib(n int) (*big.Int, *big.Int) {
	if n == 0 {
		n0, n1 := big.NewInt(0), big.NewInt(1)
		return n0, n1
	} else {
		a, b := fib(n / 2)
		var b2 big.Int
		b2.Mul(b, big.NewInt(2))
		var b2a big.Int
		b2a.Sub(&b2, a)
		var c big.Int
		c.Mul(a, &b2a)
		var bsq big.Int
		bsq.Mul(b, b)
		var asq big.Int
		asq.Mul(a, a)
		var d big.Int
		d.Add(&asq, &bsq)
		if n%2 == 0 {
			return &c, &d
		} else {
			c.Add(&c, &d)
			return &d, &c
		}
	}
}

func main() {
	if len(os.Args) > 1 {
		value := os.Args[1]
		n, err := strconv.Atoi(value)
		if err == nil {
			n0, _ := fib(n)
			fmt.Println("GoFIB n = ", n, ",", n0.Text(10))
		} else {
			fmt.Println(value, "not a number")
		}
	} else {
		fmt.Println("Expected f_n")
	}
}
