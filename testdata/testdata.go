package testdata

import _ "embed"

//go:embed nvidia-smi.xml
var NvidiaSmiXML []byte

//go:embed docker-info-nvidia.json
var DockerInfoNvidiaJson []byte

//go:embed podman-info-amd.json
var DockerInfoAMDJson []byte

//go:embed clinfo.txt
var CLInfoOutput []byte

//go:embed rocm-smi.json
var RocmSmiJSON []byte
