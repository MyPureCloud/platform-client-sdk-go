package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Workplanvalidationmessageargument
type Workplanvalidationmessageargument struct { 
	// VarType - The type of the argument associated with violation messages
	VarType *string `json:"type,omitempty"`


	// Value - The value of the argument
	Value *string `json:"value,omitempty"`

}

func (o *Workplanvalidationmessageargument) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Workplanvalidationmessageargument
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		Value *string `json:"value,omitempty"`
		*Alias
	}{ 
		VarType: o.VarType,
		
		Value: o.Value,
		Alias:    (*Alias)(o),
	})
}

func (o *Workplanvalidationmessageargument) UnmarshalJSON(b []byte) error {
	var WorkplanvalidationmessageargumentMap map[string]interface{}
	err := json.Unmarshal(b, &WorkplanvalidationmessageargumentMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := WorkplanvalidationmessageargumentMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Value, ok := WorkplanvalidationmessageargumentMap["value"].(string); ok {
		o.Value = &Value
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Workplanvalidationmessageargument) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
