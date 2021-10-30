package fibonacci

import (
    "math/big"
)

type Fibonacci interface {
    Fib(n int) *big.Int
}

type Calculator struct {
    fibn map[int]*big.Int
}

func (calc Calculator) Fib(n int) *big.Int {
    n0, _ := fib_fast(n, calc.fibn)
    return n0
}

func NewPrecompute(size int) Fibonacci {
    calc := new(Calculator)
    calc.fibn = make(map[int]*big.Int)
    for i := 0; i < size; i++ {
        ni, _ := fib(i)
        calc.fibn[i] = ni
    }
    return calc
}

func NewCalc() Fibonacci {
    calc := new(Calculator)
    calc.fibn = make(map[int]*big.Int)
    return calc
}

func fast_doubler(n int, a *big.Int, b *big.Int) (*big.Int, *big.Int) {
    var cx big.Int
    cx.Mul(b, big.NewInt(2))
    var b2a big.Int
    b2a.Sub(&cx, a)
    cx.Mul(a, &b2a)
    var bsq big.Int
    bsq.Mul(b, b)
    var asq big.Int
    asq.Mul(a, a)
    var d big.Int
    d.Add(&asq, &bsq)
    if n%2 == 0 {
        return &cx, &d
    } else {
        cx.Add(&cx, &d)
        return &d, &cx
    }
}

func fib_fast(n int, fibn map[int]*big.Int) (*big.Int, *big.Int) {
    value, containsKey := fibn[n+1]
    if containsKey {
        return fibn[n], value
    } else if n == 0 {
        n0, n1 := big.NewInt(0), big.NewInt(1)
        return n0, n1
    } else {
        a, b := fib_fast(n/2, fibn)
        return fast_doubler(n, a, b)
    }
}

func fib(n int) (*big.Int, *big.Int) {
    if n == 0 {
        n0, n1 := big.NewInt(0), big.NewInt(1)
        return n0, n1
    } else {
        a, b := fib(n / 2)
        return fast_doubler(n, a, b)
    }
}
