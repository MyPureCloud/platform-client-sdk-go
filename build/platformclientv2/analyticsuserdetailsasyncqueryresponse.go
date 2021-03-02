package platformclientv2
import (
	"time"
	"encoding/json"
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

// String returns a JSON representation of the model
func (o *Analyticsuserdetailsasyncqueryresponse) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
