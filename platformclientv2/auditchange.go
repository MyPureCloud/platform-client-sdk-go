package platformclientv2
import (
	"encoding/json"
)

// Auditchange
type Auditchange struct { 
	// Property
	Property *string `json:"property,omitempty"`


	// Entity
	Entity *Auditentityreference `json:"entity,omitempty"`


	// OldValues
	OldValues *[]string `json:"oldValues,omitempty"`


	// NewValues
	NewValues *[]string `json:"newValues,omitempty"`

}

// String returns a JSON representation of the model
func (o *Auditchange) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
