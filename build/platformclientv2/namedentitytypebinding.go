package platformclientv2
import (
	"encoding/json"
)

// Namedentitytypebinding
type Namedentitytypebinding struct { 
	// EntityType - The named entity type of the binding. It can be a built-in one such as builtin:number or a custom entity type such as BeverageType.
	EntityType *string `json:"entityType,omitempty"`


	// EntityName - The name that this named entity type is bound to.
	EntityName *string `json:"entityName,omitempty"`

}

// String returns a JSON representation of the model
func (o *Namedentitytypebinding) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
