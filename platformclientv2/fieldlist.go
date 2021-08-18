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

func (u *Fieldlist) MarshalJSON() ([]byte, error) {
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
		CustomLabels: u.CustomLabels,
		
		InstructionText: u.InstructionText,
		
		Key: u.Key,
		
		LabelKeys: u.LabelKeys,
		
		Params: u.Params,
		
		Repeatable: u.Repeatable,
		
		State: u.State,
		
		VarType: u.VarType,
		
		Required: u.Required,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Fieldlist) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
