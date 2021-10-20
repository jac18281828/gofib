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
			const NTHREAD = 12
			calc := fibonacci.NewPrecompute(10000)
			for n := start; n < end; {
				group_count := NTHREAD
				if group_count > end-n {
					group_count = end - n
				}
				var wg sync.WaitGroup
				wg.Add(group_count)
				for i := 0; i < group_count; i++ {
					go func(n int) {
						fn := calc.Fib(n)
						if fn.ProbablyPrime(8) {
							text := fn.Text(10)
							if len(text) < 5000 {
								fmt.Println("GoFIB ProbablePrime n =", n, fn.Text(10))
							} else {
								fmt.Println("GoFIB ProbablePrime n =", n, len(text), "digits")
							}
						}
						wg.Done()
					}(n + i)
				}
				wg.Wait()
				n += group_count
				if n%100000 == 0 {
					fmt.Println("Trying", n)
				}
			}
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
