package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Fieldlist
type Fieldlist struct { 
	// CustomLabels
	CustomLabels *bool `json:"customLabels,omitempty"`


	// InstructionText
	InstructionText *string `json:"instructionText,omitempty"`


	// Key
	Key *string `json:"key,omitempty"`


	// LabelKeys
	LabelKeys *[]string `json:"labelKeys,omitempty"`


	// Params
	Params *map[string]interface{} `json:"params,omitempty"`


	// Repeatable
	Repeatable *bool `json:"repeatable,omitempty"`


	// State
	State *string `json:"state,omitempty"`


	// VarType
	VarType *string `json:"type,omitempty"`


	// Required
	Required *bool `json:"required,omitempty"`

}

func (o *Fieldlist) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Fieldlist
	
	return json.Marshal(&struct { 
		CustomLabels *bool `json:"customLabels,omitempty"`
		
		InstructionText *string `json:"instructionText,omitempty"`
		
		Key *string `json:"key,omitempty"`
		
		LabelKeys *[]string `json:"labelKeys,omitempty"`
		
		Params *map[string]interface{} `json:"params,omitempty"`
		
		Repeatable *bool `json:"repeatable,omitempty"`
		
		State *string `json:"state,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		Required *bool `json:"required,omitempty"`
		*Alias
	}{ 
		CustomLabels: o.CustomLabels,
		
		InstructionText: o.InstructionText,
		
		Key: o.Key,
		
		LabelKeys: o.LabelKeys,
		
		Params: o.Params,
		
		Repeatable: o.Repeatable,
		
		State: o.State,
		
		VarType: o.VarType,
		
		Required: o.Required,
		Alias:    (*Alias)(o),
	})
}

func (o *Fieldlist) UnmarshalJSON(b []byte) error {
	var FieldlistMap map[string]interface{}
	err := json.Unmarshal(b, &FieldlistMap)
	if err != nil {
		return err
	}
	
	if CustomLabels, ok := FieldlistMap["customLabels"].(bool); ok {
		o.CustomLabels = &CustomLabels
	}
	
	if InstructionText, ok := FieldlistMap["instructionText"].(string); ok {
		o.InstructionText = &InstructionText
	}
	
	if Key, ok := FieldlistMap["key"].(string); ok {
		o.Key = &Key
	}
	
	if LabelKeys, ok := FieldlistMap["labelKeys"].([]interface{}); ok {
		LabelKeysString, _ := json.Marshal(LabelKeys)
		json.Unmarshal(LabelKeysString, &o.LabelKeys)
	}
	
	if Params, ok := FieldlistMap["params"].(map[string]interface{}); ok {
		ParamsString, _ := json.Marshal(Params)
		json.Unmarshal(ParamsString, &o.Params)
	}
	
	if Repeatable, ok := FieldlistMap["repeatable"].(bool); ok {
		o.Repeatable = &Repeatable
	}
	
	if State, ok := FieldlistMap["state"].(string); ok {
		o.State = &State
	}
	
	if VarType, ok := FieldlistMap["type"].(string); ok {
		o.VarType = &VarType
	}
	
	if Required, ok := FieldlistMap["required"].(bool); ok {
		o.Required = &Required
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Fieldlist) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
