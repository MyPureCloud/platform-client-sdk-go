package platformclientv2
import (
	"encoding/json"
)

// Callhistoryconversationentitylisting
type Callhistoryconversationentitylisting struct { 
	// Entities
	Entities *[]Callhistoryconversation `json:"entities,omitempty"`


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


	// PreviousUri
	PreviousUri *string `json:"previousUri,omitempty"`


	// NextUri
	NextUri *string `json:"nextUri,omitempty"`


	// LastUri
	LastUri *string `json:"lastUri,omitempty"`


	// PageCount
	PageCount *int32 `json:"pageCount,omitempty"`

}

// String returns a JSON representation of the model
func (o *Callhistoryconversationentitylisting) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
