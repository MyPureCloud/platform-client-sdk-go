package platformclientv2
import (
	"time"
	"encoding/json"
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

// String returns a JSON representation of the model
func (o *Buagentschedulehistorydeletedchange) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
