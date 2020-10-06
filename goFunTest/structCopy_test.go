package main

import (
	"fmt"
	"testing"
)

type User01 struct {
	Name string
	Age int
	Books []int
}

func TestStructCopy(t *testing.T){
	u1 := &User01{
		Name:"landon",
		Age:28,
		Books: []int{1, 3, 4},
	}

	u2 := *u1
	fmt.Printf("%#v\n", u2)

	u2.Age = 33

	u2.Books[0] = 22

	fmt.Printf("%#v\n", u1)
}
