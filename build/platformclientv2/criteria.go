package platformclientv2
import (
	"encoding/json"
)

// Criteria
type Criteria struct { 
	// Key - The criteria key.
	Key *string `json:"key,omitempty"`


	// Values - The criteria values.
	Values *[]string `json:"values,omitempty"`


	// ShouldIgnoreCase - Should criteria be case insensitive.
	ShouldIgnoreCase *bool `json:"shouldIgnoreCase,omitempty"`


	// Operator - The comparison operator.
	Operator *string `json:"operator,omitempty"`

}

// String returns a JSON representation of the model
func (o *Criteria) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
