package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
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

func (u *Wfmintradaydataupdatetopicintradayperformancepredictionqueuedata) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmintradaydataupdatetopicintradayperformancepredictionqueuedata

	

	return json.Marshal(&struct { 
		ServiceLevelPercent *float32 `json:"serviceLevelPercent,omitempty"`
		
		AverageSpeedOfAnswerSeconds *float32 `json:"averageSpeedOfAnswerSeconds,omitempty"`
		
		NumberOfInteractions *float32 `json:"numberOfInteractions,omitempty"`
		*Alias
	}{ 
		ServiceLevelPercent: u.ServiceLevelPercent,
		
		AverageSpeedOfAnswerSeconds: u.AverageSpeedOfAnswerSeconds,
		
		NumberOfInteractions: u.NumberOfInteractions,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Wfmintradaydataupdatetopicintradayperformancepredictionqueuedata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
