package platformclientv2
import (
	"encoding/json"
)

// Auditqueryentity
type Auditqueryentity struct { 
	// Name - Name of the Entity
	Name *string `json:"name,omitempty"`


	// Actions - List of Actions
	Actions *[]string `json:"actions,omitempty"`

}

// String returns a JSON representation of the model
func (o *Auditqueryentity) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
