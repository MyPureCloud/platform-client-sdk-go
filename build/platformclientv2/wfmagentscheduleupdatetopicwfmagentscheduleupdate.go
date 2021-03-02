package platformclientv2
import (
	"time"
	"encoding/json"
)

// Wfmagentscheduleupdatetopicwfmagentscheduleupdate
type Wfmagentscheduleupdatetopicwfmagentscheduleupdate struct { 
	// UpdateType
	UpdateType *string `json:"updateType,omitempty"`


	// ShiftStartDates
	ShiftStartDates *[]time.Time `json:"shiftStartDates,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wfmagentscheduleupdatetopicwfmagentscheduleupdate) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
