package main

import (
	"fmt"
	"github.com/distatus/battery"
	"time"
)

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func findPrimesSimple(max int) []int {
	var primes []int
	for i := 2; i <= max; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}
	return primes
}

func sieveOfEratosthenes(n int) []int {
	prime := make([]bool, n+1)
	for i := range prime {
		prime[i] = true
	}
	for p := 2; p*p <= n; p++ {
		if prime[p] {
			for i := p * p; i <= n; i += p {
				prime[i] = false
			}
		}
	}
	var primes []int
	for p := 2; p <= n; p++ {
		if prime[p] {
			primes = append(primes, p)
		}
	}
	return primes
}

func main() {
	ma := 1000000000 // Велике число для наглядності

	batteries, err := battery.GetAll()
	if err != nil {
		fmt.Println("Could not get initial battery info!")
		return
	}
	firstBattery := batteries[0] // Assuming single battery system

	//this code is calling the optimized function
	start := time.Now()
	_ = sieveOfEratosthenes(ma)
	fmt.Println("Час виконання (оптимізований):", time.Since(start))

	//this code is calling the non-optimized function
	//start := time.Now()
	//_ = findPrimesSimple(ma)
	//fmt.Println("Час виконання (неоптимізований):", time.Since(start))

	finalBatteries, err := battery.GetAll()
	if err != nil {
		fmt.Println("There was an error processing the battery info!")
		return
	}
	result := finalBatteries[0]

	// Calculate battery usage
	batteryUsed := firstBattery.Current - result.Current
	fmt.Printf("Used Battery: %f mWh\n", batteryUsed)
}
