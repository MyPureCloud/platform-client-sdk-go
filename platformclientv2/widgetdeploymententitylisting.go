package platformclientv2
import (
	"encoding/json"
)

// Widgetdeploymententitylisting
type Widgetdeploymententitylisting struct { 
	// Total
	Total *int `json:"total,omitempty"`


	// Entities
	Entities *[]Widgetdeployment `json:"entities,omitempty"`


	// SelfUri
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Widgetdeploymententitylisting) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
