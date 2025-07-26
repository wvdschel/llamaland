package monitor

import (
	"context"

	"github.com/wvdschel/llamaland/llamaland/runtime/common"
)

type GPU interface {
	ComputeLoad() *Gauge
	Memory() *Gauge
	Temperature() *SimpleValue
}

func ListGPUs(ctx context.Context, rt common.Runtime) ([]GPU, error) {
	amdgpus, err := getAMDGPUs(ctx, rt)
	if err != nil {
		return nil, err
	}

	nvgpus, err := getNvidiaGPUs(ctx, rt)
	if err != nil {
		return nil, err
	}

	gpus := make([]GPU, len(amdgpus)+len(nvgpus))

	for idx, g := range amdgpus {
		gpus[idx] = g
	}

	for idx, g := range nvgpus {
		gpus[idx+len(amdgpus)] = g
	}

	return gpus, nil
}
