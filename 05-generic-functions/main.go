// Write your answer here, and then test your code.
// Your job is to implement the CalculateMean() function.

package main

import (
	"errors"
	// "constraints"
)

// Change these boolean values to control whether you see
// the expected answer and/or hints.
const showExpectedResult = false;
const showHints = false;

// CalculateMean returns mean of the numbers slice, or an error.
func CalculateMean[T any](numbers []T) (*float64, error) {
    // Your code goes here
	if numbers == nil || len(numbers) == 0{
		return nil, errors.New("empty set")	
	} else {
		var mean float64
		n := float64(len(numbers))
		
		mean = float64(Sum(numbers)) / n
		return &mean, nil
	}
}

func Sum[T int | uint | float64](x []T) T {
	var sum T
	for _, num := range x {
		sum += num
	}
	return sum
}
