package main

import (
	"fmt"
	"log"

	"github.com/ericpko/go-algorithms/pkg/LCS"
)

func main() {

	var str1, str2 string
	fmt.Print("Enter string 1: ")

	_, err := fmt.Scanf("%s", &str1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Enter string 2: ")
	_, err = fmt.Scanf("%s", &str2)
	if err != nil {
		log.Fatal(err)
	}

	m, n := len(str1), len(str2)
	var M *[][]int

	lcsLength, M := LCS.LCSLengthDP(str1, str2)
	fmt.Printf("<%s> is of length %d and <%s> is of length %d\n\n", str1, m, str2, n)
	fmt.Printf("The Longest Common Subsequence between <%s> and <%s> is: %d\n", str1, str2, lcsLength)

	fmt.Printf("The LCS between <%s> and <%s> is:\n", str1, str2)
	LCS.PrintLCS(str1, str2, M)

	substr := LCS.GetLCS(str1, str2)
	fmt.Printf("The LCS is: <%s>", substr)
}
