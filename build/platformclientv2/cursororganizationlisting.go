package platformclientv2
import (
	"encoding/json"
)

// Cursororganizationlisting
type Cursororganizationlisting struct { 
	// Entities
	Entities *[]Externalorganization `json:"entities,omitempty"`


	// NextUri
	NextUri *string `json:"nextUri,omitempty"`


	// SelfUri
	SelfUri *string `json:"selfUri,omitempty"`


	// PreviousUri
	PreviousUri *string `json:"previousUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Cursororganizationlisting) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
