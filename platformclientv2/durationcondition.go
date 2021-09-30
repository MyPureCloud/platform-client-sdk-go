package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Durationcondition
type Durationcondition struct { 
	// DurationTarget
	DurationTarget *string `json:"durationTarget,omitempty"`


	// DurationOperator
	DurationOperator *string `json:"durationOperator,omitempty"`


	// DurationRange
	DurationRange *string `json:"durationRange,omitempty"`


	// DurationMode
	DurationMode *string `json:"durationMode,omitempty"`

}

func (o *Durationcondition) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Durationcondition
	
	return json.Marshal(&struct { 
		DurationTarget *string `json:"durationTarget,omitempty"`
		
		DurationOperator *string `json:"durationOperator,omitempty"`
		
		DurationRange *string `json:"durationRange,omitempty"`
		
		DurationMode *string `json:"durationMode,omitempty"`
		*Alias
	}{ 
		DurationTarget: o.DurationTarget,
		
		DurationOperator: o.DurationOperator,
		
		DurationRange: o.DurationRange,
		
		DurationMode: o.DurationMode,
		Alias:    (*Alias)(o),
	})
}

func (o *Durationcondition) UnmarshalJSON(b []byte) error {
	var DurationconditionMap map[string]interface{}
	err := json.Unmarshal(b, &DurationconditionMap)
	if err != nil {
		return err
	}
	
	if DurationTarget, ok := DurationconditionMap["durationTarget"].(string); ok {
		o.DurationTarget = &DurationTarget
	}
	
	if DurationOperator, ok := DurationconditionMap["durationOperator"].(string); ok {
		o.DurationOperator = &DurationOperator
	}
	
	if DurationRange, ok := DurationconditionMap["durationRange"].(string); ok {
		o.DurationRange = &DurationRange
	}
	
	if DurationMode, ok := DurationconditionMap["durationMode"].(string); ok {
		o.DurationMode = &DurationMode
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Durationcondition) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
