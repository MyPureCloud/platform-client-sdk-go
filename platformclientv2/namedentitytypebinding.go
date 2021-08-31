package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Namedentitytypebinding
type Namedentitytypebinding struct { 
	// EntityType - The named entity type of the binding. It can be a built-in one such as builtin:number or a custom entity type such as BeverageType.
	EntityType *string `json:"entityType,omitempty"`


	// EntityName - The name that this named entity type is bound to.
	EntityName *string `json:"entityName,omitempty"`

}

func (o *Namedentitytypebinding) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Namedentitytypebinding
	
	return json.Marshal(&struct { 
		EntityType *string `json:"entityType,omitempty"`
		
		EntityName *string `json:"entityName,omitempty"`
		*Alias
	}{ 
		EntityType: o.EntityType,
		
		EntityName: o.EntityName,
		Alias:    (*Alias)(o),
	})
}

func (o *Namedentitytypebinding) UnmarshalJSON(b []byte) error {
	var NamedentitytypebindingMap map[string]interface{}
	err := json.Unmarshal(b, &NamedentitytypebindingMap)
	if err != nil {
		return err
	}
	
	if EntityType, ok := NamedentitytypebindingMap["entityType"].(string); ok {
		o.EntityType = &EntityType
	}
	
	if EntityName, ok := NamedentitytypebindingMap["entityName"].(string); ok {
		o.EntityName = &EntityName
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Namedentitytypebinding) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
