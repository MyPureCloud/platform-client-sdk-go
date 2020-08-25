package platformclientv2
import (
	"encoding/json"
)

// Cursorrelationshiplisting
type Cursorrelationshiplisting struct { 
	// Entities
	Entities *[]Relationship `json:"entities,omitempty"`


	// NextUri
	NextUri *string `json:"nextUri,omitempty"`


	// SelfUri
	SelfUri *string `json:"selfUri,omitempty"`


	// PreviousUri
	PreviousUri *string `json:"previousUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Cursorrelationshiplisting) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
