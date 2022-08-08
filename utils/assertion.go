package utils

import (
	"fmt"
	"github.com/go-test/deep"
	"time"
)

func Elapsed(what string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", what, time.Since(start))
	}
}

func AssertEqual0[K any](what string, expected K, fn func() K) {
	defer Elapsed(what)()
	if diff := deep.Equal(fn(), expected); diff != nil {
		emsg := fmt.Sprintf("%s error: %v\n", what, diff)
		panic(emsg)
	}
	fmt.Printf("%s succeeded\n", what)
}

func AssertEqual1[K any, P any](what string, expected K, fn func(P) K, p P) {
	defer Elapsed(what)()
	if diff := deep.Equal(fn(p), expected); diff != nil {
		emsg := fmt.Sprintf("%s error: %v\n", what, diff)
		panic(emsg)
	}
	fmt.Printf("%s succeeded\n", what)
}

func AssertEqual2[K any, P1 any, P2 any](what string, expected K, fn func(P1, P2) K, p1 P1, p2 P2) {
	defer Elapsed(what)()
	if diff := deep.Equal(fn(p1, p2), expected); diff != nil {
		emsg := fmt.Sprintf("%s error: %v\n", what, diff)
		panic(emsg)
	}
	fmt.Printf("%s succeeded\n", what)
}

func AssertEqual3[K any, P1 any, P2 any, P3 any](what string, expected K, fn func(P1, P2, P3) K, p1 P1, p2 P2, p3 P3) {
	defer Elapsed(what)()
	if diff := deep.Equal(fn(p1, p2, p3), expected); diff != nil {
		emsg := fmt.Sprintf("%s error: %v\n", what, diff)
		panic(emsg)
	}
	fmt.Printf("%s succeeded\n", what)
}
