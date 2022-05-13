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

func (o *Wfmintradaydataupdatetopicintradayperformancepredictionqueuedata) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmintradaydataupdatetopicintradayperformancepredictionqueuedata
	
	return json.Marshal(&struct { 
		ServiceLevelPercent *float32 `json:"serviceLevelPercent,omitempty"`
		
		AverageSpeedOfAnswerSeconds *float32 `json:"averageSpeedOfAnswerSeconds,omitempty"`
		
		NumberOfInteractions *float32 `json:"numberOfInteractions,omitempty"`
		*Alias
	}{ 
		ServiceLevelPercent: o.ServiceLevelPercent,
		
		AverageSpeedOfAnswerSeconds: o.AverageSpeedOfAnswerSeconds,
		
		NumberOfInteractions: o.NumberOfInteractions,
		Alias:    (*Alias)(o),
	})
}

func (o *Wfmintradaydataupdatetopicintradayperformancepredictionqueuedata) UnmarshalJSON(b []byte) error {
	var WfmintradaydataupdatetopicintradayperformancepredictionqueuedataMap map[string]interface{}
	err := json.Unmarshal(b, &WfmintradaydataupdatetopicintradayperformancepredictionqueuedataMap)
	if err != nil {
		return err
	}
	
	if ServiceLevelPercent, ok := WfmintradaydataupdatetopicintradayperformancepredictionqueuedataMap["serviceLevelPercent"].(float64); ok {
		ServiceLevelPercentFloat32 := float32(ServiceLevelPercent)
		o.ServiceLevelPercent = &ServiceLevelPercentFloat32
	}
    
	if AverageSpeedOfAnswerSeconds, ok := WfmintradaydataupdatetopicintradayperformancepredictionqueuedataMap["averageSpeedOfAnswerSeconds"].(float64); ok {
		AverageSpeedOfAnswerSecondsFloat32 := float32(AverageSpeedOfAnswerSeconds)
		o.AverageSpeedOfAnswerSeconds = &AverageSpeedOfAnswerSecondsFloat32
	}
    
	if NumberOfInteractions, ok := WfmintradaydataupdatetopicintradayperformancepredictionqueuedataMap["numberOfInteractions"].(float64); ok {
		NumberOfInteractionsFloat32 := float32(NumberOfInteractions)
		o.NumberOfInteractions = &NumberOfInteractionsFloat32
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmintradaydataupdatetopicintradayperformancepredictionqueuedata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
