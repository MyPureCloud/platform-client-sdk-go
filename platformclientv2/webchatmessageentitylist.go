package platformclientv2
import (
	"encoding/json"
)

// Webchatmessageentitylist
type Webchatmessageentitylist struct { 
	// PageSize
	PageSize *int `json:"pageSize,omitempty"`


	// Entities
	Entities *[]Webchatmessage `json:"entities,omitempty"`


	// PreviousPage
	PreviousPage *string `json:"previousPage,omitempty"`


	// Next
	Next *string `json:"next,omitempty"`


	// SelfUri
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Webchatmessageentitylist) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
