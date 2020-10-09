package main

import (
	"fmt"

	"github.com/fatih/color"
	color1 "github.com/gookit/color"
)

func main01(){

	// Use handy standard colors
	color.Set(color.FgYellow)

	fmt.Println("Existing text will now be in yellow")
	fmt.Printf("This one %s\n", "too")

	color.Unset() // Don't forget to unset

	// You can mix up parameters
	color.Set(color.FgMagenta, color.Bold)
	defer color.Unset() // Use it in your function

	fmt.Println("All text will now be bold magenta.")
}

func main02(){
	// Create a new color object
	c := color.New(color.FgCyan).Add(color.Underline)
	_, _ = c.Println("Prints cyan text with an underline.")

	// Or just add them to New()
	d := color.New(color.FgCyan, color.Bold)
	_, _ = d.Printf("This prints bold cyan %s\n", "too!.")

	// Mix up foreground and background colors, create new mixes!
	red := color.New(color.FgRed)

	boldRed := red.Add(color.Bold)
	_, _ = boldRed.Println("This will print text in bold red.")

	whiteBackground := red.Add(color.BgWhite)
	_, _ = whiteBackground.Println("Red text with white background.")

}

func main() {
	// simple usage
	color1.Cyan.Printf("Simple to use %s\n", "color")

	// use like func
	red := color1.FgRed.Render
	green := color1.FgGreen.Render
	fmt.Printf("%s line %s library\n", red("Command"), green("color"))

	// 自定义颜色
	color1.New(color1.FgWhite, color1.BgBlack).Println("custom color style")

	// 也可以:
	color1.Style{color1.FgCyan, color1.OpBold}.Println("custom color style")

	// internal style:
	color1.Info.Println("message")
	color1.Warn.Println("message")
	color1.Error.Println("message")

	// 使用颜色标签
	color1.Print("<suc>he</><comment>llo</>, <cyan>wel</><red>come</>\n")

	// apply a style tag
	color1.Tag("info").Println("info style text")

	// prompt message
	color1.Info.Prompt("prompt style message")
	color1.Warn.Prompt("prompt style message")

	// tips message
	color1.Info.Tips("tips style message")
	color1.Warn.Tips("tips style message")
}