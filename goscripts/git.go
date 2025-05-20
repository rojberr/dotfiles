package goscripts

import (
	"fmt"
)

/*
Install git and change git config
This method performs:
  - installs git package
  - changes git config
  - sets up git aliases
*/
type Git struct {
	Logm LogManager
	Err  error
}

func (d Git) Install(homedir string) error {

	d.Err = d.Logm.RunCommand("brew install git")
	if d.Err != nil {
		d.Logm.AddLog(Failure, "brew install git failed")
		return d.Err
	}

	// Prompt for user name and email
	var gitName, gitEmail string
	fmt.Print("Enter your git user.name: ")
	fmt.Scanln(&gitName)
	fmt.Print("Enter your git user.email: ")
	fmt.Scanln(&gitEmail)

	d.Err = d.Logm.RunCommand(fmt.Sprintf("git config --global user.name '%s'", gitName))
	if d.Err != nil {
		d.Logm.AddLog(Failure, "Failed to set git user.name")
		return d.Err
	}

	d.Err = d.Logm.RunCommand(fmt.Sprintf("git config --global user.email '%s'", gitEmail))
	if d.Err != nil {
		d.Logm.AddLog(Failure, "Failed to set git user.email")
		return d.Err
	}

	// TODO: Add git aliases

	d.Logm.AddLog(Success, "git successfully installed and configured!")
	return nil
}
