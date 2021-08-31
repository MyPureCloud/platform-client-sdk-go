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

func (o *Section) MarshalJSON() ([]byte, error) {
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
		FieldList: o.FieldList,
		
		InstructionText: o.InstructionText,
		
		Key: o.Key,
		
		State: o.State,
		Alias:    (*Alias)(o),
	})
}

func (o *Section) UnmarshalJSON(b []byte) error {
	var SectionMap map[string]interface{}
	err := json.Unmarshal(b, &SectionMap)
	if err != nil {
		return err
	}
	
	if FieldList, ok := SectionMap["fieldList"].([]interface{}); ok {
		FieldListString, _ := json.Marshal(FieldList)
		json.Unmarshal(FieldListString, &o.FieldList)
	}
	
	if InstructionText, ok := SectionMap["instructionText"].(string); ok {
		o.InstructionText = &InstructionText
	}
	
	if Key, ok := SectionMap["key"].(string); ok {
		o.Key = &Key
	}
	
	if State, ok := SectionMap["state"].(string); ok {
		o.State = &State
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Section) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
