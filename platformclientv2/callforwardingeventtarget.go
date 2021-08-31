package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Callforwardingeventtarget
type Callforwardingeventtarget struct { 
	// VarType
	VarType *string `json:"type,omitempty"`


	// Value
	Value *string `json:"value,omitempty"`

}

func (o *Callforwardingeventtarget) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Callforwardingeventtarget
	
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

func (o *Callforwardingeventtarget) UnmarshalJSON(b []byte) error {
	var CallforwardingeventtargetMap map[string]interface{}
	err := json.Unmarshal(b, &CallforwardingeventtargetMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := CallforwardingeventtargetMap["type"].(string); ok {
		o.VarType = &VarType
	}
	
	if Value, ok := CallforwardingeventtargetMap["value"].(string); ok {
		o.Value = &Value
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Callforwardingeventtarget) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
