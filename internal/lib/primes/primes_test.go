package primes

import (
	"slices"
	"testing"
)

func TestPrimes(t *testing.T) {
	primes := FindPrimes(10)
	if !slices.Equal(primes, []int{2, 3, 5, 7}) {
		t.Errorf("primes(10) = %v; want [2, 3, 5, 7]", primes)
	}

	primes = FindPrimes(1)
	if !slices.Equal(primes, []int{}) {
		t.Errorf("primes(1) = %v; want []", primes)
	}
}

func TestGetPrimeFactors(t *testing.T) {
	primes := FindPrimes(10)
	primeFactors := GetPrimeFactors(10, &primes)
	if !slices.Equal(primeFactors, []int{2, 5}) {
		t.Errorf("primeFactors(10) = %v; want [2, 5]", primeFactors)
	}

	primes = FindPrimes(20)
	primeFactors = GetPrimeFactors(20, &primes)
	if !slices.Equal(primeFactors, []int{2, 2, 5}) {
		t.Errorf("primeFactors(20) = %v; want [2, 2, 5]", primeFactors)
	}
}

func TestGetLCM(t *testing.T) {
	lcm := GetLCM(&[]int{4, 21})
	if lcm != 84 {
		t.Errorf("lcm([4, 21]) = %v; want 84", lcm)
	}

	lcm = GetLCM(&[]int{63, 24})
	if lcm != 504 {
		t.Errorf("lcm([63, 24]) = %v; want 504", lcm)
	}
}
