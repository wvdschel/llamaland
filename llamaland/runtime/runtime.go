package runtime

import (
	"fmt"

	"github.com/wvdschel/llamaland/llamaland/runtime/common"
	"github.com/wvdschel/llamaland/llamaland/runtime/docker"
)

func ForServiceType(svcType string) (common.Runtime, error) {
	switch svcType {
	case "docker-nvidia":
		return docker.NewRuntime(docker.DefaultOpts().WithRuntime("nvidia"))
	case "docker":
		return docker.NewRuntime(docker.DefaultOpts())
	default:
		return nil, fmt.Errorf("unsupported service type '%s'", svcType)
	}
}
