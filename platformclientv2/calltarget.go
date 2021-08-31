package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Calltarget
type Calltarget struct { 
	// VarType - The type of call
	VarType *string `json:"type,omitempty"`


	// Value - The id of the station or an E.164 formatted phone number
	Value *string `json:"value,omitempty"`

}

func (o *Calltarget) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Calltarget
	
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

func (o *Calltarget) UnmarshalJSON(b []byte) error {
	var CalltargetMap map[string]interface{}
	err := json.Unmarshal(b, &CalltargetMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := CalltargetMap["type"].(string); ok {
		o.VarType = &VarType
	}
	
	if Value, ok := CalltargetMap["value"].(string); ok {
		o.Value = &Value
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Calltarget) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
