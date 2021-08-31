package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Servicelevel
type Servicelevel struct { 
	// Percentage - The desired Service Level. A value between 0 and 1.
	Percentage *float64 `json:"percentage,omitempty"`


	// DurationMs - Service Level target in milliseconds.
	DurationMs *int `json:"durationMs,omitempty"`

}

func (o *Servicelevel) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Servicelevel
	
	return json.Marshal(&struct { 
		Percentage *float64 `json:"percentage,omitempty"`
		
		DurationMs *int `json:"durationMs,omitempty"`
		*Alias
	}{ 
		Percentage: o.Percentage,
		
		DurationMs: o.DurationMs,
		Alias:    (*Alias)(o),
	})
}

func (o *Servicelevel) UnmarshalJSON(b []byte) error {
	var ServicelevelMap map[string]interface{}
	err := json.Unmarshal(b, &ServicelevelMap)
	if err != nil {
		return err
	}
	
	if Percentage, ok := ServicelevelMap["percentage"].(float64); ok {
		o.Percentage = &Percentage
	}
	
	if DurationMs, ok := ServicelevelMap["durationMs"].(float64); ok {
		DurationMsInt := int(DurationMs)
		o.DurationMs = &DurationMsInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Servicelevel) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
