package platformclientv2
import (
	"encoding/json"
)

// Unpublishedprogramsentitylisting
type Unpublishedprogramsentitylisting struct { 
	// Entities
	Entities *[]Program `json:"entities,omitempty"`


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
func (o *Unpublishedprogramsentitylisting) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
