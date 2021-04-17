package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Webchatmessageentitylist
type Webchatmessageentitylist struct { 
	// PageSize
	PageSize *int `json:"pageSize,omitempty"`


	// Entities
	Entities *[]Webchatmessage `json:"entities,omitempty"`


	// PreviousPage
	PreviousPage *string `json:"previousPage,omitempty"`


	// SelfUri
	SelfUri *string `json:"selfUri,omitempty"`


	// Next
	Next *string `json:"next,omitempty"`

}

// String returns a JSON representation of the model
func (o *Webchatmessageentitylist) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
