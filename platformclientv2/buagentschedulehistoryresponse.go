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

func (o *Buagentschedulehistoryresponse) MarshalJSON() ([]byte, error) {
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
		PriorPublishedSchedules: o.PriorPublishedSchedules,
		
		BasePublishedSchedule: o.BasePublishedSchedule,
		
		DroppedChanges: o.DroppedChanges,
		
		Changes: o.Changes,
		Alias:    (*Alias)(o),
	})
}

func (o *Buagentschedulehistoryresponse) UnmarshalJSON(b []byte) error {
	var BuagentschedulehistoryresponseMap map[string]interface{}
	err := json.Unmarshal(b, &BuagentschedulehistoryresponseMap)
	if err != nil {
		return err
	}
	
	if PriorPublishedSchedules, ok := BuagentschedulehistoryresponseMap["priorPublishedSchedules"].([]interface{}); ok {
		PriorPublishedSchedulesString, _ := json.Marshal(PriorPublishedSchedules)
		json.Unmarshal(PriorPublishedSchedulesString, &o.PriorPublishedSchedules)
	}
	
	if BasePublishedSchedule, ok := BuagentschedulehistoryresponseMap["basePublishedSchedule"].(map[string]interface{}); ok {
		BasePublishedScheduleString, _ := json.Marshal(BasePublishedSchedule)
		json.Unmarshal(BasePublishedScheduleString, &o.BasePublishedSchedule)
	}
	
	if DroppedChanges, ok := BuagentschedulehistoryresponseMap["droppedChanges"].([]interface{}); ok {
		DroppedChangesString, _ := json.Marshal(DroppedChanges)
		json.Unmarshal(DroppedChangesString, &o.DroppedChanges)
	}
	
	if Changes, ok := BuagentschedulehistoryresponseMap["changes"].([]interface{}); ok {
		ChangesString, _ := json.Marshal(Changes)
		json.Unmarshal(ChangesString, &o.Changes)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Buagentschedulehistoryresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
