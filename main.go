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

// part4
func ExploreProcess() {
	fmt.Println("=== Process Information ===")
	fmt.Println("PID:", os.Getpid())
	fmt.Println("PPID:", os.Getppid())
	data := []int{1, 2, 3}
	fmt.Printf("Slice address: %p\n", &data)
	fmt.Printf("First element address: %p\n", &data[0])
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

// escape analysis
func AnalyzeEscape() {
	stackVal := CreateOnStack()
	heapVal := CreateOnHeap()

	fmt.Println("Stack value:", stackVal)
	fmt.Println("Heap value:", *heapVal)

}

//part6

func main() {
	fmt.Println("=== Part 1: Math Utilities ===")
	fact, err := Factorial(5)
	if err == nil {
		fmt.Println("Factorial(5):", fact)
	} else {
		fmt.Println("Factorial error:", err)
	}

	prime, err := IsPrime(13)
	if err == nil {
		fmt.Println("IsPrime(13):", prime)
	} else {
		fmt.Println("IsPrime error:", err)
	}

	fmt.Println("\n=== Part 2: Closures ===")
	counter := MakeCounter(0)
	fmt.Println("Counter calls:", counter(), counter(), counter())
	double := MakeMultiplier(2)
	fmt.Println("Double(5):", double(5))
	add, sub, get := MakeAccumulator(100)
	add(50)
	sub(30)
	fmt.Println("Accumulator value:", get())

	fmt.Println("\n=== Part 3: Higher-Order Functions ===")
	nums := []int{1, 2, 3, 4, 5}
	squared := Apply(nums, func(x int) int { return x * x })
	fmt.Println("Apply square:", squared)
	evens := Filter(nums, func(x int) bool { return x%2 == 0 })
	fmt.Println("Filter evens:", evens)
	sum := Reduce(nums, 0, func(acc, curr int) int { return acc + curr })
	fmt.Println("Reduce sum:", sum)
	composed := Compose(func(x int) int { return x + 1 }, func(x int) int { return x * 2 })
	fmt.Println("Compose double then addOne:", composed(5))

	fmt.Println("\n=== Part 4: Process Explorer ===")
	ExploreProcess()

	fmt.Println("\n=== Part 5: Pointer Playground & Escape Analysis ===")
	val := 10
	DoubleValue(&val)
	fmt.Println("DoubleValue result:", val)

	a, b := 5, 10
	a, b = SwapValues(a, b)
	fmt.Println("SwapValues result: a =", a, "b =", b)

	SwapPointers(&a, &b)
	fmt.Println("SwapPointers result: a =", a, "b =", b)

	AnalyzeEscape()
}
