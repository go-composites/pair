package main

import (
	"fmt"

	Pair "github.com/go-composites/pair/src"
)

func main() {
	entry := Pair.New("answer", 42)
	fmt.Println("entry.ToGoString():", entry.ToGoString())
	fmt.Println("entry.First():", entry.First())
	fmt.Println("entry.Second():", entry.Second())
	fmt.Println("entry.Equal(Pair.New(\"answer\", 42)):", entry.Equal(Pair.New("answer", 42)))
}
