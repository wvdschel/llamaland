package nvidia

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/wvdschel/llamaland/testdata"
)

func TestXMLDecode(t *testing.T) {
	smi := &NvidiaSmi{}

	require.NoError(t, xml.Unmarshal(testdata.NvidiaSmiXML, smi))
	assert.Equal(t, 1, smi.AttachedGPUs)

	require.Equal(t, 1, len(smi.GPU))
	gpu := smi.GPU[0]
	assert.Equal(t, "30.92 W", gpu.GPUPowerReadings.InstantPowerDraw)
	assert.Equal(t, "90.00 W", gpu.GPUPowerReadings.CurrentPowerLimit)
}
