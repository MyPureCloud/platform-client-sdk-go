package platformclientv2
import (
	"encoding/json"
)

// Wfmintradaydataupdatetopicintradayhistoricalqueuedata
type Wfmintradaydataupdatetopicintradayhistoricalqueuedata struct { 
	// Offered
	Offered *int32 `json:"offered,omitempty"`


	// Completed
	Completed *int32 `json:"completed,omitempty"`


	// Answered
	Answered *int32 `json:"answered,omitempty"`


	// Abandoned
	Abandoned *int32 `json:"abandoned,omitempty"`


	// AverageTalkTimeSeconds
	AverageTalkTimeSeconds *float32 `json:"averageTalkTimeSeconds,omitempty"`


	// AverageAfterCallWorkSeconds
	AverageAfterCallWorkSeconds *float32 `json:"averageAfterCallWorkSeconds,omitempty"`


	// ServiceLevelPercent
	ServiceLevelPercent *float32 `json:"serviceLevelPercent,omitempty"`


	// AverageSpeedOfAnswerSeconds
	AverageSpeedOfAnswerSeconds *float32 `json:"averageSpeedOfAnswerSeconds,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wfmintradaydataupdatetopicintradayhistoricalqueuedata) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
