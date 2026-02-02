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

//part3

// Apply applies a function to each element of a slice
func Apply(nums []int, f func(int) int) []int {
	result := make([]int, len(nums))
	for i, v := range nums {
		result[i] = f(v)
	}
	return result
}

// Filter returns a new slice containing only elements that satisfy the predicate
func Filter(nums []int, predicate func(int) bool) []int {
	result := []int{}
	for _, v := range nums {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

// Reduce reduces a slice to a single value using the operation function
func Reduce(nums []int, initial int, operation func(int, int) int) int {
	acc := initial
	for _, v := range nums {
		acc = operation(acc, v)
	}
	return acc
}

// Compose returns a new function that applies g first, then f
func Compose(f func(int) int, g func(int) int) func(int) int {
	return func(x int) int {
		return f(g(x))
	}
}
