package platformclientv2
import (
	"encoding/json"
)

// Authzsubject
type Authzsubject struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Grants
	Grants *[]Authzgrant `json:"grants,omitempty"`


	// Version
	Version *int32 `json:"version,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Authzsubject) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
