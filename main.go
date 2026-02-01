package main

import (
	"errors"
)

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


