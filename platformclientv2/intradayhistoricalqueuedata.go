package platformclientv2
import (
	"encoding/json"
)

// Intradayhistoricalqueuedata
type Intradayhistoricalqueuedata struct { 
	// Offered - The number of interactions routed into the queue for the given media type(s) for an agent to answer
	Offered *int32 `json:"offered,omitempty"`


	// Completed - The number of interactions completed
	Completed *int32 `json:"completed,omitempty"`


	// Answered - The number of interactions answered by an agent in a given period
	Answered *int32 `json:"answered,omitempty"`


	// Abandoned - The number of customers who disconnect before connecting with an agent
	Abandoned *int32 `json:"abandoned,omitempty"`


	// AverageTalkTimeSeconds - The average time in seconds an agent spends interacting with a customer per talk segment for a defined period of time
	AverageTalkTimeSeconds *float64 `json:"averageTalkTimeSeconds,omitempty"`


	// AverageAfterCallWorkSeconds - The average time in seconds spent in after-call work. After-call work is the work that an agent performs immediately following an interaction
	AverageAfterCallWorkSeconds *float64 `json:"averageAfterCallWorkSeconds,omitempty"`


	// ServiceLevelPercent - Percent of interactions answered in X seconds, where X is the service level objective configured in the service goal group matching this intraday group
	ServiceLevelPercent *float64 `json:"serviceLevelPercent,omitempty"`


	// AverageSpeedOfAnswerSeconds - The average time in seconds it takes to answer an interaction once the interaction becomes available to be routed
	AverageSpeedOfAnswerSeconds *float64 `json:"averageSpeedOfAnswerSeconds,omitempty"`

}

// String returns a JSON representation of the model
func (o *Intradayhistoricalqueuedata) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
