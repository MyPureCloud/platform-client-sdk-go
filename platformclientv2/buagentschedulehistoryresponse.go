package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Buagentschedulehistoryresponse
type Buagentschedulehistoryresponse struct { 
	// PriorPublishedSchedules - The list of previously published schedules
	PriorPublishedSchedules *[]Buschedulereference `json:"priorPublishedSchedules,omitempty"`


	// BasePublishedSchedule - The originally published agent schedules
	BasePublishedSchedule *Buagentschedulehistorychange `json:"basePublishedSchedule,omitempty"`


	// DroppedChanges - The changes dropped from the schedule history. This will happen if the schedule history is too large
	DroppedChanges *[]Buagentschedulehistorydroppedchange `json:"droppedChanges,omitempty"`


	// Changes - The list of changes for the schedule history
	Changes *[]Buagentschedulehistorychange `json:"changes,omitempty"`

}

// String returns a JSON representation of the model
func (o *Buagentschedulehistoryresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
