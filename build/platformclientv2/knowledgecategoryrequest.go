package platformclientv2
import (
	"encoding/json"
)

// Knowledgecategoryrequest
type Knowledgecategoryrequest struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name - Category name
	Name *string `json:"name,omitempty"`


	// Description - Category description
	Description *string `json:"description,omitempty"`


	// Parent - Category parent
	Parent *Documentcategoryinput `json:"parent,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Knowledgecategoryrequest) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
