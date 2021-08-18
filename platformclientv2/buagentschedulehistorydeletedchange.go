package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Buagentschedulehistorydeletedchange
type Buagentschedulehistorydeletedchange struct { 
	// ShiftIds - The IDs of deleted shifts
	ShiftIds *[]string `json:"shiftIds,omitempty"`


	// FullDayTimeOffMarkerDates - The dates of any deleted full day time off markers
	FullDayTimeOffMarkerDates *[]time.Time `json:"fullDayTimeOffMarkerDates,omitempty"`


	// AgentSchedule - Whether the entire agent schedule was deleted
	AgentSchedule *bool `json:"agentSchedule,omitempty"`

}

func (u *Buagentschedulehistorydeletedchange) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Buagentschedulehistorydeletedchange

	

	return json.Marshal(&struct { 
		ShiftIds *[]string `json:"shiftIds,omitempty"`
		
		FullDayTimeOffMarkerDates *[]time.Time `json:"fullDayTimeOffMarkerDates,omitempty"`
		
		AgentSchedule *bool `json:"agentSchedule,omitempty"`
		*Alias
	}{ 
		ShiftIds: u.ShiftIds,
		
		FullDayTimeOffMarkerDates: u.FullDayTimeOffMarkerDates,
		
		AgentSchedule: u.AgentSchedule,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Buagentschedulehistorydeletedchange) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
