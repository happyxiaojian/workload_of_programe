package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:   "hello",
		Usage:  "hello world example",
		Action: pri,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func pri(c *cli.Context) error {
	fmt.Println("hello world")
	return nil
}
