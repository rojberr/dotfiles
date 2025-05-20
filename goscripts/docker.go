package goscripts

import ()

/*
Install and setup docker
*/
type Docker struct {
	Logm LogManager
	Err  error
}

func (d Docker) Install(homedir string) error {

	d.Err = d.Logm.RunCommand("brew install docker")
	if d.Err != nil {
		d.Logm.AddLog(Failure, "brew install docker failed")
		return d.Err
	}

	d.Logm.AddLog(Success, "docker successfully installed!")
	return nil
}
