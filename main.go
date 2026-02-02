//part 1

package main

import (
	"errors"
	"fmt"
	"math"
	"os"
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

//part4

// ExploreProcess prints process and memory information
func ExploreProcess() {
	fmt.Println("=== Process Information ===")

	// Current process ID
	pid := os.Getpid()
	fmt.Printf("Current Process ID: %d\n", pid)

	// Parent process ID
	ppid := os.Getppid()
	fmt.Printf("Parent Process ID: %d\n", ppid)

	// Slice and memory addresses
	data := []int{1, 2, 3, 4, 5}
	fmt.Printf("Memory address of slice: %p\n", &data)
	fmt.Printf("Memory address of first element: %p\n", &data[0])

	// Explanation
	fmt.Println("Note: Other processes cannot access these memory addresses due to process isolation")
}

//part5

// DoubleValue takes a pointer to an int and doubles the value
func DoubleValue(x *int) {
    *x = *x * 2
    // This modifies the original variable because we dereference the pointer
}

// CreateOnStack creates a local variable and returns its value
func CreateOnStack() int {
    val := 42
    // This variable stays on the stack
    return val
}

// CreateOnHeap creates a local variable and returns a pointer to it
func CreateOnHeap() *int {
    val := 99
    // This variable escapes to the heap because we return a pointer
    return &val
}

// SwapValues swaps two integers (value semantics)
func SwapValues(a, b int) (int, int) {
    return b, a
}

// SwapPointers swaps the values pointed to by two pointers
func SwapPointers(a, b *int) {
    *a, *b = *b, *a
}

//escape analysis
func AnalyzeEscape() {
    stackVal := CreateOnStack()
    heapVal := CreateOnHeap()

    fmt.Println("Stack value:", stackVal)
    fmt.Println("Heap value:", *heapVal)

    /*
        Explanation:
        - stackVal stays on the stack because we return a plain int.
        - heapVal escapes to the heap because we return a pointer.
        - "Escapes to heap" means Go allocates the variable on the heap so it
          can live beyond the function call (since a pointer to it is returned).
    */
}

func main() {
	AnalyzeEscape()
}