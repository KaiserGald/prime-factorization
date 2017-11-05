package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	err := errors.New("exceeded expected arguments")
	err1 := errors.New("please enter one real, whole number")
	args := os.Args[1:]

	if len(args) > 1 {
		log.Fatal(err)
	} else if len(args) < 1 {
		log.Fatal(err1)
	}
	num, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(findPrimeFactors(num))
}

func findPrimeFactors(num int) (primes []int) {
	// Algorithm found at www.geeksforgeeks.org/print-all-prime-factors-of-a-given-number/

	// Print the number of 2s that divide num
	for num%2 == 0 {
		primes = append(primes, 2)
		num /= 2
	}

	// n must be odd at this point. So we can skip one element
	for i := 3; i*i < num; i = i + 2 {
		// While i divides num, append i and divide num
		for num%i == 0 {
			primes = append(primes, i)
			num /= i
		}
	}

	// This condition is to hanle the case when n is a prime number geater than 2
	if num > 2 {
		primes = append(primes, num)
	}

	return
}
