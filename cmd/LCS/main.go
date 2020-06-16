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

	var lcsLength int = LCS.LCSLengthExponential(str1, str2, m, n)
	fmt.Printf("The Longest Common Subsequence between <%s> and <%s> is: %d\n", str1, str2, lcsLength)
}
