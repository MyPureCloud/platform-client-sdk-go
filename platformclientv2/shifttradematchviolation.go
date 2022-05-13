package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Shifttradematchviolation
type Shifttradematchviolation struct { 
	// VarType - The type of constraint violation
	VarType *string `json:"type,omitempty"`


	// Params - Clarifying user params for constructing helpful error messages
	Params *map[string]string `json:"params,omitempty"`

}

func (o *Shifttradematchviolation) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Shifttradematchviolation
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		Params *map[string]string `json:"params,omitempty"`
		*Alias
	}{ 
		VarType: o.VarType,
		
		Params: o.Params,
		Alias:    (*Alias)(o),
	})
}

func (o *Shifttradematchviolation) UnmarshalJSON(b []byte) error {
	var ShifttradematchviolationMap map[string]interface{}
	err := json.Unmarshal(b, &ShifttradematchviolationMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := ShifttradematchviolationMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Params, ok := ShifttradematchviolationMap["params"].(map[string]interface{}); ok {
		ParamsString, _ := json.Marshal(Params)
		json.Unmarshal(ParamsString, &o.Params)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Shifttradematchviolation) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
