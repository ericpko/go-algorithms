package main

import (
	"fmt"
	"github.com/ericpko/go-algorithms/pkg/fibonacci"
	"log"
)


func main() {
	fmt.Print("Enter a number: ");

	var n int;
	_, err := fmt.Scanf("%d", &n);
	if (err != nil) {
		log.Fatal(err);
	}

	fmt.Printf("The %dth Fibonacci number is: %d\n", n, fibonacci.MemoizedFib(n));
}
