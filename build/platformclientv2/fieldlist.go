package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Fieldlist) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
