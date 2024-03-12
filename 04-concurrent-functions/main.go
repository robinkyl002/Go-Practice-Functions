// Write your answer here, and then test your code.
// Your job is to implement the CalculateVariance() method.

package main

import (
	"errors"
	"math"
	"sync"
)

// Change these boolean values to control whether you see 
// the expected answer and/or hints.
const showExpectedResult = false;
const showHints = false;

// CalculateVariance returns variance of the numbers slice, or an error.
func CalculateVariance(numbers []int) (*float64, error) {
	if numbers == nil || len(numbers) == 0{
		return nil, errors.New("not implemented")
	} else {
		n := float64(len(numbers))
		sum := func(f func(int) float64) float64{
			var s float64
			var wg sync.WaitGroup
			wg.Add(n)
			for _, num := range numbers {
				defer wg.Done()
				s += f(num)
			}
			wg.Wait()
			return s
		}

		mean := sum(func(i int) float64 {
			return float64(i)
		}) / n

		sumRes := sum(func(i int) float64{
			return math.Pow(float64(i) - mean, 2)
		})
		
		variance := sumRes / n
		
		return &variance, nil
	}
	
}