package main

import (
	"fmt"
	"testing"
	"time"
)

// 测试for + select 的刷新频率
func TestSelect(t *testing.T){
	//tt := time.Tick(time.Second * 3)
	//c := make(chan bool, 0)

	var i int
	var j int

	go func() {
		for {
			j++
			time.Sleep(time.Second)
			fmt.Printf("g2 --> %d\n", j)
		}
	}()

	//for {
		i++
		fmt.Printf("i -> %d\n", i)
		fmt.Println(time.Now().UnixNano())
		select{
		//case <- c:
		//	fmt.Printf("close chan\n")
		//case <-tt:
		//	fmt.Printf("from time.Ticker\n")

		}
	//}
}

func TestSelect01(t *testing.T){
	c := make(chan int)
	go func(){
		for range time.Tick(time.Second){
			c <- 1
		}
	}()

	for {
		select {
		case <- c:
			fmt.Println("case1")
		case <- c:
			fmt.Println("case2")
		case <- c:
			fmt.Println("case3")
		}
	}
}
