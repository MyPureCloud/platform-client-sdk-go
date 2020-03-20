package platformclientv2
import (
	"encoding/json"
)

// Analyticsuserdetailsasyncqueryresponse
type Analyticsuserdetailsasyncqueryresponse struct { 
	// UserDetails
	UserDetails *[]Analyticsuserdetail `json:"userDetails,omitempty"`


	// Cursor - Optional cursor to indicate where to resume the results
	Cursor *string `json:"cursor,omitempty"`

}

// String returns a JSON representation of the model
func (o *Analyticsuserdetailsasyncqueryresponse) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
