package bits

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

var sumVersionsScenarios = []struct {
	in          string
	expectedSum int
}{
	//{"8A004A801A8002F478", 16},
	//{"620080001611562C8802118E34", 12},
	{"C0015000016115A2E0802F182340", 23},
	//{"A0016C880162017C3686B18A3D4780", 31},
	//{"38006F45291200", 9},
}

func TestSumVersions(t *testing.T) {

	for _, scenario := range sumVersionsScenarios {
		t.Run(fmt.Sprintf("SumVersions(%s)", scenario.in), func(t *testing.T) {
			packet, _, err := readPacket(hexCharChannelToBitChannel(stringToRuneStream(scenario.in)))
			if assert.NoError(t, err) {
				sum := SumVersions(packet)
				assert.Equal(t, scenario.expectedSum, sum)
			}
		})
	}
}

func TestHexCharChannelToBitChannel(t *testing.T) {
	buf := &bytes.Buffer{}
	for bit := range hexCharChannelToBitChannel(stringToRuneStream("38006F45291200")) {
		buf.WriteString(strconv.Itoa(int(bit)))
	}
	assert.Equal(t, "00111000000000000110111101000101001010010001001000000000", buf.String())
}
