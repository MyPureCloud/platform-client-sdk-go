package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Bulkupdateshifttradestateresultitem
type Bulkupdateshifttradestateresultitem struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// State - The state of the shift trade after the update request is processed
	State *string `json:"state,omitempty"`


	// ReviewedBy - The user who reviewed the request, if applicable
	ReviewedBy *Userreference `json:"reviewedBy,omitempty"`


	// ReviewedDate - The date the request was reviewed, if applicable. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ReviewedDate *time.Time `json:"reviewedDate,omitempty"`


	// FailureReason - The reason the update failed, if applicable
	FailureReason *string `json:"failureReason,omitempty"`


	// Metadata - Version metadata for the shift trade
	Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`

}

func (u *Bulkupdateshifttradestateresultitem) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Bulkupdateshifttradestateresultitem

	
	ReviewedDate := new(string)
	if u.ReviewedDate != nil {
		
		*ReviewedDate = timeutil.Strftime(u.ReviewedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ReviewedDate = nil
	}
	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		State *string `json:"state,omitempty"`
		
		ReviewedBy *Userreference `json:"reviewedBy,omitempty"`
		
		ReviewedDate *string `json:"reviewedDate,omitempty"`
		
		FailureReason *string `json:"failureReason,omitempty"`
		
		Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`
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
func (o *Bulkupdateshifttradestateresultitem) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
