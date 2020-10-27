package main

import (
	"errors"
	"fmt"
	"github.com/davecgh/go-spew/spew"

	"github.com/reactivex/rxgo/v2"
)

func main01() {
	observable := rxgo.Just(1, 2, errors.New("unknown"), 3, 4, 5)()
	ch := observable.Observe()
	for item := range ch {
		if item.Error() {
			fmt.Println("error:", item.E)
		} else {
			fmt.Println(item.V)
		}
	}
}

func main() {
	observable := rxgo.Just(1, 2, 3,errors.New("unknown"), 4, 5)()
	x := <-observable.ForEach(func(v interface{}) {
		fmt.Println("received:", v)
	}, func(err error) {
		fmt.Println("error:", err)
	}, func() {
		fmt.Println("completed")
	})

	spew.Dump(x)
}


