package euler

import (
    "math"
)

// A simple implementation of the Sieve of Eratosthanes
// Modified from: https://gist.githubusercontent.com/dazfuller/3940659/raw/f9ef272aea070860d9cfe67715b5b401de132745/solution3_snippet.go
func Sieve(n uint64) []uint64 {
    // If n is less than 2 return an empty array
    if n < uint64(2) {
        return make([]uint64, 0)
    }
    
    // Go defaults a bool array to false
    sieve := make([]bool, n)
    sieve[0] = true
    sieve[1] = true
    
    limit := uint64(math.Sqrt(float64(n))) + uint64(1)
    
    // Generate the sieve by removing multiples of primes
    lp := uint64(2)
    for lp < limit {
        for i := lp*2; i < n; i += lp {
            sieve[i] = true
        }

        // Find next prime candidate. Will be between last prime and lp*2
        for i := lp + 1; i < lp*2; i++ {
            if sieve[i] == false { lp = i; break }
        }
    }
    
    // Count the number of primes in the sieve
    count := 0
    for _, v := range sieve {
        if v == false {
            count++
        }
    }
    
    // Build the prime list by looking for the primes in the sieve
    result := make([]uint64, count)
    index := uint64(0)
    for i, v := range sieve {
        if v == false {
            result[index] = uint64(i)
            index++
        }
    }
    
    return result
}


func IsPrime(n uint64) (bool, uint64) {
    primes := Sieve(1000000)
    f := firstFactor(n, primes)

    if f == uint64(1) {
        return true, f
    } else {
        return false, f
    }
}


func checkFactors(n uint64, primes []uint64) uint64 {
    for _, v := range(primes) {
        if n == v { return n }
        if n % v == 0 { return v }
    }

    // No prime factor was found.
    return 1
}


func firstFactor(n uint64, primes []uint64) uint64 {
    // Test small group of primes first since we only need the first factor.
    f := checkFactors(n, primes)

    // Test the full range of primes
    if f == uint64(1) {
        // Only have to test primes upto the square root of the number.
        limit := uint64(math.Sqrt(float64(n))) + uint64(1)
        primes := Sieve(limit)
        f = checkFactors(n, primes)
    }

    return f
}


func Factor(n uint64) []uint64 {

    var factors []uint64
    primes := Sieve(1000000)

    for n > 1 {
        f := firstFactor(n, primes)
        
        if f == 1 {
            factors = append(factors, n)
            break
        }

        factors = append(factors, f)
        n = uint64(n/f)
    }

    if len(factors) == 1 {
        factors = []uint64{1, factors[0]}
    }
    
    return factors
}
