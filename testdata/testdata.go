package testdata

import _ "embed"

//go:embed nvidia-smi.xml
var NvidiaSmiXML []byte

//go:embed docker-info.json
var DockerInfoJson []byte
