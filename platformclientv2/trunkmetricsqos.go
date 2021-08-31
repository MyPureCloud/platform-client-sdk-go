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

func (o *Trunkmetricsqos) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Trunkmetricsqos
	
	return json.Marshal(&struct { 
		MismatchCount *int `json:"mismatchCount,omitempty"`
		*Alias
	}{ 
		MismatchCount: o.MismatchCount,
		Alias:    (*Alias)(o),
	})
}

func (o *Trunkmetricsqos) UnmarshalJSON(b []byte) error {
	var TrunkmetricsqosMap map[string]interface{}
	err := json.Unmarshal(b, &TrunkmetricsqosMap)
	if err != nil {
		return err
	}
	
	if MismatchCount, ok := TrunkmetricsqosMap["mismatchCount"].(float64); ok {
		MismatchCountInt := int(MismatchCount)
		o.MismatchCount = &MismatchCountInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Trunkmetricsqos) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
