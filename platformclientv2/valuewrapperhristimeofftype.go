package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Valuewrapperhristimeofftype
type Valuewrapperhristimeofftype struct { 
	// Value - The value for the associated field
	Value *Hristimeofftype `json:"value,omitempty"`

}

func (o *Valuewrapperhristimeofftype) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Valuewrapperhristimeofftype
	
	return json.Marshal(&struct { 
		Value *Hristimeofftype `json:"value,omitempty"`
		*Alias
	}{ 
		Value: o.Value,
		Alias:    (*Alias)(o),
	})
}

func (o *Valuewrapperhristimeofftype) UnmarshalJSON(b []byte) error {
	var ValuewrapperhristimeofftypeMap map[string]interface{}
	err := json.Unmarshal(b, &ValuewrapperhristimeofftypeMap)
	if err != nil {
		return err
	}
	
	if Value, ok := ValuewrapperhristimeofftypeMap["value"].(map[string]interface{}); ok {
		ValueString, _ := json.Marshal(Value)
		json.Unmarshal(ValueString, &o.Value)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Valuewrapperhristimeofftype) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
