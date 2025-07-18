package opencl

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/wvdschel/llamaland/testdata"
)

func Test_CLInfoParser(t *testing.T) {
	deviceInfo, err := ParseCLInfoOutput(string(testdata.CLInfoOutput))

	require.NoError(t, err)

	for deviceID, info := range deviceInfo {
		switch deviceID {
		case "[AMD/0]":
			assert.Equal(t, "gfx1151", info.Name)
			assert.Equal(t, "Advanced Micro Devices, Inc.", info.Vendor)
			assert.Equal(t, "OpenCL 2.0", info.Version)
			assert.Equal(t, uint64(33323491328), info.TotalMemoryBytes)
			assert.Equal(t, uint64(28324967624), info.MaxAllocationBytes)
			assert.True(t, info.UnifiedMemory)
		case "[MESA/0]":
			assert.Equal(t, "AMD Radeon Graphics (radeonsi, gfx1151, LLVM 19.1.7, DRM 3.63, 6.15.5-arch1-1)", info.Name)
			assert.Equal(t, "AMD", info.Vendor)
			assert.Equal(t, "OpenCL 3.0", info.Version)
			assert.Equal(t, uint64(536870912), info.TotalMemoryBytes)
			assert.Equal(t, uint64(2147483648), info.MaxAllocationBytes)
			assert.False(t, info.UnifiedMemory)
		case "[MESA/1]":
			assert.Equal(t, "llvmpipe (LLVM 19.1.7, 256 bits)", info.Name)
			assert.Equal(t, "Mesa", info.Vendor)
			assert.Equal(t, "OpenCL 3.0", info.Version)
			assert.Equal(t, uint64(66646982656), info.TotalMemoryBytes)
			assert.Equal(t, uint64(2147483648), info.MaxAllocationBytes)
			assert.True(t, info.UnifiedMemory)
		default:
			t.Fatalf("unexpected device ID: %s", deviceID)
		}
	}
}
