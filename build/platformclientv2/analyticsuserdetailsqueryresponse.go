package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Analyticsuserdetailsqueryresponse
type Analyticsuserdetailsqueryresponse struct { 
	// UserDetails
	UserDetails *[]Analyticsuserdetail `json:"userDetails,omitempty"`


	// Aggregations
	Aggregations *[]Aggregationresult `json:"aggregations,omitempty"`


	// TotalHits
	TotalHits *int `json:"totalHits,omitempty"`

}

// String returns a JSON representation of the model
func (o *Analyticsuserdetailsqueryresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
