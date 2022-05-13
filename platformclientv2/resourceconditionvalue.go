package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Resourceconditionvalue
type Resourceconditionvalue struct { 
	// VarType
	VarType *string `json:"type,omitempty"`


	// Value
	Value *string `json:"value,omitempty"`

}

func (o *Resourceconditionvalue) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Resourceconditionvalue
	
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

func (o *Resourceconditionvalue) UnmarshalJSON(b []byte) error {
	var ResourceconditionvalueMap map[string]interface{}
	err := json.Unmarshal(b, &ResourceconditionvalueMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := ResourceconditionvalueMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Value, ok := ResourceconditionvalueMap["value"].(string); ok {
		o.Value = &Value
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Resourceconditionvalue) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
