// Write your answer here, and then test your code.
// Your job is to implement the CalculateMean() function.

package main

import "errors"

// Change these boolean values to control whether you see
// the expected answer and/or hints.
const showExpectedResult = false;
const showHints = false;

// CalculateMean returns mean of the numbers slice, or an error.
func CalculateMean[T any](numbers []T) (*float64, error) {
    // Your code goes here
    return nil, errors.New("empty set")
}
