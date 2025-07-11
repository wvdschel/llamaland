package opencl

import (
	"strconv"
	"strings"
)

type CLDeviceInfo struct {
	Name               string
	Vendor             string
	Version            string
	TotalMemoryBytes   uint64
	MaxAllocationBytes uint64
	UnifiedMemory      bool
	Properties         map[string]string
}

func ParseCLInfoOutput(clinfo string) (map[string]CLDeviceInfo, error) {
	values := make(map[string]map[string]string)
	for _, line := range strings.Split(clinfo, "\n") {
		if len(line) == 0 || line[0] != '[' {
			continue
		}
		words := strings.Fields(line)
		if len(words) < 3 {
			continue
		}
		deviceId := words[0]
		clProperty := words[1]
		value := strings.Join(words[2:], " ")

		if clProperty == "#DEVICES" {
			continue
		}

		if _, ok := values[deviceId]; !ok {
			values[deviceId] = map[string]string{}
		}

		values[deviceId][clProperty] = value
	}
	res := make(map[string]CLDeviceInfo)
	for deviceId, deviceInfo := range values {
		if strings.HasSuffix(deviceId, "/*]") {
			continue
		}

		platformKey := strings.SplitN(deviceId, "/", 2)[0] + "/*]"

		properties := values[platformKey]

		for k, v := range deviceInfo {
			properties[k] = v
		}

		globalMem, _ := strconv.ParseUint(properties["CL_DEVICE_GLOBAL_MEM_SIZE"], 10, 64)
		maxAlloc, _ := strconv.ParseUint(properties["CL_DEVICE_MAX_MEM_ALLOC_SIZE"], 10, 64)
		unifiedMem := properties["CL_DEVICE_HOST_UNIFIED_MEMORY"] == "CL_TRUE"

		res[deviceId] = CLDeviceInfo{
			Name:               properties["CL_DEVICE_NAME"],
			Vendor:             properties["CL_DEVICE_VENDOR"],
			Version:            properties["CL_DEVICE_VERSION"],
			Properties:         properties,
			TotalMemoryBytes:   globalMem,
			MaxAllocationBytes: maxAlloc,
			UnifiedMemory:      unifiedMem,
		}
	}

	return res, nil
}
