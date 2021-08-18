package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
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

func (u *Wfmbulkshifttradestateupdatenotificationtopicbulkshifttradestateupdateresult) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmbulkshifttradestateupdatenotificationtopicbulkshifttradestateupdateresult

	
	ReviewedDate := new(string)
	if u.ReviewedDate != nil {
		
		*ReviewedDate = timeutil.Strftime(u.ReviewedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ReviewedDate = nil
	}
	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		State *string `json:"state,omitempty"`
		
		ReviewedBy *Wfmbulkshifttradestateupdatenotificationtopicuserreference `json:"reviewedBy,omitempty"`
		
		ReviewedDate *string `json:"reviewedDate,omitempty"`
		
		FailureReason *string `json:"failureReason,omitempty"`
		
		Metadata *Wfmbulkshifttradestateupdatenotificationtopicwfmversionedentitymetadata `json:"metadata,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		State: u.State,
		
		ReviewedBy: u.ReviewedBy,
		
		ReviewedDate: ReviewedDate,
		
		FailureReason: u.FailureReason,
		
		Metadata: u.Metadata,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Wfmbulkshifttradestateupdatenotificationtopicbulkshifttradestateupdateresult) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
