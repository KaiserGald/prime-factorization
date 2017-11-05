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

	fmt.Println(num)
	var primes []int
	primes = findPrimeFactors(float32(num))
	fmt.Println(primes)
}

func findPrimeFactors(num float32) []int {
	var primes []int
	return primes
}
