package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Replacementterm
type Replacementterm struct { 
	// VarType
	VarType *string `json:"type,omitempty"`


	// ExistingValue
	ExistingValue *string `json:"existingValue,omitempty"`


	// UpdatedValue
	UpdatedValue *string `json:"updatedValue,omitempty"`

}

func (o *Replacementterm) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Replacementterm
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		ExistingValue *string `json:"existingValue,omitempty"`
		
		UpdatedValue *string `json:"updatedValue,omitempty"`
		*Alias
	}{ 
		VarType: o.VarType,
		
		ExistingValue: o.ExistingValue,
		
		UpdatedValue: o.UpdatedValue,
		Alias:    (*Alias)(o),
	})
}

func (o *Replacementterm) UnmarshalJSON(b []byte) error {
	var ReplacementtermMap map[string]interface{}
	err := json.Unmarshal(b, &ReplacementtermMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := ReplacementtermMap["type"].(string); ok {
		o.VarType = &VarType
	}
	
	if ExistingValue, ok := ReplacementtermMap["existingValue"].(string); ok {
		o.ExistingValue = &ExistingValue
	}
	
	if UpdatedValue, ok := ReplacementtermMap["updatedValue"].(string); ok {
		o.UpdatedValue = &UpdatedValue
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Replacementterm) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
