package platformclientv2
import (
	"github.com/leekchan/timeutil"
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


	// Next
	Next *string `json:"next,omitempty"`


	// SelfUri
	SelfUri *string `json:"selfUri,omitempty"`

}

func (u *Webchatmessageentitylist) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Webchatmessageentitylist

	

	return json.Marshal(&struct { 
		PageSize *int `json:"pageSize,omitempty"`
		
		Entities *[]Webchatmessage `json:"entities,omitempty"`
		
		PreviousPage *string `json:"previousPage,omitempty"`
		
		Next *string `json:"next,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		PageSize: u.PageSize,
		
		Entities: u.Entities,
		
		PreviousPage: u.PreviousPage,
		
		Next: u.Next,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Webchatmessageentitylist) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
