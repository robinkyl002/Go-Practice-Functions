package main

/**
My solution for this challenge
**/

/*
The following is the test code used

numbers := []int{4, 6, 9, 45, 8, 17, 3}
learnerResult, learnerError := CalculateVariance(numbers)

*/

import (
	"errors"
	"math"
	// "reflect"
)

// Change these boolean values to control whether you see 
// the expected answer and/or hints.
const showExpectedResult = false;
const showHints = false;

// CalculateVariance returns variance of the numbers slice, or an error.
func CalculateVariance(numbers []int) (*float64, error) {
	if numbers == nil || len(numbers) == 0 {
		return nil, errors.New("not implemented")
	} else {
		floatSlice := make([]float64, len(numbers)) 
		for i, number := range numbers {
			floatSlice[i] = float64(number)
		}
		var variance float64
		variance = SumRes(floatSlice) / float64(len(floatSlice))
		
		return &variance, nil
	}
	
}

func Mean(x []float64) float64{
	mean := Sum(x) / float64(len(x))

	return mean
}

func Sum(x []float64) float64 {
	var sum float64
	for _, i := range x {
		sum += i
	}
	return sum
}

func SumRes(x []float64) float64{
	res := make([]float64, len(x))

	for j, s := range x {
		res[j] = Square(s-Mean(x))
	}
	
	return Sum(res)
}

func Square(x float64) float64{
	squared := math.Pow(x, 2)
	return squared
}