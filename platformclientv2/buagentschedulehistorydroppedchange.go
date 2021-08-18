package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Buagentschedulehistorydroppedchange
type Buagentschedulehistorydroppedchange struct { 
	// Metadata - The metadata of the change, including who and when the change was made
	Metadata *Buagentschedulehistorychangemetadata `json:"metadata,omitempty"`


	// ShiftIds - The IDs of deleted shifts
	ShiftIds *[]string `json:"shiftIds,omitempty"`


	// FullDayTimeOffMarkerDates - The dates of any deleted full day time off markers
	FullDayTimeOffMarkerDates *[]time.Time `json:"fullDayTimeOffMarkerDates,omitempty"`


	// Deletes - The deleted shifts, full day time off markers, or the entire agent schedule
	Deletes *Buagentschedulehistorydeletedchange `json:"deletes,omitempty"`

}

func (u *Buagentschedulehistorydroppedchange) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Buagentschedulehistorydroppedchange

	

	return json.Marshal(&struct { 
		Metadata *Buagentschedulehistorychangemetadata `json:"metadata,omitempty"`
		
		ShiftIds *[]string `json:"shiftIds,omitempty"`
		
		FullDayTimeOffMarkerDates *[]time.Time `json:"fullDayTimeOffMarkerDates,omitempty"`
		
		Deletes *Buagentschedulehistorydeletedchange `json:"deletes,omitempty"`
		*Alias
	}{ 
		Metadata: u.Metadata,
		
		ShiftIds: u.ShiftIds,
		
		FullDayTimeOffMarkerDates: u.FullDayTimeOffMarkerDates,
		
		Deletes: u.Deletes,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Buagentschedulehistorydroppedchange) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
