package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

// This program is used to test set the sub argument

func flagUsage (){
	usageText := "This program is an example cli tool \n" +
		"Usage: \n" +
		"ArgsSub command [arguments] \n" +
		"The commands are \n" +
		"uppercase uppercase a string \n" +
		"lowercase lowercase a string \n" +
		"use \"ArgsSubTest.txt [command] -- help \" for more information about a command "

	fmt.Println(os.Stderr, "%s \n", usageText)
}

func main(){
	flag.Usage = flagUsage
	uppercaseCmd := flag.NewFlagSet("uppercase", flag.ExitOnError)
	lowercaseCmd := flag.NewFlagSet("lowercase", flag.ExitOnError)
	if len(os.Args) == 1 {
		flag.Usage()
		return
	}
	switch os.Args[1]{
	case "uppercase":
		s := uppercaseCmd.String("s", "", "A string of text to be uppercase")
		uppercaseCmd.Parse(os.Args[2:])
		fmt.Println(strings.ToUpper(*s))
	case "lowercase":
		s := lowercaseCmd.String("s", "", "A string of text to be lowercase")
		lowercaseCmd.Parse(os.Args[2:])
		fmt.Println(strings.ToLower(*s))
	default:
		flag.Usage()
	}



}

