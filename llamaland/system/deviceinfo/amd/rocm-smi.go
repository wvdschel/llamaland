package amd

import (
	"encoding/json"
	"strconv"
)

// TODO do we even need this jank?
// Seems like everything here is present under /sys/class/drm/card1/device/
// Maybe the sysfs interface is even (partially) cross-vendor compatible?

const (
	TEMP_C                                = "Temperature (Sensor edge) (C)"
	SHADER_CLOCK_SPEED                    = "sclk clock speed:"
	SHADER_CLOCK_LEVEL                    = "sclk clock level:"
	PERF_LEVEL                            = "Performance Level"
	CURRENT_SOCKET_GRAPHICS_PACKAGE_POWER = "Current Socket Graphics Package Power (W)"
	GPU_USE_PERCENTAGE                    = "GPU use (%)"
	VRAM_TOTAL_MEMORY                     = "VRAM Total Memory (B)"
	VRAM_TOTAL_USED_MEMORY                = "VRAM Total Used Memory (B)"
	VIS_VRAM_TOTAL_MEMORY                 = "VIS_VRAM Total Memory (B)"
	VIS_VRAM_TOTAL_USED_MEMORY            = "VIS_VRAM Total Used Memory (B)"
	GTT_TOTAL_MEMORY                      = "GTT Total Memory (B)"
	GTT_TOTAL_USED_MEMORY                 = "GTT Total Used Memory (B)"
	CARD_SERIES                           = "Card Series"
	CARD_MODEL                            = "Card Model"
	CARD_VENDOR                           = "Card Vendor"
	CARD_SKU                              = "Card SKU"
	SUBSYSTEM_ID                          = "Subsystem ID"
	DEVICE_REV                            = "Device Rev"
	NODE_ID                               = "Node ID"
	GUID                                  = "GUID"
	GFX_VERSION                           = "GFX Version"
)

type RocmSmiInfo struct {
	GPUs map[string]GPUInfo
}

type GPUInfo struct {
	Name string

	TemperatureC   float64
	CoreClockMhz   int
	CoreUsePercent int
	PowerDrawW     float64

	MemoryTotal int64
	MemoryUsed  int64

	Properties map[string]string
}

func ParseRocmInfo(output []byte) (*RocmSmiInfo, error) {
	dict := make(map[string]map[string]string)
	if err := json.Unmarshal(output, &dict); err != nil {
		return nil, err
	}

	res := &RocmSmiInfo{
		GPUs: make(map[string]GPUInfo),
	}
	for gpuid, info := range dict {
		gpuInfo := GPUInfo{
			Name:           info[CARD_SERIES],
			TemperatureC:   parseFloat64(info[TEMP_C]),
			CoreClockMhz:   parseClockSpeed(info[SHADER_CLOCK_SPEED]),
			CoreUsePercent: parseInt(info[GPU_USE_PERCENTAGE]),
			PowerDrawW:     parseFloat64(info[CURRENT_SOCKET_GRAPHICS_PACKAGE_POWER]),
			MemoryTotal:    parseInt64(info[VRAM_TOTAL_MEMORY]),
			MemoryUsed:     parseInt64(info[VRAM_TOTAL_USED_MEMORY]),
			Properties:     info,
		}
		res.GPUs[gpuid] = gpuInfo
	}

	return res, nil
}

func parseClockSpeed(val string) int {
	// clock speed values are formatted as '(%dMhz)' for some silly reason
	if val == "" {
		return 0
	}
	// Extract the numeric part before 'Mhz'
	start := 0
	for start < len(val) && val[start] != '(' {
		start++
	}
	start++ // Move past the '('
	end := start
	for end < len(val) && val[end] >= '0' && val[end] <= '9' {
		end++
	}
	if start >= end {
		return 0
	}
	// Parse the numeric part
	numStr := val[start:end]
	num, err := strconv.Atoi(numStr)
	if err != nil {
		return 0
	}
	return num
}
