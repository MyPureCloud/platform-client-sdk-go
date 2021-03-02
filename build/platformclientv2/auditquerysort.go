package platformclientv2
import (
	"encoding/json"
)

// Auditquerysort
type Auditquerysort struct { 
	// Name - Name of the property to sort.
	Name *string `json:"name,omitempty"`


	// SortOrder - Sort Order
	SortOrder *string `json:"sortOrder,omitempty"`

}

// String returns a JSON representation of the model
func (o *Auditquerysort) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
