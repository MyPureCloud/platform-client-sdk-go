package platformclientv2
import (
	"encoding/json"
)

// Propertychange
type Propertychange struct { 
	// Property - The property that was changed
	Property *string `json:"property,omitempty"`


	// OldValues - Previous values for the property.
	OldValues *[]string `json:"oldValues,omitempty"`


	// NewValues - New values for the property.
	NewValues *[]string `json:"newValues,omitempty"`

}

// String returns a JSON representation of the model
func (o *Propertychange) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
