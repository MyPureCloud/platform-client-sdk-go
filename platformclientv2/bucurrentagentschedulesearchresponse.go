package platformclientv2
import (
	"time"
	"encoding/json"
)

// Bucurrentagentschedulesearchresponse
type Bucurrentagentschedulesearchresponse struct { 
	// AgentSchedules - The requested agent schedules
	AgentSchedules *[]Buagentschedulesearchresponse `json:"agentSchedules,omitempty"`


	// BusinessUnitTimeZone - The time zone configured for the business unit to which this schedule applies
	BusinessUnitTimeZone *string `json:"businessUnitTimeZone,omitempty"`


	// PublishedSchedules - References to all published week schedules overlapping the start/end date query parameters
	PublishedSchedules *[]Buagentschedulepublishedschedulereference `json:"publishedSchedules,omitempty"`


	// StartDate - The start date of the schedules. Only populated on notifications. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	StartDate *time.Time `json:"startDate,omitempty"`


	// EndDate - The end date of the schedules. Only populated on notifications. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	EndDate *time.Time `json:"endDate,omitempty"`


	// Updates - The list of updates for the schedule. Only used in notifications
	Updates *[]Buagentscheduleupdate `json:"updates,omitempty"`

}

// String returns a JSON representation of the model
func (o *Bucurrentagentschedulesearchresponse) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
