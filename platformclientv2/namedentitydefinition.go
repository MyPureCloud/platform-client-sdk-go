package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Namedentitydefinition
type Namedentitydefinition struct { 
	// Name - The name of the entity.
	Name *string `json:"name,omitempty"`


	// VarType - The name of the entity type.
	VarType *string `json:"type,omitempty"`

}

func (o *Namedentitydefinition) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Namedentitydefinition
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		*Alias
	}{ 
		Name: o.Name,
		
		VarType: o.VarType,
		Alias:    (*Alias)(o),
	})
}

func (o *Namedentitydefinition) UnmarshalJSON(b []byte) error {
	var NamedentitydefinitionMap map[string]interface{}
	err := json.Unmarshal(b, &NamedentitydefinitionMap)
	if err != nil {
		return err
	}
	
	if Name, ok := NamedentitydefinitionMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if VarType, ok := NamedentitydefinitionMap["type"].(string); ok {
		o.VarType = &VarType
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Namedentitydefinition) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
