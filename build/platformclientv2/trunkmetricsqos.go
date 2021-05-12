package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Trunkmetricsqos
type Trunkmetricsqos struct { 
	// MismatchCount - Total number of QoS mismatches over the course of the last 24-hour period (sliding window).
	MismatchCount *int `json:"mismatchCount,omitempty"`

}

// String returns a JSON representation of the model
func (o *Trunkmetricsqos) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
