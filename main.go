package main

import (
	"fmt"
	"strings"
	survey "github.com/AlecAivazis/survey/v2"
	"github.com/bitfield/script"
	"github.com/mgutz/ansi"
	"github.com/rojberr/dotfiles/goscripts"
)

// ASCII art logo for the installer
const logo = `
      ___      __    _____.__.__                 
   __| _/_____/  |__/ ____\__|  |   ____   ______
  / __ |/  _ \   __\   __\|  |  | _/ __ \ /  ___/
 / /_/ (  <_> )  |  |  |  |  |  |_\  ___/ \___ \ 
 \____ |\____/|__|  |__|  |__|____/\___  >____  >
      \/                               \/     \/ `

// Survey questions for install selection
var installQuestions = []*survey.Question{
	{
		Name: "selection",
		Prompt: &survey.MultiSelect{
			Message: "Select components to install:",
			Options: []string{
				fmt.Sprint(goscripts.DockerType),
			},
		},
	},
}

func main() {
	// Color codes for output
	red := ansi.ColorCode("red+bB")
	cyan := ansi.ColorCode("cyan+b")
	reset := ansi.ColorCode("reset")

	// Get user's home directory
	home, _ := script.Exec("bash -c 'echo $HOME'").String()
	home = strings.Replace(home, "\n", "", -1)

	choices := []string{}

	fmt.Println(cyan, logo, reset)
	fmt.Println(red + "Welcome to the " + cyan + "dotfiles" + red + " installer script" + reset)
	fmt.Println("Your homedir: ", home)

	// Prompt user for install options
	err := survey.Ask(installQuestions, &choices, survey.WithPageSize(10))
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Prepare installers based on user selection
	rootInstaller := &goscripts.Root{
		Logm: goscripts.LogManager{},
		Err:  nil,
	}
	installers := []goscripts.Installer{rootInstaller}

	for _, choice := range choices {
		switch choice {
		case goscripts.DockerType.String():
			dockerInstaller := &goscripts.Docker{
				Logm: goscripts.LogManager{},
				Err:  nil,
			}
			installers = append(installers, dockerInstaller)
		default:
			fmt.Printf("❔ Unexpected option selected: %s.\n", choice)
		}
	}

	// Run each installer in sequence
	for _, inst := range installers {
		err = inst.Install(home)
		if err != nil {
			fmt.Println("❌ Installation failed, aborting remaining steps")
			return
		}
		// TODO: Aggregate logs and write to a generated logfile
		// TODO: Open logfile in VSCode after install: code --reuse-window --goto ~/<genfile>.log
	}

	fmt.Println(ansi.ColorCode("magenta+b"), "🎆🎆🎆 Finished! Load bash again using: $ bash 🎆🎆🎆", reset)
}
