package main

import (
	"fmt"

	"github.com/ericpko/go-algorithms/pkg/knapsack"
)

func main() {

	item1 := knapsack.Item{Value: 60, Weight: 10}
	item2 := knapsack.Item{Value: 100, Weight: 20}
	item3 := knapsack.Item{Value: 120, Weight: 30}

	items := []knapsack.Item{item1, item2, item3}

	// var M *[][]int
	max_val, _ := knapsack.KnapsackSum(items, 50)

	fmt.Printf("The max value the knapsack can hold is: %d\n", max_val)
}
