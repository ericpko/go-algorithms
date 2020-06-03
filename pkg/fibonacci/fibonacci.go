package fibonacci

import ()


// Recursive Fibonacci function
func Fib(n int) int {
	if (n <= 1) {
		return n;
	}
	return Fib(n - 1) + Fib(n - 2);
}

// Iterative dynamic programming Fibonacci function
func MemoizedFib(n int) int {
	if (n <= 1) {
		return n;
	}

	var arr []int = []int{0, 1}						// arr[0] = 0, arr[1] = 1
	// var arr = []int{0, 1}							// this also works

	for i := 2; i <= n; i++ {
		arr = append(arr, arr[i - 1] + arr[i - 2]);
	}

	return arr[n];
}
