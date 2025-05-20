package goscripts

type Installtype int64

const (
	IdleType Installtype = iota
	DockerType
)

func (i Installtype) String() string {
	switch i {
	case DockerType:
		return "install docker"
	}
	return "unknown"
}

type Installer interface {
	Install(homedir string) error
}
