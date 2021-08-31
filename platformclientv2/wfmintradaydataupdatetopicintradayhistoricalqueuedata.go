package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmintradaydataupdatetopicintradayhistoricalqueuedata
type Wfmintradaydataupdatetopicintradayhistoricalqueuedata struct { 
	// Offered
	Offered *int `json:"offered,omitempty"`


	// Completed
	Completed *int `json:"completed,omitempty"`


	// Answered
	Answered *int `json:"answered,omitempty"`


	// Abandoned
	Abandoned *int `json:"abandoned,omitempty"`


	// AverageTalkTimeSeconds
	AverageTalkTimeSeconds *float32 `json:"averageTalkTimeSeconds,omitempty"`


	// AverageAfterCallWorkSeconds
	AverageAfterCallWorkSeconds *float32 `json:"averageAfterCallWorkSeconds,omitempty"`


	// ServiceLevelPercent
	ServiceLevelPercent *float32 `json:"serviceLevelPercent,omitempty"`


	// AverageSpeedOfAnswerSeconds
	AverageSpeedOfAnswerSeconds *float32 `json:"averageSpeedOfAnswerSeconds,omitempty"`

}

func (o *Wfmintradaydataupdatetopicintradayhistoricalqueuedata) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmintradaydataupdatetopicintradayhistoricalqueuedata
	
	return json.Marshal(&struct { 
		Offered *int `json:"offered,omitempty"`
		
		Completed *int `json:"completed,omitempty"`
		
		Answered *int `json:"answered,omitempty"`
		
		Abandoned *int `json:"abandoned,omitempty"`
		
		AverageTalkTimeSeconds *float32 `json:"averageTalkTimeSeconds,omitempty"`
		
		AverageAfterCallWorkSeconds *float32 `json:"averageAfterCallWorkSeconds,omitempty"`
		
		ServiceLevelPercent *float32 `json:"serviceLevelPercent,omitempty"`
		
		AverageSpeedOfAnswerSeconds *float32 `json:"averageSpeedOfAnswerSeconds,omitempty"`
		*Alias
	}{ 
		Offered: o.Offered,
		
		Completed: o.Completed,
		
		Answered: o.Answered,
		
		Abandoned: o.Abandoned,
		
		AverageTalkTimeSeconds: o.AverageTalkTimeSeconds,
		
		AverageAfterCallWorkSeconds: o.AverageAfterCallWorkSeconds,
		
		ServiceLevelPercent: o.ServiceLevelPercent,
		
		AverageSpeedOfAnswerSeconds: o.AverageSpeedOfAnswerSeconds,
		Alias:    (*Alias)(o),
	})
}

func (o *Wfmintradaydataupdatetopicintradayhistoricalqueuedata) UnmarshalJSON(b []byte) error {
	var WfmintradaydataupdatetopicintradayhistoricalqueuedataMap map[string]interface{}
	err := json.Unmarshal(b, &WfmintradaydataupdatetopicintradayhistoricalqueuedataMap)
	if err != nil {
		return err
	}
	
	if Offered, ok := WfmintradaydataupdatetopicintradayhistoricalqueuedataMap["offered"].(float64); ok {
		OfferedInt := int(Offered)
		o.Offered = &OfferedInt
	}
	
	if Completed, ok := WfmintradaydataupdatetopicintradayhistoricalqueuedataMap["completed"].(float64); ok {
		CompletedInt := int(Completed)
		o.Completed = &CompletedInt
	}
	
	if Answered, ok := WfmintradaydataupdatetopicintradayhistoricalqueuedataMap["answered"].(float64); ok {
		AnsweredInt := int(Answered)
		o.Answered = &AnsweredInt
	}
	
	if Abandoned, ok := WfmintradaydataupdatetopicintradayhistoricalqueuedataMap["abandoned"].(float64); ok {
		AbandonedInt := int(Abandoned)
		o.Abandoned = &AbandonedInt
	}
	
	if AverageTalkTimeSeconds, ok := WfmintradaydataupdatetopicintradayhistoricalqueuedataMap["averageTalkTimeSeconds"].(float64); ok {
		AverageTalkTimeSecondsFloat32 := float32(AverageTalkTimeSeconds)
		o.AverageTalkTimeSeconds = &AverageTalkTimeSecondsFloat32
	}
	
	if AverageAfterCallWorkSeconds, ok := WfmintradaydataupdatetopicintradayhistoricalqueuedataMap["averageAfterCallWorkSeconds"].(float64); ok {
		AverageAfterCallWorkSecondsFloat32 := float32(AverageAfterCallWorkSeconds)
		o.AverageAfterCallWorkSeconds = &AverageAfterCallWorkSecondsFloat32
	}
	
	if ServiceLevelPercent, ok := WfmintradaydataupdatetopicintradayhistoricalqueuedataMap["serviceLevelPercent"].(float64); ok {
		ServiceLevelPercentFloat32 := float32(ServiceLevelPercent)
		o.ServiceLevelPercent = &ServiceLevelPercentFloat32
	}
	
	if AverageSpeedOfAnswerSeconds, ok := WfmintradaydataupdatetopicintradayhistoricalqueuedataMap["averageSpeedOfAnswerSeconds"].(float64); ok {
		AverageSpeedOfAnswerSecondsFloat32 := float32(AverageSpeedOfAnswerSeconds)
		o.AverageSpeedOfAnswerSeconds = &AverageSpeedOfAnswerSecondsFloat32
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmintradaydataupdatetopicintradayhistoricalqueuedata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
