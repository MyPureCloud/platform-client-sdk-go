package platformclientv2
import (
	"encoding/json"
)

// Userapp - Details for a UserApp
type Userapp struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name - The name of the userApp, used to distinguish this userApp from others of the same type.
	Name *string `json:"name,omitempty"`


	// IntegrationType - Integration Type for the userApp
	IntegrationType *Integrationtype `json:"integrationType,omitempty"`


	// Config
	Config *Userappconfigurationinfo `json:"config,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Userapp) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
