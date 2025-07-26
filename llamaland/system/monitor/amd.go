package monitor

import (
	"context"
	"io"

	"github.com/wvdschel/llamaland/cmd/maestrod/config"
	"github.com/wvdschel/llamaland/llamaland/runtime/common"
	"github.com/wvdschel/llamaland/llamaland/system/deviceinfo/amd"
)

type amdGPU struct {
	info amd.GPUInfo
}

func (g *amdGPU) ComputeLoad() *Gauge {
	return &Gauge{
		Min:     0,
		Max:     100,
		Current: Value(g.info.CoreUsePercent),
		Unit:    "%",
	}
}

func (g *amdGPU) Memory() *Gauge {
	return &Gauge{
		Min:     0,
		Max:     Value(g.info.MemoryTotal) / 1024 / 1024,
		Current: Value(g.info.MemoryUsed) / 1024 / 1024,
		Unit:    "MiB",
	}
}

func (g *amdGPU) Temperature() *SimpleValue {
	return &SimpleValue{
		Current: Value(g.info.TemperatureC),
		Unit:    "°C",
	}
}

func getAMDGPUs(ctx context.Context, rt common.Runtime) ([]*amdGPU, error) {
	svc := rt.NewService(&config.Service{
		Type: "docker",
		Spec: map[string]any{
			"image": "quay.io/llamaland/rocm-smi:latest",
		},
	})

	if err := svc.Prepare(ctx); err != nil {
		return nil, err
	}
	if err := svc.Start(ctx); err != nil {
		return nil, err
	}

	stdout, stderr, err := svc.Logs(context.Background())
	if err != nil {
		return nil, err
	}

	go func() { _, _ = io.ReadAll(stderr) }()
	logs, err := io.ReadAll(stdout)
	if err != nil {
		return nil, err
	}

	deviceInfo, err := amd.ParseRocmInfo(logs)
	if err != nil {
		return nil, err
	}

	res := []*amdGPU{}
	for _, gpu := range deviceInfo.GPUs {
		res = append(res, &amdGPU{
			info: gpu,
		})
	}

	return res, nil
}
