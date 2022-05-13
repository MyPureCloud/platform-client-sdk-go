package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Trunkerrorinfo
type Trunkerrorinfo struct { 
	// Text
	Text *string `json:"text,omitempty"`


	// Code
	Code *string `json:"code,omitempty"`


	// Details
	Details *Trunkerrorinfodetails `json:"details,omitempty"`

}

func (o *Trunkerrorinfo) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Trunkerrorinfo
	
	return json.Marshal(&struct { 
		Text *string `json:"text,omitempty"`
		
		Code *string `json:"code,omitempty"`
		
		Details *Trunkerrorinfodetails `json:"details,omitempty"`
		*Alias
	}{ 
		Text: o.Text,
		
		Code: o.Code,
		
		Details: o.Details,
		Alias:    (*Alias)(o),
	})
}

func (o *Trunkerrorinfo) UnmarshalJSON(b []byte) error {
	var TrunkerrorinfoMap map[string]interface{}
	err := json.Unmarshal(b, &TrunkerrorinfoMap)
	if err != nil {
		return err
	}
	
	if Text, ok := TrunkerrorinfoMap["text"].(string); ok {
		o.Text = &Text
	}
    
	if Code, ok := TrunkerrorinfoMap["code"].(string); ok {
		o.Code = &Code
	}
    
	if Details, ok := TrunkerrorinfoMap["details"].(map[string]interface{}); ok {
		DetailsString, _ := json.Marshal(Details)
		json.Unmarshal(DetailsString, &o.Details)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Trunkerrorinfo) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
