package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Trunkmetricsqos
type Trunkmetricsqos struct { 
	// MismatchCount - Total number of QoS mismatches over the course of the last 24-hour period (sliding window).
	MismatchCount *int `json:"mismatchCount,omitempty"`

}

func (u *Trunkmetricsqos) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Trunkmetricsqos

	

	return json.Marshal(&struct { 
		MismatchCount *int `json:"mismatchCount,omitempty"`
		*Alias
	}{ 
		MismatchCount: u.MismatchCount,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Trunkmetricsqos) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
