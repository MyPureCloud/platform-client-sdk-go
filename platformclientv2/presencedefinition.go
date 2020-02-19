package platformclientv2
import (
	"encoding/json"
)

// Presencedefinition
type Presencedefinition struct { 
	// Id - description
	Id *string `json:"id,omitempty"`


	// SystemPresence
	SystemPresence *string `json:"systemPresence,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Presencedefinition) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
