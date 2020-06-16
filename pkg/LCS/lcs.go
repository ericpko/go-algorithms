package LCS

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
 * TODO: A recursive memoized version
 */
func LCSLengthMemoized(X, Y string) int {
	return 0
}

/**
 * TODO
 */
func LCSLengthDP(X, Y string) int {
	return 0
}

/**
 * TODO
 */
func PrintLCS(X, Y string) nil {
	return
}

/**
 * TODO
 */
func GetLCS(X, Y string) string {
	return ""
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
