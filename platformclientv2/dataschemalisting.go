package platformclientv2
import (
	"encoding/json"
)

// Dataschemalisting
type Dataschemalisting struct { 
	// Total
	Total *int64 `json:"total,omitempty"`


	// Entities
	Entities *[]Dataschema `json:"entities,omitempty"`


	// SelfUri
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Dataschemalisting) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
