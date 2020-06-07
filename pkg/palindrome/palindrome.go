package palindrome

/**
 * Given a word, find the minimum number of insersions needed
 * to convert it into a palindrome
 */
func MinInsertions(word string) (int, *[][]int) {
	n := len(word)

	// Make an (n+1) x (n+1) matrix
	var M [][]int = make([][]int, n+1)
	for i := 0; i < n+1; i++ { // for each row
		M[i] = make([]int, n+1)
	}

	// Set M[i][j] = 0 for all i >= j
	for i := 0; i <= n; i++ {
		for j := 0; j <= i; j++ {
			M[i][j] = 0
		}
	}

	for k := 1; k <= n-1; k++ {
		for i := 1; i <= n-k; i++ {
			j := i + k

			if word[i-1] == word[j-1] {
				M[i][j] = M[i+1][j-1]

			} else {
				M[i][j] = 1 + min(M[i+1][j], M[i][j-1])
			}
		}
	}

	return M[1][n], &M
}

/**
 * Given a word and the Memoized 2D array of the minimum number
 * of insersions to convert it into a palindrome, return the
 * shortest palindrome of <word> by adding inserts
 *
 * TODO: Must fix on inputs such as pizzapa
 */
func MakePalindrome(word string, M *[][]int) string {
	var n int = len(word)

	// Initialize the palindrome to <word>
	var palindrome string = word

	// Initialize i and j to the start and end indices of the string.
	// (Pretend indices are as in psudocode and match indices into M).
	i, j := 1, n // start and end index of <word> in psudocode respectively
	i_offset, j_offset := 0, 0
	var numInserts = (*M)[1][n]

	for i < j && numInserts > 0 {
		if (*M)[i][j] == (*M)[i+1][j-1] {
			i++
			j--
			i_offset--
			j_offset++

		} else if (*M)[i][j] == 1+(*M)[i+1][j] {
			palindrome = palindrome[:j+j_offset] + string(word[i-1]) + palindrome[j+j_offset:]
			i++
			numInserts--

		} else {
			palindrome = palindrome[:i-2+i_offset] + string(word[j-1]) + palindrome[i-2+i_offset:]
			j--
			numInserts--
		}
	}

	// for i := 0; i <= n; i++ {
	// 	for j := 0; j <= n; j++ {
	// 		fmt.Printf("%d ", (*M)[i][j])
	// 		if j == n {
	// 			fmt.Println()
	// 		}
	// 	}
	// }
	// fmt.Println()

	return palindrome
}

// Min returns the smaller int of x or y.
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

/**
 * Given a word, find the minimum number of insersions needed
 * to convert it into a palindrome. This returns the same as
 * <MinInsertions>, but has exponential time complexity.
 * (For intuition only)
 */
func MinInsertExponential(word string, i, j int) int {

	n := len(word)

	// Base Case:
	if n == 0 || n == 1 {
		return 0
	}
	if i >= j {
		return 0
	}

	// Recursive Step:
	if word[i] == word[j] {
		return MinInsertExponential(word, i+1, j-1)
	} else {
		right := MinInsertExponential(word, i+1, j)
		left := MinInsertExponential(word, i, j-1)

		return 1 + min(left, right)
	}
}

/**
 * Returns the minimum number of insertions to convert <word>
 * into a palindrome.
 */
func MinInsertMemoization(word string) int {
	var n int = len(word)
	if n == 0 || n == 1 {
		return 0
	}

	// Define 2D memoization array. 1D for each index (i, j)
	var M [][]int = make([][]int, n)
	for i := 0; i < n; i++ { // for each row
		M[i] = make([]int, n)
	}

	// Set M[i][j] = 0 for all i >= j
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			M[i][j] = 0
		}
	}

	// NOTE: This is a Lambda work-around in golang since you can't
	// recurse on a lambda (yet?)
	// A true lambda looks like this:
	// lamb := func(...) {...}
	var minInserts func(i, j int) int // helper function
	minInserts = func(i, j int) int {

		// Base case:
		if i >= j {
			return 0
		}

		// Check if M[i][j] has already been set
		if M[i][j] > 0 {
			return M[i][j]

		} else if word[i] == word[j] {
			M[i][j] = minInserts(i+1, j-1)

		} else {
			right := minInserts(i+1, j)
			left := minInserts(i, j-1)

			M[i][j] = 1 + min(left, right)
		}

		return M[i][j]
	}

	return minInserts(0, n-1)
}

/**
 * Two other loops to iterate backwards
 */
// Iterate backwards from MinInsertions
// for k := n - 1; k >= 1; k-- {
// 	for i := 1; i <= n-k; i++ {
// 		j := i + k

// Iterate backwards from MinInsertions
// for k := n - 1; k >= 1; k-- {
// 	for j := n; j >= k+1; j-- {
// 		i := j - k

// Way to iterate forwards without index k
// for i := n - 1; i >= 1; i-- {
// 	for j in range (i + 1, n) {