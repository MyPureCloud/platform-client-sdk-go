package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Eventtyping - A Typing event.
type Eventtyping struct { 
	// VarType - Describes the type of Typing event.
	VarType *string `json:"type,omitempty"`


	// Duration - The duration of the Typing event in milliseconds.
	Duration *int `json:"duration,omitempty"`

}

func (o *Eventtyping) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Eventtyping
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		Duration *int `json:"duration,omitempty"`
		*Alias
	}{ 
		VarType: o.VarType,
		
		Duration: o.Duration,
		Alias:    (*Alias)(o),
	})
}

func (o *Eventtyping) UnmarshalJSON(b []byte) error {
	var EventtypingMap map[string]interface{}
	err := json.Unmarshal(b, &EventtypingMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := EventtypingMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Duration, ok := EventtypingMap["duration"].(float64); ok {
		DurationInt := int(Duration)
		o.Duration = &DurationInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Eventtyping) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
