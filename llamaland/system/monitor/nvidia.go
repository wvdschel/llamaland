package monitor

import (
	"context"
	"fmt"
	"io"

	"github.com/wvdschel/llamaland/cmd/maestrod/config"
	"github.com/wvdschel/llamaland/llamaland/runtime/common"
	"github.com/wvdschel/llamaland/llamaland/system/deviceinfo/nvidia"
)

type nvidiaGPU struct {
	info nvidia.GPU
}

func (ng *nvidiaGPU) ComputeLoad() *Gauge {
	util := parsePercentage(ng.info.Utilization.GpuUtil)
	return &Gauge{
		Min:     0,
		Max:     100,
		Current: util,
		Unit:    "%",
	}
}

func (ng *nvidiaGPU) Memory() *Gauge {
	capacity := parseBinarySize(ng.info.FbMemoryUsage.Total)
	util := parseBinarySize(ng.info.FbMemoryUsage.Used)
	return &Gauge{
		Min:     0,
		Max:     capacity,
		Current: util / 1024 / 1024,
		Unit:    "MiB",
	}
}

func (ng *nvidiaGPU) Temperature() *SimpleValue {
	// Implement the logic to get the temperature from ng.info
	return &SimpleValue{
		Current: parseTemp(ng.info.Temperature.GpuTemp),
		Unit:    "°C",
	}
}

func getNvidiaGPUs(ctx context.Context, rt common.Runtime) ([]*nvidiaGPU, error) {
	svc := rt.NewService(&config.Service{
		Type: "docker",
		Spec: map[string]any{
			"image": "quay.io/llamaland/nvidia-smi:latest",
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

	deviceInfo, err := nvidia.ParseNvidiaInfo(logs)
	if err != nil {
		return nil, err
	}

	res := []*nvidiaGPU{}
	for _, gpu := range deviceInfo.GPU {
		res = append(res, &nvidiaGPU{
			info: gpu,
		})
	}

	return res, nil
}

func parsePercentage(v string) Value {
	// Parse v, which is a string formatted as "X %", where X is a number
	var percentage float64
	_, err := fmt.Sscanf(v, "%f %%", &percentage)
	if err != nil {
		return 0.0 // or handle the error as appropriate
	}
	return Value(percentage)
}

func parseTemp(v string) Value {
	// Parse v, which is a string formatted as "X %", where X is a number
	var res float64
	_, err := fmt.Sscanf(v, "%f C", &res)
	if err != nil {
		return 0.0 // or handle the error as appropriate
	}
	return Value(res)
}

func parseBinarySize(v string) Value {
	// Parse v, which is formatted as "%d MiB" or "%d GiB"
	var value float64
	var unit string
	_, err := fmt.Sscanf(v, "%f %s", &value, &unit)
	if err != nil {
		return 0.0
	}
	switch unit {
	case "KiB":
		return Value(value * 1024)
	case "MiB":
		return Value(value * 1024 * 1024)
	case "GiB":
		return Value(value * 1024 * 1024 * 1024)
	case "TiB":
		return Value(value * 1024 * 1024 * 1024 * 1024)
	default:
		return 0.0
	}
}
