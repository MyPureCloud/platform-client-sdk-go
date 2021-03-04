package platformclientv2
import (
	"encoding/json"
)

// Topicsentitylisting
type Topicsentitylisting struct { 
	// Entities
	Entities *[]Listedtopic `json:"entities,omitempty"`


	// PageSize
	PageSize *int `json:"pageSize,omitempty"`


	// SelfUri
	SelfUri *string `json:"selfUri,omitempty"`


	// NextUri
	NextUri *string `json:"nextUri,omitempty"`


	// PageCount
	PageCount *int `json:"pageCount,omitempty"`

}

// String returns a JSON representation of the model
func (o *Topicsentitylisting) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
