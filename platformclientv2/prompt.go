package platformclientv2
import (
	"encoding/json"
)

// Prompt
type Prompt struct { 
	// Id - The prompt identifier
	Id *string `json:"id,omitempty"`


	// Name - The prompt name.
	Name *string `json:"name,omitempty"`


	// Description
	Description *string `json:"description,omitempty"`


	// Resources
	Resources *[]Promptasset `json:"resources,omitempty"`


	// CurrentOperation
	CurrentOperation *Operation `json:"currentOperation,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Prompt) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
