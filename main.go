package main

import (
	"fmt"
	survey "github.com/AlecAivazis/survey/v2"
	"github.com/bitfield/script"
	"github.com/mgutz/ansi"
	"strings"
)

// Use asciiart logo
const logo = `
      ___      __    _____.__.__                 
   __| _/_____/  |__/ ____\__|  |   ____   ______
  / __ |/  _ \   __\   __\|  |  | _/ __ \ /  ___/
 / /_/ (  <_> )  |  |  |  |  |  |_\  ___/ \___ \ 
 \____ |\____/|__|  |__|  |__|____/\___  >____  >
      \/                               \/     \/ `

// Wizard with install selections
var multiQs = []*survey.Question{
	{
		Name: "letter",
		Prompt: &survey.MultiSelect{
			Message: "What you want to install :",
			Options: []string{
				fmt.Sprint("aaa"),
				fmt.Sprint("bbb"),
			},
		},
	},
}

func main() {

	// Set text colors to use later
	var pan = ansi.ColorCode("red+bB")
	var cyan = ansi.ColorCode("cyan+b")
	var reset = ansi.ColorCode("reset")

	// Get the current working directory
	homedir, _ := script.Exec("bash -c 'echo $HOME'").String()
	homedir = strings.Replace(homedir, "\n", "", -1)
	options := []string{}

	fmt.Println(logo)
	fmt.Println(pan + "Welcome to the " + cyan + "dotfiles" + pan + " installer script" + reset)
	fmt.Println("Your homedir: ", homedir)

	// Show wizard with install selections
	err := survey.Ask(multiQs, &options, survey.WithPageSize(10))

	if err != nil {
		fmt.Println(err.Error())
		return
	}

}
