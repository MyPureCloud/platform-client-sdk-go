package platformclientv2
import (
	"encoding/json"
)

// Webchatdeploymententitylisting
type Webchatdeploymententitylisting struct { 
	// Total
	Total *int64 `json:"total,omitempty"`


	// Entities
	Entities *[]Webchatdeployment `json:"entities,omitempty"`


	// SelfUri
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Webchatdeploymententitylisting) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
