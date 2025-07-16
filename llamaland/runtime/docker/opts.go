package docker

import (
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/system"
)

type opts struct {
	systemInfo     *system.Info
	deviceRequests []container.DeviceRequest
	devices        []container.DeviceMapping
	runtime        string
}

var RuntimePreference = []string{
	"nvidia",
	"runc",
}

func DefaultOpts() *opts {
	return &opts{
		devices: []container.DeviceMapping{ // GPU access for Vulkan, ROCm, OpenCL
			{
				PathOnHost:        "/dev/dri",
				PathInContainer:   "/dev/dri",
				CgroupPermissions: "rwm",
			},
			{
				PathOnHost:        "/dev/kfd",
				PathInContainer:   "/dev/kfd",
				CgroupPermissions: "rwm",
			},
		},
	}
}

func (o *opts) info() *system.Info {
	if o.systemInfo == nil {
		s, err := SystemInfo()
		if err != nil {
			panic(err)
		}
		o.systemInfo = &s
	}
	return o.systemInfo
}

func (o *opts) WithRuntime(rt string) *opts {
	o.runtime = rt
	return o
}

func (o *opts) AutodetectRuntime() *opts {
	i := o.info()
	for _, name := range RuntimePreference {
		if _, ok := i.Runtimes[name]; ok {
			o.runtime = name
			break
		}
	}

	if o.runtime == "nvidia" {
		o.deviceRequests = append(o.deviceRequests, container.DeviceRequest{
			Count:        -1, // TODO: -1 is all GPUs, 0 is 1st, etc. (for nvidia runtime)
			Capabilities: [][]string{{"gpu"}},
		})
	}

	return o
}
