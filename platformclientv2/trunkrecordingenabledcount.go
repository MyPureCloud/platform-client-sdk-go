package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Trunkrecordingenabledcount
type Trunkrecordingenabledcount struct { 
	// EnabledCount - The amount of trunks that have recording enabled
	EnabledCount *int `json:"enabledCount,omitempty"`


	// DisabledCount - The amount of trunks that do not have recording enabled
	DisabledCount *int `json:"disabledCount,omitempty"`

}

func (o *Trunkrecordingenabledcount) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Trunkrecordingenabledcount
	
	return json.Marshal(&struct { 
		EnabledCount *int `json:"enabledCount,omitempty"`
		
		DisabledCount *int `json:"disabledCount,omitempty"`
		*Alias
	}{ 
		EnabledCount: o.EnabledCount,
		
		DisabledCount: o.DisabledCount,
		Alias:    (*Alias)(o),
	})
}

func (o *Trunkrecordingenabledcount) UnmarshalJSON(b []byte) error {
	var TrunkrecordingenabledcountMap map[string]interface{}
	err := json.Unmarshal(b, &TrunkrecordingenabledcountMap)
	if err != nil {
		return err
	}
	
	if EnabledCount, ok := TrunkrecordingenabledcountMap["enabledCount"].(float64); ok {
		EnabledCountInt := int(EnabledCount)
		o.EnabledCount = &EnabledCountInt
	}
	
	if DisabledCount, ok := TrunkrecordingenabledcountMap["disabledCount"].(float64); ok {
		DisabledCountInt := int(DisabledCount)
		o.DisabledCount = &DisabledCountInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Trunkrecordingenabledcount) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
