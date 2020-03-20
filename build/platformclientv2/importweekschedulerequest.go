package platformclientv2
import (
	"encoding/json"
)

// Importweekschedulerequest - Information to create a schedule for a week in management unit using imported data
type Importweekschedulerequest struct { 
	// Description - Description for the schedule
	Description *string `json:"description,omitempty"`


	// UserSchedules - User schedules
	UserSchedules *map[string]Userschedule `json:"userSchedules,omitempty"`


	// Published - Whether the schedule is published
	Published *bool `json:"published,omitempty"`


	// ShortTermForecastId - Short term forecast that should be associated with this schedule
	ShortTermForecastId *string `json:"shortTermForecastId,omitempty"`


	// PartialUploadIds - IDs of partial uploads of user schedules to import week schedule. It is applicable only for large schedules where activity count in schedule is greater than 17500
	PartialUploadIds *[]string `json:"partialUploadIds,omitempty"`

}

// String returns a JSON representation of the model
func (o *Importweekschedulerequest) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
