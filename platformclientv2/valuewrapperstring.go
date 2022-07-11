package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Valuewrapperstring
type Valuewrapperstring struct { 
	// Value - The value for the associated field
	Value *string `json:"value,omitempty"`

}

func (o *Valuewrapperstring) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Valuewrapperstring
	
	return json.Marshal(&struct { 
		Value *string `json:"value,omitempty"`
		*Alias
	}{ 
		Value: o.Value,
		Alias:    (*Alias)(o),
	})
}

func (o *Valuewrapperstring) UnmarshalJSON(b []byte) error {
	var ValuewrapperstringMap map[string]interface{}
	err := json.Unmarshal(b, &ValuewrapperstringMap)
	if err != nil {
		return err
	}
	
	if Value, ok := ValuewrapperstringMap["value"].(string); ok {
		o.Value = &Value
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Valuewrapperstring) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
