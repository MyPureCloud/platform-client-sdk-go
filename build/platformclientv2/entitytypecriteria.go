package platformclientv2
import (
	"encoding/json"
)

// Entitytypecriteria
type Entitytypecriteria struct { 
	// Key - The criteria key.
	Key *string `json:"key,omitempty"`


	// Values - The criteria values.
	Values *[]string `json:"values,omitempty"`


	// ShouldIgnoreCase - Should criteria be case insensitive.
	ShouldIgnoreCase *bool `json:"shouldIgnoreCase,omitempty"`


	// Operator - The comparison operator.
	Operator *string `json:"operator,omitempty"`


	// EntityType - The entity to match the pattern against.
	EntityType *string `json:"entityType,omitempty"`

}

// String returns a JSON representation of the model
func (o *Entitytypecriteria) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
