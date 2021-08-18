package platformclientv2
import (
	"github.com/leekchan/timeutil"
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

func (u *Buagentschedulehistoryresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Buagentschedulehistoryresponse

	

	return json.Marshal(&struct { 
		PriorPublishedSchedules *[]Buschedulereference `json:"priorPublishedSchedules,omitempty"`
		
		BasePublishedSchedule *Buagentschedulehistorychange `json:"basePublishedSchedule,omitempty"`
		
		DroppedChanges *[]Buagentschedulehistorydroppedchange `json:"droppedChanges,omitempty"`
		
		Changes *[]Buagentschedulehistorychange `json:"changes,omitempty"`
		*Alias
	}{ 
		PriorPublishedSchedules: u.PriorPublishedSchedules,
		
		BasePublishedSchedule: u.BasePublishedSchedule,
		
		DroppedChanges: u.DroppedChanges,
		
		Changes: u.Changes,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Buagentschedulehistoryresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
