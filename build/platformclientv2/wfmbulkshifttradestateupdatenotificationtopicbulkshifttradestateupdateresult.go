package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmbulkshifttradestateupdatenotificationtopicbulkshifttradestateupdateresult
type Wfmbulkshifttradestateupdatenotificationtopicbulkshifttradestateupdateresult struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// State
	State *string `json:"state,omitempty"`


	// ReviewedBy
	ReviewedBy *Wfmbulkshifttradestateupdatenotificationtopicuserreference `json:"reviewedBy,omitempty"`


	// ReviewedDate
	ReviewedDate *time.Time `json:"reviewedDate,omitempty"`


	// FailureReason
	FailureReason *string `json:"failureReason,omitempty"`


	// Metadata
	Metadata *Wfmbulkshifttradestateupdatenotificationtopicwfmversionedentitymetadata `json:"metadata,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wfmbulkshifttradestateupdatenotificationtopicbulkshifttradestateupdateresult) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
