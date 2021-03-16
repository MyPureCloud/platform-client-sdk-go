package platformclientv2
import (
	"encoding/json"
)

// Gettemplatesresponse
type Gettemplatesresponse struct { 
	// Total
	Total *int `json:"total,omitempty"`


	// Entities
	Entities *[]Objectivetemplate `json:"entities,omitempty"`


	// SelfUri
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Gettemplatesresponse) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
