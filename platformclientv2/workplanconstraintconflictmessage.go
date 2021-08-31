package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Workplanconstraintconflictmessage
type Workplanconstraintconflictmessage struct { 
	// VarType - Type of constraint conflict that can be resolved by clients in order to generate agent schedules
	VarType *string `json:"type,omitempty"`


	// Arguments - The arguments to the type of the message that can help clients resolve validation issues
	Arguments *[]Workplanvalidationmessageargument `json:"arguments,omitempty"`

}

func (o *Workplanconstraintconflictmessage) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Workplanconstraintconflictmessage
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		Arguments *[]Workplanvalidationmessageargument `json:"arguments,omitempty"`
		*Alias
	}{ 
		VarType: o.VarType,
		
		Arguments: o.Arguments,
		Alias:    (*Alias)(o),
	})
}

func (o *Workplanconstraintconflictmessage) UnmarshalJSON(b []byte) error {
	var WorkplanconstraintconflictmessageMap map[string]interface{}
	err := json.Unmarshal(b, &WorkplanconstraintconflictmessageMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := WorkplanconstraintconflictmessageMap["type"].(string); ok {
		o.VarType = &VarType
	}
	
	if Arguments, ok := WorkplanconstraintconflictmessageMap["arguments"].([]interface{}); ok {
		ArgumentsString, _ := json.Marshal(Arguments)
		json.Unmarshal(ArgumentsString, &o.Arguments)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Workplanconstraintconflictmessage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
