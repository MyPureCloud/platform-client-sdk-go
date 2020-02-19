package platformclientv2
import (
	"encoding/json"
)

// Intradayperformancepredictionqueuedata
type Intradayperformancepredictionqueuedata struct { 
	// ServiceLevelPercent - Predicted percent of interactions answered in X seconds, where X is the service level objective configured in the service goal group matching this intraday group
	ServiceLevelPercent *float64 `json:"serviceLevelPercent,omitempty"`


	// AverageSpeedOfAnswerSeconds - Predicted average time in seconds it takes to answer an interaction once the interaction becomes available to be routed
	AverageSpeedOfAnswerSeconds *float64 `json:"averageSpeedOfAnswerSeconds,omitempty"`


	// NumberOfInteractions - Predicted number of interactions
	NumberOfInteractions *float64 `json:"numberOfInteractions,omitempty"`

}

// String returns a JSON representation of the model
func (o *Intradayperformancepredictionqueuedata) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
