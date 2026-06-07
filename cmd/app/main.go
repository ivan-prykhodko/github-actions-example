package main

import "fmt"

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
