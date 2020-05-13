package platformclientv2
import (
	"encoding/json"
)

// Documentlisting
type Documentlisting struct { 
	// Entities
	Entities *[]Knowledgedocument `json:"entities,omitempty"`


	// NextUri
	NextUri *string `json:"nextUri,omitempty"`


	// SelfUri
	SelfUri *string `json:"selfUri,omitempty"`


	// PreviousUri
	PreviousUri *string `json:"previousUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Documentlisting) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
