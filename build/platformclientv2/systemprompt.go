package platformclientv2
import (
	"encoding/json"
)

// Systemprompt
type Systemprompt struct { 
	// Id - The system prompt identifier
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Description
	Description *string `json:"description,omitempty"`


	// Resources
	Resources *[]Systempromptasset `json:"resources,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Systemprompt) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
