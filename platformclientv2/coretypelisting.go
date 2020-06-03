package platformclientv2
import (
	"encoding/json"
)

// Coretypelisting
type Coretypelisting struct { 
	// Total
	Total *int64 `json:"total,omitempty"`


	// Entities
	Entities *[]Coretype `json:"entities,omitempty"`


	// SelfUri
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Coretypelisting) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
