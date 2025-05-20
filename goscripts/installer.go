package goscripts

type Installtype int64

const (
	IdleType Installtype = iota
	DockerType
	GitType
)

func (i Installtype) String() string {
	switch i {
	case DockerType:
		return "install docker"
	case GitType:
		return "install&configure git"
	}
	return "unknown"
}

type Installer interface {
	Install(homedir string) error
}
