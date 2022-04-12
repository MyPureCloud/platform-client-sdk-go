package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Supportcentercustommessage
type Supportcentercustommessage struct { 
	// DefaultValue - Default value for the message
	DefaultValue *string `json:"defaultValue,omitempty"`


	// VarType - Type of the message
	VarType *string `json:"type,omitempty"`

}

func (o *Supportcentercustommessage) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Supportcentercustommessage
	
	return json.Marshal(&struct { 
		DefaultValue *string `json:"defaultValue,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		*Alias
	}{ 
		DefaultValue: o.DefaultValue,
		
		VarType: o.VarType,
		Alias:    (*Alias)(o),
	})
}

func (o *Supportcentercustommessage) UnmarshalJSON(b []byte) error {
	var SupportcentercustommessageMap map[string]interface{}
	err := json.Unmarshal(b, &SupportcentercustommessageMap)
	if err != nil {
		return err
	}
	
	if DefaultValue, ok := SupportcentercustommessageMap["defaultValue"].(string); ok {
		o.DefaultValue = &DefaultValue
	}
	
	if VarType, ok := SupportcentercustommessageMap["type"].(string); ok {
		o.VarType = &VarType
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Supportcentercustommessage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
