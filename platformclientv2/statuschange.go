package platformclientv2
import (
	"time"
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

// String returns a JSON representation of the model
func (o *Statuschange) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
