package platformclientv2
import (
	"encoding/json"
)

// Wfmintradaydataupdatetopicintradayperformancepredictionqueuedata
type Wfmintradaydataupdatetopicintradayperformancepredictionqueuedata struct { 
	// ServiceLevelPercent
	ServiceLevelPercent *float32 `json:"serviceLevelPercent,omitempty"`


	// AverageSpeedOfAnswerSeconds
	AverageSpeedOfAnswerSeconds *float32 `json:"averageSpeedOfAnswerSeconds,omitempty"`


	// NumberOfInteractions
	NumberOfInteractions *float32 `json:"numberOfInteractions,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wfmintradaydataupdatetopicintradayperformancepredictionqueuedata) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
