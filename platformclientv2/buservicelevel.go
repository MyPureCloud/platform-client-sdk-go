package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Buservicelevel - Service goal service level configuration
type Buservicelevel struct { 
	// Include - Whether to include service level targets in the associated configuration
	Include *bool `json:"include,omitempty"`


	// Percent - Service level target percent answered. Required if include == true
	Percent *int `json:"percent,omitempty"`


	// Seconds - Service level target answer time. Required if include == true
	Seconds *int `json:"seconds,omitempty"`

}

func (o *Buservicelevel) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Buservicelevel
	
	return json.Marshal(&struct { 
		Include *bool `json:"include,omitempty"`
		
		Percent *int `json:"percent,omitempty"`
		
		Seconds *int `json:"seconds,omitempty"`
		*Alias
	}{ 
		Include: o.Include,
		
		Percent: o.Percent,
		
		Seconds: o.Seconds,
		Alias:    (*Alias)(o),
	})
}

func (o *Buservicelevel) UnmarshalJSON(b []byte) error {
	var BuservicelevelMap map[string]interface{}
	err := json.Unmarshal(b, &BuservicelevelMap)
	if err != nil {
		return err
	}
	
	if Include, ok := BuservicelevelMap["include"].(bool); ok {
		o.Include = &Include
	}
	
	if Percent, ok := BuservicelevelMap["percent"].(float64); ok {
		PercentInt := int(Percent)
		o.Percent = &PercentInt
	}
	
	if Seconds, ok := BuservicelevelMap["seconds"].(float64); ok {
		SecondsInt := int(Seconds)
		o.Seconds = &SecondsInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Buservicelevel) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
