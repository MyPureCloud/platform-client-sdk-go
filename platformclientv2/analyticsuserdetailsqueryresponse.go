package platformclientv2
import (
	"github.com/leekchan/timeutil"
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

func (u *Analyticsuserdetailsqueryresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Analyticsuserdetailsqueryresponse

	

	return json.Marshal(&struct { 
		UserDetails *[]Analyticsuserdetail `json:"userDetails,omitempty"`
		
		Aggregations *[]Aggregationresult `json:"aggregations,omitempty"`
		
		TotalHits *int `json:"totalHits,omitempty"`
		*Alias
	}{ 
		UserDetails: u.UserDetails,
		
		Aggregations: u.Aggregations,
		
		TotalHits: u.TotalHits,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Analyticsuserdetailsqueryresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
