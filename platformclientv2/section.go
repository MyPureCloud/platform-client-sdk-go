package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
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

func (u *Section) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Section

	

	return json.Marshal(&struct { 
		FieldList *[]Fieldlist `json:"fieldList,omitempty"`
		
		InstructionText *string `json:"instructionText,omitempty"`
		
		Key *string `json:"key,omitempty"`
		
		State *string `json:"state,omitempty"`
		*Alias
	}{ 
		FieldList: u.FieldList,
		
		InstructionText: u.InstructionText,
		
		Key: u.Key,
		
		State: u.State,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Section) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
