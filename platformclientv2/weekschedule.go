package platformclientv2
import (
	"encoding/json"
)

// Weekschedule - Week schedule information
type Weekschedule struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`


	// WeekDate - First day of this week schedule in yyyy-MM-dd format
	WeekDate *string `json:"weekDate,omitempty"`


	// Description - Description of the week schedule
	Description *string `json:"description,omitempty"`


	// Published - Whether the week schedule is published
	Published *bool `json:"published,omitempty"`


	// GenerationResults - Summary of the results from the schedule run
	GenerationResults *Weekschedulegenerationresult `json:"generationResults,omitempty"`


	// ShortTermForecast - Short term forecast associated with this schedule
	ShortTermForecast *Shorttermforecastreference `json:"shortTermForecast,omitempty"`


	// Metadata - Version metadata for this work plan
	Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`


	// UserSchedules - User schedules in the week
	UserSchedules *map[string]Userschedule `json:"userSchedules,omitempty"`


	// HeadcountForecast - Headcount information for the week schedule
	HeadcountForecast *Headcountforecast `json:"headcountForecast,omitempty"`


	// AgentSchedulesVersion - Version of agent schedules in the week schedule
	AgentSchedulesVersion *int `json:"agentSchedulesVersion,omitempty"`

}

// String returns a JSON representation of the model
func (o *Weekschedule) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
