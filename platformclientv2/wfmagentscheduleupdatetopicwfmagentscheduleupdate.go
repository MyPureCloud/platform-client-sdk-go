package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmagentscheduleupdatetopicwfmagentscheduleupdate
type Wfmagentscheduleupdatetopicwfmagentscheduleupdate struct { 
	// UpdateType
	UpdateType *string `json:"updateType,omitempty"`


	// ShiftStartDates
	ShiftStartDates *[]time.Time `json:"shiftStartDates,omitempty"`

}

func (u *Wfmagentscheduleupdatetopicwfmagentscheduleupdate) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmagentscheduleupdatetopicwfmagentscheduleupdate

	

	return json.Marshal(&struct { 
		UpdateType *string `json:"updateType,omitempty"`
		
		ShiftStartDates *[]time.Time `json:"shiftStartDates,omitempty"`
		*Alias
	}{ 
		UpdateType: u.UpdateType,
		
		ShiftStartDates: u.ShiftStartDates,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Wfmagentscheduleupdatetopicwfmagentscheduleupdate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
