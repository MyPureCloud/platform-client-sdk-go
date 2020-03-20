package platformclientv2
import (
	"encoding/json"
)

// Analyticsuserdetailsqueryresponse
type Analyticsuserdetailsqueryresponse struct { 
	// UserDetails
	UserDetails *[]Analyticsuserdetail `json:"userDetails,omitempty"`


	// Aggregations
	Aggregations *[]Aggregationresult `json:"aggregations,omitempty"`

}

// String returns a JSON representation of the model
func (o *Analyticsuserdetailsqueryresponse) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
