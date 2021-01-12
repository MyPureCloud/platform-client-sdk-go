package platformclientv2
import (
	"encoding/json"
)

// Timezoneentitylisting
type Timezoneentitylisting struct { 
	// Entities
	Entities *[]Regiontimezone `json:"entities,omitempty"`


	// PageSize
	PageSize *int32 `json:"pageSize,omitempty"`


	// PageNumber
	PageNumber *int32 `json:"pageNumber,omitempty"`


	// Total
	Total *int64 `json:"total,omitempty"`


	// FirstUri
	FirstUri *string `json:"firstUri,omitempty"`


	// LastUri
	LastUri *string `json:"lastUri,omitempty"`


	// PreviousUri
	PreviousUri *string `json:"previousUri,omitempty"`


	// SelfUri
	SelfUri *string `json:"selfUri,omitempty"`


	// NextUri
	NextUri *string `json:"nextUri,omitempty"`


	// PageCount
	PageCount *int32 `json:"pageCount,omitempty"`

}

// String returns a JSON representation of the model
func (o *Timezoneentitylisting) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
