package primes

import (
	"math"
	"slices"

	"github.com/mocdaniel/adventsofcode/internal/lib/lists"
)

// Use sieve of Aristothenes to find primes up to n
func FindPrimes(n int) []int {
	primes := make([]int, 0, n-1)
	for i := 2; i <= n; i++ {
		primes = append(primes, i)
	}
	for i := 0; i < len(primes); i++ {
		for j := i + 1; j < len(primes); j++ {
			if primes[j]%primes[i] == 0 {
				primes = slices.Delete(primes, j, j+1)
				j--
			}
		}
	}
	return primes
}

// gets primes that are factors of n
func GetPrimeFactors(n int, primes *[]int) []int {
	foundPrimes := make([]int, 0)
	for _, v := range *primes {
		if n%v == 0 {
			foundPrimes = append(foundPrimes, v)
			foundPrimes = append(foundPrimes, GetPrimeFactors(n/v, primes)...)
			break
		}
	}
	return foundPrimes
}

// takes a slice of numbers and returns the LCM
func GetLCM(n *[]int) int {

	primeMap := make(map[int]int)

	primes := FindPrimes(slices.Max(*n))
	for _, v := range *n {
		primeFactors := GetPrimeFactors(v, &primes)
		for _, p := range primeFactors {
			if _, ok := primeMap[p]; !ok {
				primeMap[p] = 1
			}
			count := lists.Count(p, &primeFactors)
			if count > primeMap[p] {
				primeMap[p] = count
			}
		}
	}

	lcm := 1
	for k, v := range primeMap {
		lcm *= int(math.Pow(float64(k), float64(v)))
	}

	return lcm
}
