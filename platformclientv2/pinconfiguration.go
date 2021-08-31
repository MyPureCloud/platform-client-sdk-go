package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Pinconfiguration
type Pinconfiguration struct { 
	// MinimumLength
	MinimumLength *int `json:"minimumLength,omitempty"`


	// MaximumLength
	MaximumLength *int `json:"maximumLength,omitempty"`

}

func (o *Pinconfiguration) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Pinconfiguration
	
	return json.Marshal(&struct { 
		MinimumLength *int `json:"minimumLength,omitempty"`
		
		MaximumLength *int `json:"maximumLength,omitempty"`
		*Alias
	}{ 
		MinimumLength: o.MinimumLength,
		
		MaximumLength: o.MaximumLength,
		Alias:    (*Alias)(o),
	})
}

func (o *Pinconfiguration) UnmarshalJSON(b []byte) error {
	var PinconfigurationMap map[string]interface{}
	err := json.Unmarshal(b, &PinconfigurationMap)
	if err != nil {
		return err
	}
	
	if MinimumLength, ok := PinconfigurationMap["minimumLength"].(float64); ok {
		MinimumLengthInt := int(MinimumLength)
		o.MinimumLength = &MinimumLengthInt
	}
	
	if MaximumLength, ok := PinconfigurationMap["maximumLength"].(float64); ok {
		MaximumLengthInt := int(MaximumLength)
		o.MaximumLength = &MaximumLengthInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Pinconfiguration) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
