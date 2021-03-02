package platformclientv2
import (
	"encoding/json"
)

// Createintegrationrequest - Details for an Integration
type Createintegrationrequest struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name - The name of the integration, used to distinguish this integration from others of the same type.
	Name *string `json:"name,omitempty"`


	// IntegrationType - Type of the integration to create.
	IntegrationType *Integrationtype `json:"integrationType,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Createintegrationrequest) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
