package main

import "fmt"

import golangexamples "github.com/hammadnadeemx/golangexample"

func main() {
	var carr = "letsencryptthistext"
	var temp string
	temp = golangexamples.ConcatSlice(carr)
	fmt.Println(temp)
	temp = golangexamples.Encrypt(carr, 5)
	fmt.Println(temp)
	temp = golangexamples.EZGreetings("hammy")
	fmt.Println(temp)
}
