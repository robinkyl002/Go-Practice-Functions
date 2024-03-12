package main

// My solution

import "errors"

// Change these boolean values to control whether you see
// the expected answer and/or hints.
const showExpectedResult = true;
const showHints = false;

// CalculateMean returns mean of the numbers slice, or an error.
func CalculateMean(numbers []int) (*float64, error) {
    // Your code goes here.
    if (len(numbers) == 0) {
        return nil, errors.New("not implemented")
    } else {
        var sum int
        for _, number := range numbers {
            sum = sum + number
        }
        var mean float64
        mean =  float64(sum) / float64(len(numbers))
        return &mean, nil
    }
    
}
