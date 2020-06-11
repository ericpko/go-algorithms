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

	// for i := 0; i <= n; i++ {
	// 	for j := 0; j <= n; j++ {
	// 		fmt.Printf("%d ", M[i][j])
	// 		if j == n {
	// 			fmt.Println()
	// 		}
	// 	}
	// }
	// fmt.Println()

	return M[1][n], &M
}

/**
 * Given a word and the Memoized 2D array of the minimum number
 * of insersions to convert it into a palindrome, return the
 * shortest palindrome of <word> by adding inserts
 */
func MakePalindrome(word string, M *[][]int) string {
	var n int = len(word)

	// Initialize the palindrome
	pal_i, pal_j := "", ""

	// Initialize i and j to the start and end indices of the string.
	// (Pretend indices are as in psudocode and match indices into M).
	i, j := 1, n // start and end index of <word> in psudocode respectively

	for i <= j {
		if i == j {
			pal_i = pal_i + string(word[i-1])
			i++

		} else if word[i-1] == word[j-1] {
			pal_i = pal_i + string(word[i-1])
			pal_j = string(word[j-1]) + pal_j
			i++
			j--

		} else if (*M)[i][j] == 1+(*M)[i+1][j] {
			pal_i = pal_i + string(word[i-1])
			pal_j = string(word[i-1]) + pal_j
			i++

		} else if (*M)[i][j] == 1+(*M)[i][j-1] {
			pal_i = pal_i + string(word[j-1])
			pal_j = string(word[j-1]) + pal_j
			j--
		}
	}

	return pal_i + pal_j
}

/**
 * Given a word and the Memoized 2D array of the minimum number
 * of insersions to convert it into a palindrome, return the
 * shortest palindrome of <word> by adding inserts
 *
 * TODO: Fix on inputs such as pizzapar
 */
func MakePalindrome2(word string, M *[][]int) string {
	var n int = len(word)

	// Initialize the palindrome to <word>
	var pal_i string = ""
	var pal_j string = ""

	// Iterate in opposite direction of MinInsertions
	for k := n - 1; k >= 1; k-- {
		for j := n; j >= k+1; j-- {
			i := j - k

			if (*M)[i][j] == (*M)[i+1][j-1] && word[i-1] == word[j-1] {
				pal_i = pal_i + string(word[i-1])
				pal_j = string(word[j-1]) + pal_j

			} else if (*M)[i][j] == 1+(*M)[i+1][j] {
				pal_i = pal_i + string(word[i-1])
				pal_j = string(word[i-1]) + pal_j

			} else if (*M)[i][j] == 1+(*M)[i][j-1] {
				pal_i = pal_i + string(word[j-1])
				pal_j = string(word[j-1]) + pal_j
			}
		}
	}

	return pal_i + pal_j
}

// TODO: fix
func MakePalindromeRec(word, palindrome string, M *[][]int, i, j int) string {

	if i >= j {
		return palindrome
	}

	if (*M)[i][j] == (*M)[i+1][j-1] && word[i-1] == word[j-1] {
		palindrome = MakePalindromeRec(word, palindrome, M, i+1, j-1)

	} else if (*M)[i][j] == 1+(*M)[i+1][j] {
		palindrome = palindrome[:j] + string(word[i-1]) + palindrome[j:]
		palindrome = MakePalindromeRec(word, palindrome, M, i+1, j)

	} else if (*M)[i][j] == 1+(*M)[i][j-1] {
		palindrome = palindrome[:i-1] + string(word[j-1]) + palindrome[i-1:]
		palindrome = MakePalindromeRec(word, palindrome, M, i, j-1)
	}

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
