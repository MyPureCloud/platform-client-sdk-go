package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
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

func (u *Wfmbuintradaydataupdatetopicintradayperformancepredictiondata) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmbuintradaydataupdatetopicintradayperformancepredictiondata

	

	return json.Marshal(&struct { 
		ServiceLevelPercent *float32 `json:"serviceLevelPercent,omitempty"`
		
		AverageSpeedOfAnswerSeconds *float32 `json:"averageSpeedOfAnswerSeconds,omitempty"`
		
		OccupancyPercent *float32 `json:"occupancyPercent,omitempty"`
		*Alias
	}{ 
		ServiceLevelPercent: u.ServiceLevelPercent,
		
		AverageSpeedOfAnswerSeconds: u.AverageSpeedOfAnswerSeconds,
		
		OccupancyPercent: u.OccupancyPercent,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Wfmbuintradaydataupdatetopicintradayperformancepredictiondata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
