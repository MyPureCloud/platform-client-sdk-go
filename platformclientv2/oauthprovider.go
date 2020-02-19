package platformclientv2
import (
	"encoding/json"
)

// Oauthprovider
type Oauthprovider struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Disabled
	Disabled *bool `json:"disabled,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Oauthprovider) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
