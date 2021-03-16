package platformclientv2
import (
	"encoding/json"
)

// Getmetricsresponse
type Getmetricsresponse struct { 
	// Total
	Total *int `json:"total,omitempty"`


	// Entities
	Entities *[]Metrics `json:"entities,omitempty"`


	// SelfUri
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Getmetricsresponse) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
