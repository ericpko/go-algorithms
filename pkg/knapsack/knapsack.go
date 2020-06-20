package knapsack

/**
 * Type Item.
 * NOTE: had to make the attributes start with a capital, so they could be used
 * in other packages outside of package <knapsack>
 */
type Item struct {
	Value  int
	Weight int
}

/**
 * W is the maximum weight that the knapsack can hold
 */
func KnapsackSum(items []Item, W int) (int, *[][]int) {

	n := len(items)

	// Define memoization array and initialize it
	var M [][]int = make([][]int, n+1)
	for i := 0; i <= n; i++ { // for each row/item
		M[i] = make([]int, W+1)

		// Initialize the value to be zero when the knapsack has a max capacity of 0
		M[i][0] = 0
	}
	// Initialize item zero for any weight to be 0. (There is no real item zero)
	for w := 0; w <= W; w++ {
		M[0][w] = 0
	}

	/**
	 * NOTE: We have n items from 1 to n. When we index into array M,
	 * M[i][w] is the value for the ith item because we actually made the
	 * size of M (n+1) rows and (W+1) columns, but when we index into
	 * the array <items>, the "ith" item in M is actually corresponds to
	 * index i-1 in array <items>.
	 */
	for i := 1; i <= n; i++ { // for each item
		// Get the ith item. Just pretend this says item = items[i] as in psudocode.
		item := items[i-1]

		for w := 1; w <= W; w++ { // for each knapsack capacity/weight
			// Check if this item's weight is greater than the knapsack weight/cap
			if item.Weight > w {
				// Skip this item, it doesn't fit in the knapsack
				M[i][w] = M[i-1][w]

			} else {
				M[i][w] = max(M[i-1][w], item.Value+M[i-1][w-item.Weight])
			}
		}
	}

	return M[n][W], &M
}

/**
 * TODO
 */
func PrintKnapsackItems(items []Item, M *[][]int) {
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
