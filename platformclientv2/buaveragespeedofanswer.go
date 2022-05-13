package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Buaveragespeedofanswer - Service goal average speed of answer configuration
type Buaveragespeedofanswer struct { 
	// Include - Whether to include average speed of answer (ASA) in the associated configuration
	Include *bool `json:"include,omitempty"`


	// Seconds - The target average speed of answer (ASA) in seconds. Required if include == true
	Seconds *int `json:"seconds,omitempty"`

}

func (o *Buaveragespeedofanswer) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Buaveragespeedofanswer
	
	return json.Marshal(&struct { 
		Include *bool `json:"include,omitempty"`
		
		Seconds *int `json:"seconds,omitempty"`
		*Alias
	}{ 
		Include: o.Include,
		
		Seconds: o.Seconds,
		Alias:    (*Alias)(o),
	})
}

func (o *Buaveragespeedofanswer) UnmarshalJSON(b []byte) error {
	var BuaveragespeedofanswerMap map[string]interface{}
	err := json.Unmarshal(b, &BuaveragespeedofanswerMap)
	if err != nil {
		return err
	}
	
	if Include, ok := BuaveragespeedofanswerMap["include"].(bool); ok {
		o.Include = &Include
	}
    
	if Seconds, ok := BuaveragespeedofanswerMap["seconds"].(float64); ok {
		SecondsInt := int(Seconds)
		o.Seconds = &SecondsInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Buaveragespeedofanswer) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
