package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Schedulegenerationmessage
type Schedulegenerationmessage struct { 
	// VarType - The type of the message
	VarType *string `json:"type,omitempty"`


	// Arguments - The arguments describing the message
	Arguments *[]Schedulermessageargument `json:"arguments,omitempty"`

}

func (o *Schedulegenerationmessage) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Schedulegenerationmessage
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		Arguments *[]Schedulermessageargument `json:"arguments,omitempty"`
		*Alias
	}{ 
		VarType: o.VarType,
		
		Arguments: o.Arguments,
		Alias:    (*Alias)(o),
	})
}

func (o *Schedulegenerationmessage) UnmarshalJSON(b []byte) error {
	var SchedulegenerationmessageMap map[string]interface{}
	err := json.Unmarshal(b, &SchedulegenerationmessageMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := SchedulegenerationmessageMap["type"].(string); ok {
		o.VarType = &VarType
	}
	
	if Arguments, ok := SchedulegenerationmessageMap["arguments"].([]interface{}); ok {
		ArgumentsString, _ := json.Marshal(Arguments)
		json.Unmarshal(ArgumentsString, &o.Arguments)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Schedulegenerationmessage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
