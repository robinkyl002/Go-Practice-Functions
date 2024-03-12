package main

// Solution by instructor
import (
  "math"
  "error"
)

func CalculateVariance(numbers []int) (*float64, error) {
	if numbers == nil || len(numbers) == 0 {
		return nil, errors.New("empty numbers")
	} else {
		n := float64(len(numbers))
		sum := func(f func(int) float64) float64{
			var s float64
			for _, num := range numbers {
				s += f(num)
			}
			
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
