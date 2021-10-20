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

func NewCalc() Fibonacci {
	calc := new(Calculator)
	calc.fibn = make(map[int]*big.Int)
	for i := 0; i < 5000; i++ {
		ni, _ := fib(i)
		calc.fibn[i] = ni
	}
	return calc
}

func fib_fast(n int, fibn map[int]*big.Int) (*big.Int, *big.Int) {
	value, containsKey := fibn[n+1]
	if containsKey {
		return fibn[n], value
	} else if n == 0 {
		n0, n1 := big.NewInt(0), big.NewInt(1)
		return n0, n1
	} else {
		a, b := fib_fast(n / 2, fibn)
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
