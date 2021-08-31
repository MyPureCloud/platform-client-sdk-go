package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Manager - Defines a SCIM manager.
type Manager struct { 
	// Value - The ID of the manager.
	Value *string `json:"value,omitempty"`


	// Ref - The reference URI of the manager's user record.
	Ref *string `json:"$ref,omitempty"`

}

func (o *Manager) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Manager
	
	return json.Marshal(&struct { 
		Value *string `json:"value,omitempty"`
		
		Ref *string `json:"$ref,omitempty"`
		*Alias
	}{ 
		Value: o.Value,
		
		Ref: o.Ref,
		Alias:    (*Alias)(o),
	})
}

func (o *Manager) UnmarshalJSON(b []byte) error {
	var ManagerMap map[string]interface{}
	err := json.Unmarshal(b, &ManagerMap)
	if err != nil {
		return err
	}
	
	if Value, ok := ManagerMap["value"].(string); ok {
		o.Value = &Value
	}
	
	if Ref, ok := ManagerMap["$ref"].(string); ok {
		o.Ref = &Ref
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Manager) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
