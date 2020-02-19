package platformclientv2
import (
	"encoding/json"
)

// Scimserviceproviderconfigsimplefeature - Defines a request in the SCIM service provider's configuration.
type Scimserviceproviderconfigsimplefeature struct { 
	// Supported - Indicates whether configuration options are supported.
	Supported *bool `json:"supported,omitempty"`

}

// String returns a JSON representation of the model
func (o *Scimserviceproviderconfigsimplefeature) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
