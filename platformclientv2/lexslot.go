package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Lexslot
type Lexslot struct { 
	// Name - The slot name
	Name *string `json:"name,omitempty"`


	// Description - The slot description
	Description *string `json:"description,omitempty"`


	// VarType - The slot type
	VarType *string `json:"type,omitempty"`


	// Priority - The priority of the slot
	Priority *int `json:"priority,omitempty"`

}

func (o *Lexslot) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Lexslot
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		Priority *int `json:"priority,omitempty"`
		*Alias
	}{ 
		Name: o.Name,
		
		Description: o.Description,
		
		VarType: o.VarType,
		
		Priority: o.Priority,
		Alias:    (*Alias)(o),
	})
}

func (o *Lexslot) UnmarshalJSON(b []byte) error {
	var LexslotMap map[string]interface{}
	err := json.Unmarshal(b, &LexslotMap)
	if err != nil {
		return err
	}
	
	if Name, ok := LexslotMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Description, ok := LexslotMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if VarType, ok := LexslotMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Priority, ok := LexslotMap["priority"].(float64); ok {
		PriorityInt := int(Priority)
		o.Priority = &PriorityInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Lexslot) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
