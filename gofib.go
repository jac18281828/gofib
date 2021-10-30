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
			checkRange(start, end)
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

func min(x1 int, x2 int) int {
	if x1 < x2 {
		return x1
	} else {
		return x2
	}
}

func checkRange(start int, end int) {
	const NGO = 100
	const RABINROUNDS = 8
	const DECIMAL = 10
	const MAX_LEN = 5000
	calc := fibonacci.NewPrecompute(1000)
	for n := start; n < end; {
		group_count := min(NGO, end-n)
		var wg sync.WaitGroup
		wg.Add(group_count)
		for i := 0; i < group_count; i++ {
			go func(n int) {
				fn := calc.Fib(n)
				if fn.ProbablyPrime(RABINROUNDS) {
					text := fn.Text(DECIMAL)
					if len(text) < MAX_LEN {
						fmt.Println("GoFIB ProbablePrime n =", n, fn.Text(DECIMAL))
					} else {
						fmt.Println("GoFIB ProbablePrime n =", n, len(text), "digits")
					}
				}
				wg.Done()
			}(n + i)
		}
		wg.Wait()
		n += group_count
	}
}
