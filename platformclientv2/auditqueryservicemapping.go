package platformclientv2
import (
	"encoding/json"
)

// Auditqueryservicemapping
type Auditqueryservicemapping struct { 
	// Services - List of Services
	Services *[]Auditqueryservice `json:"services,omitempty"`

}

// String returns a JSON representation of the model
func (o *Auditqueryservicemapping) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
