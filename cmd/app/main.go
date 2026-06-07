package main

import (
	"fmt"

	"github.com/ivan-prykhodko/go-slices"
)

func main() {
	l := lyrics()
	fmt.Println(l[0])
	fmt.Println(l[1])
}

func lyrics() []string {
	return []string{
		"Hello",
		"Is it me you're looking for?",
	}
}

func fooBar() {
	// This is just for using any dependency
	_ = slices.Unique(lyrics())
}
