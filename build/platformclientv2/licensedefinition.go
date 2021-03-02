package platformclientv2
import (
	"encoding/json"
)

// Licensedefinition
type Licensedefinition struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Description
	Description *string `json:"description,omitempty"`


	// Permissions
	Permissions *Permissions `json:"permissions,omitempty"`


	// Prerequisites
	Prerequisites *[]Addressablelicensedefinition `json:"prerequisites,omitempty"`


	// Comprises
	Comprises *[]Licensedefinition `json:"comprises,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Licensedefinition) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
