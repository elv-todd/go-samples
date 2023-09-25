package main

import "testing"
import "fmt"

//
// pointer ranges for speed and correctness
//
// https://medium.com/@haaawk/i-thought-i-understood-how-iteration-over-an-array-works-but-apparently-not-in-golang-441a7abd6540
//

type item struct {
	value  int
	weight int
}

func knapsack(items []item, limit int) int {
	var dp [10000]int
	for idx, v := range &dp { // Passing pointer here
		for _, item := range items {
			if idx+item.weight <= limit {
				if dp[idx+item.weight] < v+item.value {
					dp[idx+item.weight] = v + item.value
				}
			}
		}
	}
	result := 0
	for idx := 0; idx <= limit; idx++ {
		if result < dp[idx] {
			result = dp[idx]
		}
	}
	return result
}

func TestKnapsack(t *testing.T) {
	items := []item{
		{value: 6, weight: 2},
		{value: 3, weight: 2},
		{value: 5, weight: 6},
		{value: 4, weight: 5},
		{value: 6, weight: 4},
	}
	limit := 10
	t.Log(knapsack(items, limit))
}

func main() {
	items := []item{
		{value: 6, weight: 2},
		{value: 3, weight: 2},
		{value: 5, weight: 6},
		{value: 4, weight: 5},
		{value: 6, weight: 4},
	}
	limit := 10
	fmt.Print(knapsack(items, limit))
}

