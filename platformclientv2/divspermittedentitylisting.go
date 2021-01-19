package platformclientv2
import (
	"encoding/json"
)

// Divspermittedentitylisting
type Divspermittedentitylisting struct { 
	// Entities
	Entities *[]Authzdivision `json:"entities,omitempty"`


	// PageSize
	PageSize *int32 `json:"pageSize,omitempty"`


	// PageNumber
	PageNumber *int32 `json:"pageNumber,omitempty"`


	// Total
	Total *int64 `json:"total,omitempty"`


	// AllDivsPermitted
	AllDivsPermitted *bool `json:"allDivsPermitted,omitempty"`


	// FirstUri
	FirstUri *string `json:"firstUri,omitempty"`


	// SelfUri
	SelfUri *string `json:"selfUri,omitempty"`


	// LastUri
	LastUri *string `json:"lastUri,omitempty"`


	// NextUri
	NextUri *string `json:"nextUri,omitempty"`


	// PreviousUri
	PreviousUri *string `json:"previousUri,omitempty"`


	// PageCount
	PageCount *int32 `json:"pageCount,omitempty"`

}

// String returns a JSON representation of the model
func (o *Divspermittedentitylisting) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
