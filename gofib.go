package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"

	"github.com/jac18281828/gofib/fibonacci"
)

func main() {
	if len(os.Args) > 2 {
		arg1str := os.Args[1]
		arg2str := os.Args[2]
		start, err1 := strconv.Atoi(arg1str)
		end, err2 := strconv.Atoi(arg2str)
		if err1 == nil && err2 == nil {
			var wg sync.WaitGroup
			wg.Add(end - start)
			calc := fibonacci.NewPrecompute(500)
			for n := start; n < end; n++ {
				go func(n int) {
					fn := calc.Fib(n)
					if fn.ProbablyPrime(2) {
						fmt.Println("GoFIB ProbablePrime n =", n, fn.Text(10))
					}
					wg.Done()
				}(n)
			}
			wg.Wait()
		} else {
			if err1 != nil {
				fmt.Println(arg1str, "not a number")
			}
			if err2 != nil {
				fmt.Println(arg2str, "not a number")
			}
		}
	} else if len(os.Args) > 1 {
		arg1str := os.Args[1]
		n, err := strconv.Atoi(arg1str)
		if err == nil {
			calc := fibonacci.NewCalc()
			fn := calc.Fib(n)
			fmt.Println("GoFIB n =", n, fn.Text(10))
		} else {
			fmt.Println(arg1str, "not a number")
		}
	} else {
		fmt.Println("Expected f_n")
	}
}
