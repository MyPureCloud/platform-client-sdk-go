package platformclientv2
import (
	"encoding/json"
)

// Integrationactionfields
type Integrationactionfields struct { 
	// IntegrationAction - Reference to the Integration Action to be used when integrationAction type is qualified
	IntegrationAction *Integrationaction `json:"integrationAction,omitempty"`


	// RequestMappings - Collection of Request Mappings to use
	RequestMappings *[]Requestmapping `json:"requestMappings,omitempty"`

}

// String returns a JSON representation of the model
func (o *Integrationactionfields) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
