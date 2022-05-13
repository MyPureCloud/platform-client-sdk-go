package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Namedentitytypedefinition
type Namedentitytypedefinition struct { 
	// Name - The name of the entity type.
	Name *string `json:"name,omitempty"`


	// Description - Description of the of the named entity type.
	Description *string `json:"description,omitempty"`


	// Mechanism - The mechanism enabling detection of the named entity type.
	Mechanism *Namedentitytypemechanism `json:"mechanism,omitempty"`

}

func (o *Namedentitytypedefinition) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Namedentitytypedefinition
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		Mechanism *Namedentitytypemechanism `json:"mechanism,omitempty"`
		*Alias
	}{ 
		Name: o.Name,
		
		Description: o.Description,
		
		Mechanism: o.Mechanism,
		Alias:    (*Alias)(o),
	})
}

func (o *Namedentitytypedefinition) UnmarshalJSON(b []byte) error {
	var NamedentitytypedefinitionMap map[string]interface{}
	err := json.Unmarshal(b, &NamedentitytypedefinitionMap)
	if err != nil {
		return err
	}
	
	if Name, ok := NamedentitytypedefinitionMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Description, ok := NamedentitytypedefinitionMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if Mechanism, ok := NamedentitytypedefinitionMap["mechanism"].(map[string]interface{}); ok {
		MechanismString, _ := json.Marshal(Mechanism)
		json.Unmarshal(MechanismString, &o.Mechanism)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Namedentitytypedefinition) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
