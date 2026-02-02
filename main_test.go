package main

import (
	"testing"
)

func TestFactorial(t *testing.T) {
	tests := []struct {
		name    string
		input   int
		want    int
		wantErr bool
	}{
		{"zero", 0, 1, false},
		{"one", 1, 1, false},
		{"three", 3, 6, false},
		{"five", 5, 120, false},
		{"seven", 7, 5040, false},
		{"negative number", -1, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Factorial(tt.input)

			if (err != nil) != tt.wantErr {
				t.Errorf("Factorial(%d) error = %v, wantErr %v",
					tt.input, err, tt.wantErr)
				return
			}

			if got != tt.want {
				t.Errorf("Factorial(%d) = %d, want %d",
					tt.input, got, tt.want)
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
		{"two is prime", 2, true, false},
		{"three is prime", 3, true, false},
		{"four is not prime", 4, false, false},
		{"seventeen is prime", 17, true, false},
		{"twenty is not prime", 20, false, false},
		{"edge case one", 1, false, true},
		{"negative", -5, false, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IsPrime(tt.input)

			if (err != nil) != tt.wantErr {
				t.Errorf("IsPrime(%d) error = %v, wantErr %v",
					tt.input, err, tt.wantErr)
				return
			}

			if got != tt.want {
				t.Errorf("IsPrime(%d) = %v, want %v",
					tt.input, got, tt.want)
			}
		})
	}
}

func TestPower(t *testing.T) {
	tests := []struct {
		name     string
		base     int
		exponent int
		want     int
		wantErr  bool
	}{
		{"two to the third power", 2, 3, 8, false},
		{"five to the zero power", 5, 0, 1, false},
		{"zero to the fifth", 0, 5, 0, false},
		{"one to any power", 1, 10, 1, false},
		{"negative exponent error", 2, -1, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Power(tt.base, tt.exponent)

			if (err != nil) != tt.wantErr {
				t.Errorf("Power(%d, %d) error = %v, wantErr %v",
					tt.base, tt.exponent, err, tt.wantErr)
				return
			}

			if got != tt.want {
				t.Errorf("Power(%d, %d) = %d, want %d",
					tt.base, tt.exponent, got, tt.want)
			}
		})
	}
}

func TestMakeCounter(t *testing.T) {
	tests := []struct {
		name     string
		start    int
		calls    int
		expected []int
	}{
		{
			name:     "start at 0, three calls",
			start:    0,
			calls:    3,
			expected: []int{1, 2, 3},
		},
		{
			name:     "start at 10, two calls",
			start:    10,
			calls:    2,
			expected: []int{11, 12},
		},
		{
			name:     "start at -5, four calls",
			start:    -5,
			calls:    4,
			expected: []int{-4, -3, -2, -1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			counter := MakeCounter(tt.start)
			got := []int{}
			for i := 0; i < tt.calls; i++ {
				got = append(got, counter())
			}

			for i := range got {
				if got[i] != tt.expected[i] {
					t.Errorf("counter() call %d = %d, want %d", i+1, got[i], tt.expected[i])
				}
			}
		})
	}
}

func TestMakeMultiplier(t *testing.T) {
	tests := []struct {
		name     string
		factor   int
		inputs   []int
		expected []int
	}{
		{"double", 2, []int{1, 3, 5}, []int{2, 6, 10}},
		{"triple", 3, []int{2, 4, 6}, []int{6, 12, 18}},
		{"factor zero", 0, []int{1, 2, 3}, []int{0, 0, 0}},
		{"negative factor", -1, []int{1, 2, 3}, []int{-1, -2, -3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			multiplier := MakeMultiplier(tt.factor)
			for i, input := range tt.inputs {
				got := multiplier(input)
				if got != tt.expected[i] {
					t.Errorf("multiplier(%d) = %d, want %d", input, got, tt.expected[i])
				}
			}
		})
	}
}

func TestMakeAccumulator(t *testing.T) {
	tests := []struct {
		name     string
		initial  int
		adds     []int
		subs     []int
		expected int
	}{
		{"basic operations", 100, []int{50, 25}, []int{30}, 145},
		{"only subtract", 50, []int{}, []int{10, 5}, 35},
		{"add and subtract zero", 10, []int{0}, []int{0}, 10},
		{"negative adds and subs", 20, []int{-5, -10}, []int{-5}, 0}, //error
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			add, sub, get := MakeAccumulator(tt.initial)

			for _, a := range tt.adds {
				add(a)
			}

			for _, s := range tt.subs {
				sub(s)
			}

			got := get()
			if got != tt.expected {
				t.Errorf("accumulator result = %d, want %d", got, tt.expected)
			}
		})
	}
}

func TestApply(t *testing.T) {
	tests := []struct {
		name      string
		nums      []int
		operation func(int) int
		expected  []int
	}{
		{"square numbers", []int{1, 2, 3, 4}, func(x int) int { return x * x }, []int{1, 4, 9, 16}},
		{"double numbers", []int{1, 2, 3}, func(x int) int { return x * 2 }, []int{2, 4, 6}},
		{"negate numbers", []int{-1, 2, -3}, func(x int) int { return -x }, []int{1, -2, 3}},
		{"empty slice", []int{}, func(x int) int { return x * x }, []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Apply(tt.nums, tt.operation)
			for i := range got {
				if got[i] != tt.expected[i] {
					t.Errorf("Apply(%v) = %v, want %v", tt.nums, got, tt.expected)
					return
				}
			}
			if len(got) != len(tt.expected) {
				t.Errorf("Apply(%v) length = %d, want %d", tt.nums, len(got), len(tt.expected))
			}
		})
	}
}

func TestFilter(t *testing.T) {
	tests := []struct {
		name      string
		nums      []int
		predicate func(int) bool
		expected  []int
	}{
		{"even numbers", []int{1, 2, 3, 4, 5}, func(x int) bool { return x%2 == 0 }, []int{2, 4}},
		{"positive numbers", []int{-2, -1, 0, 1, 2}, func(x int) bool { return x > 0 }, []int{1, 2}},
		{"greater than 3", []int{1, 3, 5, 7}, func(x int) bool { return x > 3 }, []int{5, 7}},
		{"empty slice", []int{}, func(x int) bool { return x > 0 }, []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Filter(tt.nums, tt.predicate)
			if len(got) != len(tt.expected) {
				t.Errorf("Filter(%v) length = %d, want %d", tt.nums, len(got), len(tt.expected))
			}
			for i := range got {
				if got[i] != tt.expected[i] {
					t.Errorf("Filter(%v) = %v, want %v", tt.nums, got, tt.expected)
				}
			}
		})
	}
}

func TestReduce(t *testing.T) {
	tests := []struct {
		name      string
		nums      []int
		initial   int
		operation func(int, int) int
		expected  int
	}{
		{"sum", []int{1, 2, 3, 4}, 0, func(acc, cur int) int { return acc + cur }, 10},
		{"product", []int{1, 2, 3, 4}, 1, func(acc, cur int) int { return acc * cur }, 24},
		{"max", []int{1, 5, 3, 4}, 0, func(acc, cur int) int {
			if cur > acc {
				return cur
			}
			return acc
		}, 5},
		{"min", []int{7, 2, 9, 4}, 100, func(acc, cur int) int {
			if cur < acc {
				return cur
			}
			return acc
		}, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Reduce(tt.nums, tt.initial, tt.operation)
			if got != tt.expected {
				t.Errorf("Reduce(%v) = %d, want %d", tt.nums, got, tt.expected)
			}
		})
	}
}

func TestCompose(t *testing.T) {
	double := func(x int) int { return x * 2 }
	addTen := func(x int) int { return x + 10 }

	tests := []struct {
		name     string
		f        func(int) int
		g        func(int) int
		input    int
		expected int
	}{
		{"double then add 10", addTen, double, 5, 20}, // (5*2)+10
		{"add 10 then double", double, addTen, 5, 30}, // (5+10)*2
		{"double then double", double, double, 3, 12}, // (3*2)*2
		{"add10 then add10", addTen, addTen, 4, 24},   // (4+10)+10
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			composed := Compose(tt.f, tt.g)
			got := composed(tt.input)
			if got != tt.expected {
				t.Errorf("Compose result = %d, want %d", got, tt.expected)
			}
		})
	}
}
