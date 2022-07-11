package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Valuewrapperplanningperiodsettings
type Valuewrapperplanningperiodsettings struct { 
	// Value - The value for the associated field
	Value *Planningperiodsettings `json:"value,omitempty"`

}

func (o *Valuewrapperplanningperiodsettings) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Valuewrapperplanningperiodsettings
	
	return json.Marshal(&struct { 
		Value *Planningperiodsettings `json:"value,omitempty"`
		*Alias
	}{ 
		Value: o.Value,
		Alias:    (*Alias)(o),
	})
}

func (o *Valuewrapperplanningperiodsettings) UnmarshalJSON(b []byte) error {
	var ValuewrapperplanningperiodsettingsMap map[string]interface{}
	err := json.Unmarshal(b, &ValuewrapperplanningperiodsettingsMap)
	if err != nil {
		return err
	}
	
	if Value, ok := ValuewrapperplanningperiodsettingsMap["value"].(map[string]interface{}); ok {
		ValueString, _ := json.Marshal(Value)
		json.Unmarshal(ValueString, &o.Value)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Valuewrapperplanningperiodsettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
