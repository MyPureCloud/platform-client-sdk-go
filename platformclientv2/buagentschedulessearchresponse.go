package platformclientv2
import (
	"encoding/json"
)

// Buagentschedulessearchresponse
type Buagentschedulessearchresponse struct { 
	// AgentSchedules - The requested agent schedules
	AgentSchedules *[]Buagentschedulesearchresponse `json:"agentSchedules,omitempty"`


	// BusinessUnitTimeZone - The time zone configured for the business unit to which this schedule applies
	BusinessUnitTimeZone *string `json:"businessUnitTimeZone,omitempty"`


	// PublishedSchedules - References to all published week schedules overlapping the start/end date query parameters
	PublishedSchedules *[]Buagentschedulepublishedschedulereference `json:"publishedSchedules,omitempty"`

}

// String returns a JSON representation of the model
func (o *Buagentschedulessearchresponse) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
