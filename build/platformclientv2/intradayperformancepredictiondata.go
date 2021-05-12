package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Intradayperformancepredictiondata
type Intradayperformancepredictiondata struct { 
	// ServiceLevelPercent - Percentage of interactions that meets service level target as defined in the matching service goal templates
	ServiceLevelPercent *float64 `json:"serviceLevelPercent,omitempty"`


	// AverageSpeedOfAnswerSeconds - Predicted average time in seconds it takes to answer an interaction once the interaction becomes available to be routed
	AverageSpeedOfAnswerSeconds *float64 `json:"averageSpeedOfAnswerSeconds,omitempty"`


	// OccupancyPercent - Percentage of on-queue time for all agents in this group that are occupied handling interactions
	OccupancyPercent *float64 `json:"occupancyPercent,omitempty"`

}

// String returns a JSON representation of the model
func (o *Intradayperformancepredictiondata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
