package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Minlength
type Minlength struct { 
	// Min - A non-negative integer for a text-based schema field denoting the minimum smallest length a string field can contain for a schema instance.
	Min *int `json:"min,omitempty"`


	// Max - A non-negative integer for a text-based schema field denoting the maximum smallest length string the field can contain for a schema instance.
	Max *int `json:"max,omitempty"`

}

func (o *Minlength) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Minlength
	
	return json.Marshal(&struct { 
		Min *int `json:"min,omitempty"`
		
		Max *int `json:"max,omitempty"`
		*Alias
	}{ 
		Min: o.Min,
		
		Max: o.Max,
		Alias:    (*Alias)(o),
	})
}

func (o *Minlength) UnmarshalJSON(b []byte) error {
	var MinlengthMap map[string]interface{}
	err := json.Unmarshal(b, &MinlengthMap)
	if err != nil {
		return err
	}
	
	if Min, ok := MinlengthMap["min"].(float64); ok {
		MinInt := int(Min)
		o.Min = &MinInt
	}
	
	if Max, ok := MinlengthMap["max"].(float64); ok {
		MaxInt := int(Max)
		o.Max = &MaxInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Minlength) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
