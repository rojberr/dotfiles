package main

import (
	"fmt"
	"strings"
	"github.com/mgutz/ansi"
	"github.com/bitfield/script"
)


// Use asciiart logo
const logo = `
      ___      __    _____.__.__                 
   __| _/_____/  |__/ ____\__|  |   ____   ______
  / __ |/  _ \   __\   __\|  |  | _/ __ \ /  ___/
 / /_/ (  <_> )  |  |  |  |  |  |_\  ___/ \___ \ 
 \____ |\____/|__|  |__|  |__|____/\___  >____  >
      \/                               \/     \/ `



func main() {

	// Set text colors to use later
	var pan = ansi.ColorCode("red+bB")
	var cyan = ansi.ColorCode("cyan+b")
	var reset = ansi.ColorCode("reset")

	// Get the current working directory
	homedir, _ := script.Exec("bash -c 'echo $HOME'").String()
	homedir = strings.Replace(homedir, "\n", "", -1)

	fmt.Println(pan + "Welcome to the " + cyan + "Go" + pan + " script" + reset)
	fmt.Println(homedir)
	fmt.Println(logo)
}
