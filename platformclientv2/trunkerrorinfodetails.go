package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Trunkerrorinfodetails
type Trunkerrorinfodetails struct { 
	// Code
	Code *string `json:"code,omitempty"`


	// Message
	Message *string `json:"message,omitempty"`


	// Hostname
	Hostname *string `json:"hostname,omitempty"`

}

func (o *Trunkerrorinfodetails) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Trunkerrorinfodetails
	
	return json.Marshal(&struct { 
		Code *string `json:"code,omitempty"`
		
		Message *string `json:"message,omitempty"`
		
		Hostname *string `json:"hostname,omitempty"`
		*Alias
	}{ 
		Code: o.Code,
		
		Message: o.Message,
		
		Hostname: o.Hostname,
		Alias:    (*Alias)(o),
	})
}

func (o *Trunkerrorinfodetails) UnmarshalJSON(b []byte) error {
	var TrunkerrorinfodetailsMap map[string]interface{}
	err := json.Unmarshal(b, &TrunkerrorinfodetailsMap)
	if err != nil {
		return err
	}
	
	if Code, ok := TrunkerrorinfodetailsMap["code"].(string); ok {
		o.Code = &Code
	}
    
	if Message, ok := TrunkerrorinfodetailsMap["message"].(string); ok {
		o.Message = &Message
	}
    
	if Hostname, ok := TrunkerrorinfodetailsMap["hostname"].(string); ok {
		o.Hostname = &Hostname
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Trunkerrorinfodetails) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
