package platformclientv2
import (
	"encoding/json"
)

// Section
type Section struct { 
	// FieldList
	FieldList *[]Fieldlist `json:"fieldList,omitempty"`


	// InstructionText
	InstructionText *string `json:"instructionText,omitempty"`


	// Key
	Key *string `json:"key,omitempty"`


	// State
	State *string `json:"state,omitempty"`

}

// String returns a JSON representation of the model
func (o *Section) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
