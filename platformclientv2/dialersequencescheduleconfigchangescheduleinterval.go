package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Dialersequencescheduleconfigchangescheduleinterval
type Dialersequencescheduleconfigchangescheduleinterval struct { 
	// Start - scheduled start time represented as an ISO-8601 string; for example, yyyy-MM-ddTHH:mm:ss.SSSZ
	Start *string `json:"start,omitempty"`


	// End - scheduled end time represented as an ISO-8601 string; for example, yyyy-MM-ddTHH:mm:ss.SSSZ
	End *string `json:"end,omitempty"`


	// AdditionalProperties
	AdditionalProperties *map[string]interface{} `json:"additionalProperties,omitempty"`

}

func (o *Dialersequencescheduleconfigchangescheduleinterval) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Dialersequencescheduleconfigchangescheduleinterval
	
	return json.Marshal(&struct { 
		Start *string `json:"start,omitempty"`
		
		End *string `json:"end,omitempty"`
		
		AdditionalProperties *map[string]interface{} `json:"additionalProperties,omitempty"`
		*Alias
	}{ 
		Start: o.Start,
		
		End: o.End,
		
		AdditionalProperties: o.AdditionalProperties,
		Alias:    (*Alias)(o),
	})
}

func (o *Dialersequencescheduleconfigchangescheduleinterval) UnmarshalJSON(b []byte) error {
	var DialersequencescheduleconfigchangescheduleintervalMap map[string]interface{}
	err := json.Unmarshal(b, &DialersequencescheduleconfigchangescheduleintervalMap)
	if err != nil {
		return err
	}
	
	if Start, ok := DialersequencescheduleconfigchangescheduleintervalMap["start"].(string); ok {
		o.Start = &Start
	}
    
	if End, ok := DialersequencescheduleconfigchangescheduleintervalMap["end"].(string); ok {
		o.End = &End
	}
    
	if AdditionalProperties, ok := DialersequencescheduleconfigchangescheduleintervalMap["additionalProperties"].(map[string]interface{}); ok {
		AdditionalPropertiesString, _ := json.Marshal(AdditionalProperties)
		json.Unmarshal(AdditionalPropertiesString, &o.AdditionalProperties)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Dialersequencescheduleconfigchangescheduleinterval) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
