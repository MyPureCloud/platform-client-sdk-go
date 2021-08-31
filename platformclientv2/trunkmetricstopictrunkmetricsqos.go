package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Trunkmetricstopictrunkmetricsqos
type Trunkmetricstopictrunkmetricsqos struct { 
	// MismatchCount
	MismatchCount *int `json:"mismatchCount,omitempty"`

}

func (o *Trunkmetricstopictrunkmetricsqos) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Trunkmetricstopictrunkmetricsqos
	
	return json.Marshal(&struct { 
		MismatchCount *int `json:"mismatchCount,omitempty"`
		*Alias
	}{ 
		MismatchCount: o.MismatchCount,
		Alias:    (*Alias)(o),
	})
}

func (o *Trunkmetricstopictrunkmetricsqos) UnmarshalJSON(b []byte) error {
	var TrunkmetricstopictrunkmetricsqosMap map[string]interface{}
	err := json.Unmarshal(b, &TrunkmetricstopictrunkmetricsqosMap)
	if err != nil {
		return err
	}
	
	if MismatchCount, ok := TrunkmetricstopictrunkmetricsqosMap["mismatchCount"].(float64); ok {
		MismatchCountInt := int(MismatchCount)
		o.MismatchCount = &MismatchCountInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Trunkmetricstopictrunkmetricsqos) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
