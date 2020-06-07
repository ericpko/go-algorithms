package main

import (
	"fmt"
	"log"

	"github.com/ericpko/go-algorithms/pkg/palindrome"
)

func main() {
	fmt.Print("Enter a word: ")

	var word string
	_, err := fmt.Scanf("%s", &word)
	if err != nil {
		log.Fatal(err)
	}

	n := len(word)
	var M *[][]int
	numInsertions, M := palindrome.MinInsertions(word)
	fmt.Printf("The minimum number of insersions is: %d\n", numInsertions)

	var minPalindrome string = palindrome.MakePalindrome(word, M)
	fmt.Printf("The shortest length palindrome of <%s> is: <%s>\n", word, minPalindrome)

	// Clear the memory?
	M = nil

	numInsertions = palindrome.MinInsertExponential(word, 0, n-1)
	fmt.Printf("The minimum number of insersions called from <MinInsertExponential> is: %d\n", numInsertions)

	numInsertions = palindrome.MinInsertMemoization(word)
	fmt.Printf("The minimum number of insersions called from <MinInsertMemoization> is: %d\n", numInsertions)
}
