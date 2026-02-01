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
