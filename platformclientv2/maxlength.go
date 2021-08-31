package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Maxlength
type Maxlength struct { 
	// Min - A non-negative integer for a text-based schema field denoting the minimum largest length string the field can contain for a schema instance.
	Min *int `json:"min,omitempty"`


	// Max - A non-negative integer for a text-based schema field denoting the maximum largest string the field can contain for a schema instance.
	Max *int `json:"max,omitempty"`

}

func (o *Maxlength) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Maxlength
	
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

func (o *Maxlength) UnmarshalJSON(b []byte) error {
	var MaxlengthMap map[string]interface{}
	err := json.Unmarshal(b, &MaxlengthMap)
	if err != nil {
		return err
	}
	
	if Min, ok := MaxlengthMap["min"].(float64); ok {
		MinInt := int(Min)
		o.Min = &MinInt
	}
	
	if Max, ok := MaxlengthMap["max"].(float64); ok {
		MaxInt := int(Max)
		o.Max = &MaxInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Maxlength) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
