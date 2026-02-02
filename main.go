//part 1

package main

import (
	"errors"
	"math"
)

// Factorial calculates n!
func Factorial(n int) (int, error) {
	if n < 0 {
		return 0, errors.New("factorial is not defined for negative numbers")
	}
	if n == 0 {
		return 1, nil
	}
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result, nil
}

// IsPrime checks if n is prime
func IsPrime(n int) (bool, error) {
	if n < 2 {
		return false, errors.New("prime check requires number >= 2")
	}
	// Only need to check up to sqrt(n)
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false, nil
		}
	}
	return true, nil
}

// Power calculates base^exponent
func Power(base, exponent int) (int, error) {
	if exponent < 0 {
		return 0, errors.New("negative exponents not supported")
	}
	result := 1
	for i := 0; i < exponent; i++ {
		result *= base
	}
	return result, nil
}

//part2

// MakeCounter returns a function that increments and returns a counter
func MakeCounter(start int) func() int {
	counter := start
	return func() int {
		counter++
		return counter
	}
}

// MakeMultiplier returns a function that multiplies input by the captured factor
func MakeMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

// MakeAccumulator returns three functions (add, subtract, get) that share state
func MakeAccumulator(initial int) (func(int), func(int), func() int) {
	acc := initial
	add := func(x int) {
		acc += x
	}
	subtract := func(x int) {
		acc -= x
	}
	get := func() int {
		return acc
	}
	return add, subtract, get
}
