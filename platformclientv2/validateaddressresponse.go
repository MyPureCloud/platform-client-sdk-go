package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Validateaddressresponse
type Validateaddressresponse struct { 
	// Valid - Was the passed in address valid
	Valid *bool `json:"valid,omitempty"`


	// Response - Subscriber schema
	Response *Subscriberresponse `json:"response,omitempty"`

}

func (o *Validateaddressresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Validateaddressresponse
	
	return json.Marshal(&struct { 
		Valid *bool `json:"valid,omitempty"`
		
		Response *Subscriberresponse `json:"response,omitempty"`
		*Alias
	}{ 
		Valid: o.Valid,
		
		Response: o.Response,
		Alias:    (*Alias)(o),
	})
}

func (o *Validateaddressresponse) UnmarshalJSON(b []byte) error {
	var ValidateaddressresponseMap map[string]interface{}
	err := json.Unmarshal(b, &ValidateaddressresponseMap)
	if err != nil {
		return err
	}
	
	if Valid, ok := ValidateaddressresponseMap["valid"].(bool); ok {
		o.Valid = &Valid
	}
    
	if Response, ok := ValidateaddressresponseMap["response"].(map[string]interface{}); ok {
		ResponseString, _ := json.Marshal(Response)
		json.Unmarshal(ResponseString, &o.Response)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Validateaddressresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
