package platformclientv2
import (
	"encoding/json"
)

// Wfmbuintradaydataupdatetopicintradayperformancepredictiondata
type Wfmbuintradaydataupdatetopicintradayperformancepredictiondata struct { 
	// ServiceLevelPercent
	ServiceLevelPercent *float32 `json:"serviceLevelPercent,omitempty"`


	// AverageSpeedOfAnswerSeconds
	AverageSpeedOfAnswerSeconds *float32 `json:"averageSpeedOfAnswerSeconds,omitempty"`


	// OccupancyPercent
	OccupancyPercent *float32 `json:"occupancyPercent,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wfmbuintradaydataupdatetopicintradayperformancepredictiondata) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
