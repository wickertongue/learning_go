package main

import (
	"example/hello/morestrings"
	"fmt"

	"github.com/google/go-cmp/cmp"
	"rsc.io/quote/v4"
)

func main() {
	fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
	fmt.Println(quote.Go())
}
