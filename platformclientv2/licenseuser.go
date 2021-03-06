package platformclientv2
import (
	"encoding/json"
)

// Licenseuser
type Licenseuser struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Licenses
	Licenses *[]Licensedefinition `json:"licenses,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Licenseuser) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
