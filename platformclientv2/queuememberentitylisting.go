package platformclientv2
import (
	"encoding/json"
)

// Queuememberentitylisting
type Queuememberentitylisting struct { 
	// Entities
	Entities *[]Queuemember `json:"entities,omitempty"`


	// PageSize
	PageSize *int `json:"pageSize,omitempty"`


	// PageNumber
	PageNumber *int `json:"pageNumber,omitempty"`


	// Total
	Total *int `json:"total,omitempty"`


	// FirstUri
	FirstUri *string `json:"firstUri,omitempty"`


	// SelfUri
	SelfUri *string `json:"selfUri,omitempty"`


	// NextUri
	NextUri *string `json:"nextUri,omitempty"`


	// LastUri
	LastUri *string `json:"lastUri,omitempty"`


	// PreviousUri
	PreviousUri *string `json:"previousUri,omitempty"`


	// PageCount
	PageCount *int `json:"pageCount,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queuememberentitylisting) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
