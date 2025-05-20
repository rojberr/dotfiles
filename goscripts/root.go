package goscripts

import "github.com/bitfield/script"

/*
Root-check to perform the following installations with the needed permissions
*/
type Root struct {
	Logm LogManager
	Err  error
}

func (r Root) Install(homedir string) error {

	//l1 := make([]Logmessage, 0)
	r.Logm.AddLog(Todo, "Checking root privileges")
	_, err := script.Exec("sudo mkdir -p " + homedir + "/.upss").Stdout()
	if err != nil {
		r.Logm.AddLog(Failure, "Wrong sudo pw entered, cancel installation..")
		return err
	}
	r.Logm.AddLog(Success, "Acquired sudo, continue with installation")
	return nil
}
