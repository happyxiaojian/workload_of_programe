package main

import (
	"flag"
	"fmt"
)

var b bool

func InitFlag(){
	flag.BoolVar(&b, "open", false, "bool值的设置")
}


func main(){

	InitFlag()

	flag.Parse()

	fmt.Println(b)
}

