// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package intset

import (
	"fmt"
	"testing"
)

func Example_one() {
	//!+main
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String()) // "{1 9 144}"

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String()) // "{9 42}"

	x.UnionWith(&y)
	x.Remove(144)
	//x.Clear()
	fmt.Println(x.String()) // "{1 9 42 144}"

	fmt.Println(x.Has(9), x.Has(123)) // "true false"
	fmt.Println(x.Len())

	c := x.Copy()
	x.Add(999)
	c.Add(888)
	fmt.Println(c)
	fmt.Println(x.String())
	//!-main

	// Output:
	// {1 9 144}
	// {9 42}
	// {1 9 42}
	// true false
	// 3
}

func Example_two() {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Add(42)

	//!+note
	fmt.Println(&x)         // "{1 9 42 144}"
	fmt.Println(x.String()) // "{1 9 42 144}"
	fmt.Println(x)          // "{[4398046511618 0 65536]}"
	//!-note

	// Output:
	// {1 9 42 144}
	// {1 9 42 144}
	// {[4398046511618 0 65536]}
}

func TestDemo(t *testing.T) {
	//Example_one()

	s := 32 << (^uint(0) >> 63)
	fmt.Println((^uint(0) >> 63))
	fmt.Println(s)
}