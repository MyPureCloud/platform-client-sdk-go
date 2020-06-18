package platformclientv2
import (
	"encoding/json"
)

// Auditqueryservice
type Auditqueryservice struct { 
	// Name - Name of the Service
	Name *string `json:"name,omitempty"`


	// Entities - List of Entities
	Entities *[]Auditqueryentity `json:"entities,omitempty"`

}

// String returns a JSON representation of the model
func (o *Auditqueryservice) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
