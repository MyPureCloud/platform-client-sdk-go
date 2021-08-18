package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Analyticsuserdetailsasyncqueryresponse
type Analyticsuserdetailsasyncqueryresponse struct { 
	// UserDetails
	UserDetails *[]Analyticsuserdetail `json:"userDetails,omitempty"`


	// Cursor - Optional cursor to indicate where to resume the results
	Cursor *string `json:"cursor,omitempty"`


	// DataAvailabilityDate - Data available up to at least this datetime. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DataAvailabilityDate *time.Time `json:"dataAvailabilityDate,omitempty"`

}

func (u *Analyticsuserdetailsasyncqueryresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Analyticsuserdetailsasyncqueryresponse

	
	DataAvailabilityDate := new(string)
	if u.DataAvailabilityDate != nil {
		
		*DataAvailabilityDate = timeutil.Strftime(u.DataAvailabilityDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DataAvailabilityDate = nil
	}
	

	return json.Marshal(&struct { 
		UserDetails *[]Analyticsuserdetail `json:"userDetails,omitempty"`
		
		Cursor *string `json:"cursor,omitempty"`
		
		DataAvailabilityDate *string `json:"dataAvailabilityDate,omitempty"`
		*Alias
	}{ 
		UserDetails: u.UserDetails,
		
		Cursor: u.Cursor,
		
		DataAvailabilityDate: DataAvailabilityDate,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Analyticsuserdetailsasyncqueryresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
