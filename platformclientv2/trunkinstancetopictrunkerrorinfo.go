package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Trunkinstancetopictrunkerrorinfo
type Trunkinstancetopictrunkerrorinfo struct { 
	// Text
	Text *string `json:"text,omitempty"`


	// Code
	Code *string `json:"code,omitempty"`


	// Details
	Details *Trunkinstancetopictrunkerrorinfodetails `json:"details,omitempty"`

}

func (o *Trunkinstancetopictrunkerrorinfo) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Trunkinstancetopictrunkerrorinfo
	
	return json.Marshal(&struct { 
		Text *string `json:"text,omitempty"`
		
		Code *string `json:"code,omitempty"`
		
		Details *Trunkinstancetopictrunkerrorinfodetails `json:"details,omitempty"`
		*Alias
	}{ 
		Text: o.Text,
		
		Code: o.Code,
		
		Details: o.Details,
		Alias:    (*Alias)(o),
	})
}

func (o *Trunkinstancetopictrunkerrorinfo) UnmarshalJSON(b []byte) error {
	var TrunkinstancetopictrunkerrorinfoMap map[string]interface{}
	err := json.Unmarshal(b, &TrunkinstancetopictrunkerrorinfoMap)
	if err != nil {
		return err
	}
	
	if Text, ok := TrunkinstancetopictrunkerrorinfoMap["text"].(string); ok {
		o.Text = &Text
	}
    
	if Code, ok := TrunkinstancetopictrunkerrorinfoMap["code"].(string); ok {
		o.Code = &Code
	}
    
	if Details, ok := TrunkinstancetopictrunkerrorinfoMap["details"].(map[string]interface{}); ok {
		DetailsString, _ := json.Marshal(Details)
		json.Unmarshal(DetailsString, &o.Details)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Trunkinstancetopictrunkerrorinfo) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
