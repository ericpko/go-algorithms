package LCS

import "fmt"

/**
 * Return the length of the Longest Common Subsequence between
 * <str1> and <str2>.
 * NOTE: This function runs in exponential time complexity and should
 * only be used as an example.
 */
func LCSLengthExponential(X, Y string, m, n int) int {

	if m == 0 || n == 0 {
		return 0

	} else if X[m-1] == Y[n-1] {
		return 1 + LCSLengthExponential(X, Y, m-1, n-1)

	} else {
		return max(LCSLengthExponential(X, Y, m-1, n), LCSLengthExponential(X, Y, m, n-1))
	}
}

/**
 * A memoized recursive version
 */
func LCSLengthMemoized(X, Y string) int {

	m, n := len(X), len(Y)

	// Define memoization array and initialize it
	var M [][]int = make([][]int, m+1)
	for i := 0; i <= m; i++ { // for each row
		M[i] = make([]int, n+1)

		// Set the first column of M to 0
		M[i][0] = 0
	}
	// Set the first row of M to 0
	for j := 0; j <= n; j++ {
		M[0][j] = 0
	}

	// Create the recursive "lambda" function
	var lcsLength func(i, j int) int
	lcsLength = func(i, j int) int {
		if i == 0 || j == 0 {
			return 0

		} else if M[i][j] > 0 {
			return M[i][j]

		} else if X[i-1] == Y[j-1] {
			M[i][j] = 1 + lcsLength(i-1, j-1)

		} else {
			left := lcsLength(i-1, j)
			right := lcsLength(i, j-1)

			M[i][j] = max(left, right)
		}

		return M[i][j]
	}

	return lcsLength(m, n)
}

/**
 * A bottom-up dynamic programming solution that returns the Longest
 * Common Subsequence of the two given strings.
 */
func LCSLengthDP(X, Y string) (int, *[][]int) {
	m, n := len(X), len(Y)

	// Define memoization array and initialize it
	var M [][]int = make([][]int, m+1)
	for i := 0; i <= m; i++ { // for each row
		M[i] = make([]int, n+1)

		// Set the first column of M to 0
		M[i][0] = 0
	}
	// Set the first row of M to 0
	for j := 0; j <= n; j++ {
		M[0][j] = 0
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {

			if X[i-1] == Y[j-1] {
				M[i][j] = 1 + M[i-1][j-1]

			} else {
				M[i][j] = max(M[i-1][j], M[i][j-1])
			}
		}
	}

	return M[m][n], &M
}

/**
 * Print the Longest Common Subsequence between two given strings.
 * NOTE: Could also implement this function without the given memoization
 * array M.
 */
func PrintLCS(X, Y string, M *[][]int) {

	m, n := len(X), len(Y)

	var recPrintLCS func(i, j int)
	recPrintLCS = func(i, j int) {
		if (*M)[i][j] == 0 {
			return

		} else if X[i-1] == Y[j-1] {
			recPrintLCS(i-1, j-1)
			fmt.Print(string(X[i-1]))

		} else if (*M)[i-1][j] > (*M)[i][j-1] {
			recPrintLCS(i-1, j)

		} else {
			recPrintLCS(i, j-1)
		}
	}

	recPrintLCS(m, n)
	fmt.Println()
}

/**
 * Returns the Longest Common Substring between the two given strings.
 */
func GetLCS(X, Y string) string {
	m, n := len(X), len(Y)
	var substr string

	var M *[][]int
	_, M = LCSLengthDP(X, Y)

	// Build the substring
	var recGetLCS func(i, j int) string
	recGetLCS = func(i, j int) string {
		if (*M)[i][j] == 0 {
			return ""

		} else if X[i-1] == Y[j-1] {
			return recGetLCS(i-1, j-1) + string(X[i-1])

		} else if (*M)[i-1][j] > (*M)[i][j-1] {
			return recGetLCS(i-1, j)

		} else {
			return recGetLCS(i, j-1)
		}
	}

	substr = recGetLCS(m, n)
	return substr
}

/**
 * Helper function to compute the max of two ints.
 */
func max(a, b int) int {
	if a >= b {
		return a
	}

	return b
}
