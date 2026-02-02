package main

import "testing"

func TestFactorial(t *testing.T) {
	tests := []struct {
		name    string
		input   int
		want    int
		wantErr bool
	}{
		{name: "factorial of 0", input: 0, want: 1, wantErr: false},
		{name: "factorial of 1", input: 1, want: 1, wantErr: false},
		{name: "factorial of 5", input: 5, want: 120, wantErr: false},
		{name: "factorial of 10", input: 10, want: 3628800, wantErr: false},
		{name: "negative input", input: -3, want: 0, wantErr: true},
		{name: "factorial of 3", input: 3, want: 6, wantErr: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Factorial(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Factorial() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Factorial() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsPrime(t *testing.T) {
	tests := []struct {
		name    string
		input   int
		want    bool
		wantErr bool
	}{
		{name: "prime 2", input: 2, want: true, wantErr: false},
		{name: "prime 3", input: 3, want: true, wantErr: false},
		{name: "composite 4", input: 4, want: false, wantErr: false},
		{name: "prime 17", input: 17, want: true, wantErr: false},
		{name: "composite 20", input: 20, want: false, wantErr: false},
		{name: "invalid input", input: 1, want: false, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IsPrime(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsPrime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IsPrime() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestPower(t *testing.T) {
	tests := []struct {
		name    string
		base    int
		exp     int
		want    int
		wantErr bool
	}{
		{name: "2^3", base: 2, exp: 3, want: 8, wantErr: false},
		{name: "5^0", base: 5, exp: 0, want: 1, wantErr: false},
		{name: "0^5", base: 0, exp: 5, want: 0, wantErr: false},
		{name: "negative exponent", base: 2, exp: -1, want: 0, wantErr: true},
		{name: "3^4", base: 3, exp: 4, want: 81, wantErr: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Power(tt.base, tt.exp)
			if (err != nil) != tt.wantErr {
				t.Errorf("Power() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Power() = %v, want %v", got, tt.want)
			}
		})
	}
}

// part 2
func TestMakeCounter(t *testing.T) {
	counter1 := MakeCounter(0)
	counter2 := MakeCounter(10)

	tests := []struct {
		name string
		call func() int
		want int
	}{
		{name: "counter1 first call", call: counter1, want: 1},
		{name: "counter1 second call", call: counter1, want: 2},
		{name: "counter2 first call", call: counter2, want: 11},
		{name: "counter2 second call", call: counter2, want: 12},
		{name: "counter1 third call", call: counter1, want: 3}, // independence check
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.call()
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMakeMultiplier(t *testing.T) {
	double := MakeMultiplier(2)
	triple := MakeMultiplier(3)

	tests := []struct {
		name string
		fn   func(int) int
		in   int
		want int
	}{
		{name: "double 5", fn: double, in: 5, want: 10},
		{name: "triple 5", fn: triple, in: 5, want: 15},
		{name: "double 0", fn: double, in: 0, want: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.fn(tt.in)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMakeAccumulator(t *testing.T) {
	add, sub, get := MakeAccumulator(100)

	tests := []struct {
		name string
		op   func()
		want int
	}{
		{name: "add 50", op: func() { add(50) }, want: 150},
		{name: "subtract 30", op: func() { sub(30) }, want: 120},
		{name: "get current", op: func() {}, want: 120},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.op()
			got := get()
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

// part3
func TestApply(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	tests := []struct {
		name string
		fn   func(int) int
		want []int
	}{
		{name: "square", fn: func(x int) int { return x * x }, want: []int{1, 4, 9, 16}},
		{name: "double", fn: func(x int) int { return x * 2 }, want: []int{2, 4, 6, 8}},
		{name: "negate", fn: func(x int) int { return -x }, want: []int{-1, -2, -3, -4}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Apply(nums, tt.fn)
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("Apply() = %v, want %v", got, tt.want)
					break
				}
			}
		})
	}
}

func TestFilter(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 10, 15}
	tests := []struct {
		name      string
		predicate func(int) bool
		want      []int
	}{
		{name: "evens", predicate: func(x int) bool { return x%2 == 0 }, want: []int{2, 4, 6, 10}},
		{name: "greater than 5", predicate: func(x int) bool { return x > 5 }, want: []int{6, 10, 15}},
		{name: "positives", predicate: func(x int) bool { return x > 0 }, want: []int{1, 2, 3, 4, 5, 6, 10, 15}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Filter(nums, tt.predicate)
			if len(got) != len(tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
				return
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("Filter() = %v, want %v", got, tt.want)
					break
				}
			}
		})
	}
}
func TestReduce(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	tests := []struct {
		name      string
		initial   int
		operation func(int, int) int
		want      int
	}{
		{name: "sum", initial: 0, operation: func(acc, curr int) int { return acc + curr }, want: 10},
		{name: "product", initial: 1, operation: func(acc, curr int) int { return acc * curr }, want: 24},
		{name: "max", initial: 0, operation: func(acc, curr int) int {
			if curr > acc {
				return curr
			}
			return acc
		}, want: 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Reduce(nums, tt.initial, tt.operation)
			if got != tt.want {
				t.Errorf("Reduce() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCompose(t *testing.T) {
	addTwo := func(x int) int { return x + 2 }
	double := func(x int) int { return x * 2 }

	composed := Compose(addTwo, double) // f(g(x)) = addTwo(double(x))

	tests := []struct {
		name string
		in   int
		want int
	}{
		{name: "compose double then addTwo", in: 5, want: 12}, // (5*2)+2
		{name: "compose double then addTwo with 0", in: 0, want: 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := composed(tt.in)
			if got != tt.want {
				t.Errorf("Compose() = %v, want %v", got, tt.want)
			}
		})
	}
}

// part4
func main() {
	// Demonstrate process explorer
	ExploreProcess()
}

//part5
func TestSwapValues(t *testing.T) {
    tests := []struct {
        name   string
        a, b   int
        wantA, wantB int
    }{
        {name: "swap 5 and 10", a: 5, b: 10, wantA: 10, wantB: 5},
        {name: "swap equal values", a: 7, b: 7, wantA: 7, wantB: 7},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            gotA, gotB := SwapValues(tt.a, tt.b)
            if gotA != tt.wantA || gotB != tt.wantB {
                t.Errorf("SwapValues() = (%v, %v), want (%v, %v)", gotA, gotB, tt.wantA, tt.wantB)
            }
        })
    }
}

func TestSwapPointers(t *testing.T) {
    a, b := 5, 10
    SwapPointers(&a, &b)
    if a != 10 || b != 5 {
        t.Errorf("SwapPointers() failed, got a=%v b=%v", a, b)
    }
}
