package docker

type opts struct {
	runtime string
}

var RuntimePreference = []string{
	"nvidia",
	"runc",
}

func DefaultOpts() *opts {
	return &opts{}
}

func (o *opts) WithRuntime(rt string) *opts {
	o.runtime = rt
	return o
}
