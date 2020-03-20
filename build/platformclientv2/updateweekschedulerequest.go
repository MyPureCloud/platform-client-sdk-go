package platformclientv2
import (
	"encoding/json"
)

// Updateweekschedulerequest
type Updateweekschedulerequest struct { 
	// Description - Description of the week schedule
	Description *string `json:"description,omitempty"`


	// Published - Whether the week schedule is published
	Published *bool `json:"published,omitempty"`


	// UserSchedules - User schedules in the week
	UserSchedules *map[string]Userschedule `json:"userSchedules,omitempty"`


	// PartialUploadIds - IDs of partial uploads to include in this imported schedule. It is applicable only for large schedules where activity count in schedule is greater than 17500
	PartialUploadIds *[]string `json:"partialUploadIds,omitempty"`


	// Metadata - Version metadata for this work plan
	Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`


	// AgentSchedulesVersion - Version of agent schedules in the week schedule
	AgentSchedulesVersion *int32 `json:"agentSchedulesVersion,omitempty"`


	// ShortTermForecast - Reference to optionally point the schedule at a new short term forecast
	ShortTermForecast *Shorttermforecastreference `json:"shortTermForecast,omitempty"`


	// HeadcountForecast - The headcount forecast associated with the schedule.  If not null, existing values will be irrecoverably replaced
	HeadcountForecast *Headcountforecast `json:"headcountForecast,omitempty"`


	// AgentUpdateFilter - For a published schedule, this determines whether a notification will be shown to agents in the default PureCloud user interface.  The CPC notification will always be sent and the value specified here affects what data is returned in the 'updates' property.  In the default PureCloud UI, \"None\" means that agents will not be notified, \"ShiftTimesOnly\" means agents will only be notified for changes to shift start and end times,  and \"All\" means that agents will be notified for any change to a shift or activity (except for full day off activities).  When building a custom client, use this property to specify the level of detail you need. Defaults to \"ShiftTimesOnly\".
	AgentUpdateFilter *string `json:"agentUpdateFilter,omitempty"`

}

// String returns a JSON representation of the model
func (o *Updateweekschedulerequest) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
