package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Statuschange
type Statuschange struct { 
	// DateStatusChanged - The date of this status change. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateStatusChanged *time.Time `json:"dateStatusChanged,omitempty"`


	// Status - The status the change request transitioned to
	Status *string `json:"status,omitempty"`


	// PreviousStatus - The status the change request transitioned from
	PreviousStatus *string `json:"previousStatus,omitempty"`


	// Message - A short message describing the status change
	Message *string `json:"message,omitempty"`


	// ChangedBy - If applicable, the user who updated the change request to this status
	ChangedBy *string `json:"changedBy,omitempty"`


	// RejectReason - The reason for rejecting the limit override request
	RejectReason *string `json:"rejectReason,omitempty"`

}

func (u *Statuschange) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Statuschange

	
	DateStatusChanged := new(string)
	if u.DateStatusChanged != nil {
		
		*DateStatusChanged = timeutil.Strftime(u.DateStatusChanged, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateStatusChanged = nil
	}
	

	return json.Marshal(&struct { 
		DateStatusChanged *string `json:"dateStatusChanged,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		PreviousStatus *string `json:"previousStatus,omitempty"`
		
		Message *string `json:"message,omitempty"`
		
		ChangedBy *string `json:"changedBy,omitempty"`
		
		RejectReason *string `json:"rejectReason,omitempty"`
		*Alias
	}{ 
		DateStatusChanged: DateStatusChanged,
		
		Status: u.Status,
		
		PreviousStatus: u.PreviousStatus,
		
		Message: u.Message,
		
		ChangedBy: u.ChangedBy,
		
		RejectReason: u.RejectReason,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Statuschange) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
