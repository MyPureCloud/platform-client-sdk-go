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

func (o *Wfmbuintradaydataupdatetopicintradayperformancepredictiondata) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmbuintradaydataupdatetopicintradayperformancepredictiondata
	
	return json.Marshal(&struct { 
		ServiceLevelPercent *float32 `json:"serviceLevelPercent,omitempty"`
		
		AverageSpeedOfAnswerSeconds *float32 `json:"averageSpeedOfAnswerSeconds,omitempty"`
		
		OccupancyPercent *float32 `json:"occupancyPercent,omitempty"`
		*Alias
	}{ 
		ServiceLevelPercent: o.ServiceLevelPercent,
		
		AverageSpeedOfAnswerSeconds: o.AverageSpeedOfAnswerSeconds,
		
		OccupancyPercent: o.OccupancyPercent,
		Alias:    (*Alias)(o),
	})
}

func (o *Wfmbuintradaydataupdatetopicintradayperformancepredictiondata) UnmarshalJSON(b []byte) error {
	var WfmbuintradaydataupdatetopicintradayperformancepredictiondataMap map[string]interface{}
	err := json.Unmarshal(b, &WfmbuintradaydataupdatetopicintradayperformancepredictiondataMap)
	if err != nil {
		return err
	}
	
	if ServiceLevelPercent, ok := WfmbuintradaydataupdatetopicintradayperformancepredictiondataMap["serviceLevelPercent"].(float64); ok {
		ServiceLevelPercentFloat32 := float32(ServiceLevelPercent)
		o.ServiceLevelPercent = &ServiceLevelPercentFloat32
	}
	
	if AverageSpeedOfAnswerSeconds, ok := WfmbuintradaydataupdatetopicintradayperformancepredictiondataMap["averageSpeedOfAnswerSeconds"].(float64); ok {
		AverageSpeedOfAnswerSecondsFloat32 := float32(AverageSpeedOfAnswerSeconds)
		o.AverageSpeedOfAnswerSeconds = &AverageSpeedOfAnswerSecondsFloat32
	}
	
	if OccupancyPercent, ok := WfmbuintradaydataupdatetopicintradayperformancepredictiondataMap["occupancyPercent"].(float64); ok {
		OccupancyPercentFloat32 := float32(OccupancyPercent)
		o.OccupancyPercent = &OccupancyPercentFloat32
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmbuintradaydataupdatetopicintradayperformancepredictiondata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
