package nvidia

import "encoding/xml"

func ParseNvidiaInfo(output []byte) (*NvidiaSmi, error) {
	res := &NvidiaSmi{}
	err := xml.Unmarshal(output, &res)

	return res, err
}
