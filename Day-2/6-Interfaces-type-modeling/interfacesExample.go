package main

import "fmt"

// type alias
type Helper = interface {
	Help() string
}

type HelpString string

func (hs HelpString) Help() string  {
	return string(hs)
}

type UnHelpString struct {}

func (uhs *UnHelpString) Help() string  {
	return "I cannot help you"
}

// Compile time check
var _ = Helper(HelpString(""))

func main()  {
	fmt.Println(HelpString("Hey").Help())
	//fmt.Println(UnHelpString{}.Help())
	for _, helper := range []Helper{HelpString(""), &UnHelpString{}} {
		fmt.Println(helper.Help())
	}
}
