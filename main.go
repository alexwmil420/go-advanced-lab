package main

import (
	"errors"
	"fmt"
	"os"
)

// Part 1 Table-Driven Tests and Math Operations
func Factorial(n int) (int, error) {
	if n < 0 {
		return 0, errors.New("factorial is not defined for negative numbers")
	}

	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}

	return result, nil
}

func IsPrime(n int) (bool, error) {
	if n < 2 {
		return false, errors.New("prime check requires number >= 2")
	}

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false, nil
		}
	}

	return true, nil
}

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

// =================================================================
// Part 2 Factory Functions and Closures
func MakeCounter(start int) func() int {
	count := start
	return func() int {
		count++
		return count
	}
}

func MakeMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

func MakeAccumulator(initial int) (func(int), func(int), func() int) {
	total := initial

	add := func(x int) {
		total += x
	}

	subtract := func(x int) {
		total -= x
	}

	get := func() int {
		return total
	}

	return add, subtract, get
}

// =================================================================
// Part 3 - Higher-Order Functions
func Apply(nums []int, operation func(int) int) []int {
	result := make([]int, len(nums))
	for i, n := range nums {
		result[i] = operation(n)
	}
	return result
}

func Filter(nums []int, predicate func(int) bool) []int {
	result := []int{}
	for _, n := range nums {
		if predicate(n) {
			result = append(result, n)
		}
	}
	return result
}

func Reduce(nums []int, initial int, operation func(acc, current int) int) int {
	acc := initial
	for _, n := range nums {
		acc = operation(acc, n)
	}
	return acc
}

func Compose(f func(int) int, g func(int) int) func(int) int {
	return func(x int) int {
		return f(g(x))
	}
}

// =================================================================
// Part 4 Process Explorer
func ExploreProcess() {
	/* A Process ID (PID) is a unique number the operating system
	   assigns to every running program. It lets the OS track,
	   manage, and control processes
	*/
	pid := os.Getpid()

	/* Parent Process ID (PPID) is the PID of the process that
	   started this program
	*/
	ppid := os.Getppid()

	fmt.Printf("Current Process ID: %d\n", pid)
	fmt.Printf("Parent Process ID: %d\n", ppid)

	/* Process isolation:
	   Each process runs in its own memory space and Other programs
	   cannot directly read or modify that memory. This prevents
	   crashes, bugs, or malicious programs from interfering with
	   each other.
	*/

	data := []int{1, 2, 3, 4, 5}

	/*Slice Memory:

	  &data refers to the address of the slice header and stores a pointer to teh underlying array,
	  the slice length, and its capacity.

	  &data[0] refers to the address of the first element in the underlying
	  array that holds the actual integers.

	  They are different because the slice header only contains metadata about the slice,
	  where as &data[0] points to the actual data stored in memory.

	*/
	fmt.Printf("Memory Address of the slice header: %p\n", &data)
	fmt.Printf("Memory Address of the first element: %p\n", &data[0])

	fmt.Println("Note: Other processes cannot access these memory addresses due to process isolation.")
	fmt.Println()
}

//=================================================================
// Part 5 Pointer Playground and Escape Analysis

/*
Escape Analysis Explanation
--------------------------
Normally, local variables are stored on the stack, which is fast and automatically
cleaned up when a function returns but if Go determines that a variable must live
longer than the function call or be accessed outside of its scope, it moves that
variable to the heap instead. Hence "escaping to the heap".

Variables that escape to the heap include:
1. `count` in MakeCounter:
   - Escapes because the returned closure references it, so it must persist
     after MakeCounter returns.

2. `total` in MakeAccumulator:
   - Escapes for the same reason; the returned closures (`add`, `subtract`, `get`)
     capture it.

3. Function literals / closures:
   - Escape when returned or stored outside their defining function.

4. Slices in Apply and Filter:
   - The underlying array may escape because it is returned to the caller.

5. Slice `data` in ExploreProcess:
   - Accessing element addresses (`&data[0]`) may force it to escape.

6. Values passed to fmt.Printf / fmt.Println:
   - Sometimes treated as escaping because the function may reference them beyond
     the current stack frame.
*/

// 1 DoubleValue: works on a copy, does NOT modify the original
func DoubleValue(x int) {
	x *= 2 // modifies only local copy
}

// 2 DoublePointer: works on a pointer, modifies the original
func DoublePointer(x *int) {
	*x *= 2
}

// 3 CreateOnStack: returns a value, stays on the stack
func CreateOnStack() int {
	value := 42 // declared locally
	return value
}

// 4 CreateOnHeap: returns a pointer, escapes to the heap
func CreateOnHeap() *int {
	value := 100  // declared locally
	return &value // returning pointer forces heap allocation
}

// 5 SwapValues: swaps two integers (no pointers)
func SwapValues(a, b int) (int, int) {
	return b, a
}

// 6 SwapPointers: swaps the values two pointers point to
func SwapPointers(a, b *int) {
	*a, *b = *b, *a
}

// Escape analysis demo
func AnalyzeEscape() {
	stackVar := CreateOnStack() // declared
	heapVar := CreateOnHeap()   // declared

	fmt.Println("Escape Analysis Demo:")
	fmt.Printf("Value from CreateOnStack (stack): %d\n", stackVar)
	fmt.Printf("Value from CreateOnHeap (heap): %d\n", *heapVar)
	fmt.Println("Note: stackVar stays on stack, heapVar escaped to heap")
	fmt.Println()
}

func main() {

	//Part 5 Pointer Playground and Escape Analysis
	// Demo DoubleValue vs DoublePointer
	num1 := 10
	num2 := 10

	fmt.Println("Before DoubleValue:", num1)
	DoubleValue(num1)
	fmt.Println("After DoubleValue:", num1) // unchanged

	fmt.Println("Before DoublePointer:", num2)
	DoublePointer(&num2)
	fmt.Println("After DoublePointer:", num2) // modified

	// Demo SwapValues vs SwapPointers
	a, b := 5, 10
	fmt.Println("\nBefore SwapValues:", a, b)
	a, b = SwapValues(a, b)
	fmt.Println("After SwapValues:", a, b)

	c, d := 5, 10
	fmt.Println("\nBefore SwapPointers:", c, d)
	SwapPointers(&c, &d)
	fmt.Println("After SwapPointers:", c, d)

	// Escape analysis demo
	AnalyzeEscape()

//==============================================================
	//Part 4
	// Process exploration
	ExploreProcess()

//==============================================================
	//PART 6
	// 1. Process Information
	fmt.Println("========== Process Information ==========")
	ExploreProcess()

	// 2. Math Operations Demo
	fmt.Println("\n========== Math Operations ==========")

	// Factorials
	values := []int{0, 5, 10}
	for _, n := range values {
		result, err := Factorial(n)
		if err != nil {
			fmt.Println("Factorial error:", err)
		} else {
			fmt.Printf("%d! = %d\n", n, result)
		}
	}

	// Prime checks
	primes := []int{17, 20, 25}
	for _, n := range primes {
		result, err := IsPrime(n)
		if err != nil {
			fmt.Println("Prime error:", err)
		} else {
			fmt.Printf("%d is prime? %v\n", n, result)
		}
	}

	// Powers
	powerCases := [][2]int{{2, 8}, {5, 3}}
	for _, p := range powerCases {
		result, err := Power(p[0], p[1])
		if err != nil {
			fmt.Println("Power error:", err)
		} else {
			fmt.Printf("%d^%d = %d\n", p[0], p[1], result)
		}
	}

	// 3. Closure Demo
	fmt.Println("\n========== Closures ==========")

	counterA := MakeCounter(0)
	counterB := MakeCounter(100)

	fmt.Println("Counter A calls:", counterA(), counterA(), counterA())
	fmt.Println("Counter B calls:", counterB(), counterB())

	doubler := MakeMultiplier(2)
	tripler := MakeMultiplier(3)

	number := 5
	fmt.Printf("Original: %d\n", number)
	fmt.Printf("Double: %d\n", doubler(number))
	fmt.Printf("Triple: %d\n", tripler(number))

	// 4. Higher-Order Functions Demo
	fmt.Println("\n========== Higher-Order Functions ==========")

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	squared := Apply(nums, func(x int) int {
		return x * x
	})
	fmt.Println("Squared:", squared)

	evens := Filter(nums, func(x int) bool {
		return x%2 == 0
	})
	fmt.Println("Even numbers:", evens)

	sum := Reduce(nums, 0, func(acc, curr int) int {
		return acc + curr
	})
	fmt.Println("Sum:", sum)

	doubleThenAddTen := Compose(
		func(x int) int { return x + 10 },
		func(x int) int { return x * 2 },
	)
	fmt.Println("Double then add 10 (5):", doubleThenAddTen(5))

	// 5. Pointer Demo
	fmt.Println("\n========== Pointer Demo ==========")

	a, b := 3, 7
	fmt.Println("Before SwapValues:", a, b)
	SwapValues(a, b)
	fmt.Println("After SwapValues:", a, b, "(unchanged)")

	c, d := 3, 7
	fmt.Println("Before SwapPointers:", c, d)
	SwapPointers(&c, &d)
	fmt.Println("After SwapPointers:", c, d, "(values swapped)")
}
