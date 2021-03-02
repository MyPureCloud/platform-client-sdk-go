package platformclientv2
import (
	"encoding/json"
)

// Flowoutcome
type Flowoutcome struct { 
	// Id - The flow outcome identifier
	Id *string `json:"id,omitempty"`


	// Name - The flow outcome name.
	Name *string `json:"name,omitempty"`


	// Description
	Description *string `json:"description,omitempty"`


	// CurrentOperation
	CurrentOperation *Operation `json:"currentOperation,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Flowoutcome) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
