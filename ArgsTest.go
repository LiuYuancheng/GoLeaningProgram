package main

import (
	"flag"
	"fmt"
	"os"
)

func main(){
	flag.Usage = func(){
		usageText :=  "Usage ArgsTest.go [Option]" +
			"An exmaple of customizing usage output" +
			"-s , --s 		example string arguments, default: string help text" +
			"-i , --i 		example integer argument, default: Int help text" +
			"-b , --b 		example boolean argument, default: bool help text"
			fmt.Println(os.Stderr, "%s \n", usageText)

	}
	s := flag.String("s", "testing", "String help text")
	i := flag.Int("i", 1, "Int help text")
	b := flag.Bool("b", false, "boolean help text")
	flag.Parse()
	fmt.Println("value of s:", *s)
	fmt.Println("value of i", *i)
	fmt.Println("value of b", *b)
	for i, args := range os.Args{
		fmt.Println("Argument", i, "is", args)
	}
}
