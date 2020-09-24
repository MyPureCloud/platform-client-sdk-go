package platformclientv2
import (
	"encoding/json"
)

// Flowmilestonelisting
type Flowmilestonelisting struct { 
	// Entities
	Entities *[]Flowmilestone `json:"entities,omitempty"`


	// PageSize
	PageSize *int32 `json:"pageSize,omitempty"`


	// PageNumber
	PageNumber *int32 `json:"pageNumber,omitempty"`


	// Total
	Total *int64 `json:"total,omitempty"`


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
	PageCount *int32 `json:"pageCount,omitempty"`

}

// String returns a JSON representation of the model
func (o *Flowmilestonelisting) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
