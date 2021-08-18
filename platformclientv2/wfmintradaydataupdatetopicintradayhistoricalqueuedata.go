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

func (u *Wfmintradaydataupdatetopicintradayhistoricalqueuedata) MarshalJSON() ([]byte, error) {
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
		Offered: u.Offered,
		
		Completed: u.Completed,
		
		Answered: u.Answered,
		
		Abandoned: u.Abandoned,
		
		AverageTalkTimeSeconds: u.AverageTalkTimeSeconds,
		
		AverageAfterCallWorkSeconds: u.AverageAfterCallWorkSeconds,
		
		ServiceLevelPercent: u.ServiceLevelPercent,
		
		AverageSpeedOfAnswerSeconds: u.AverageSpeedOfAnswerSeconds,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Wfmintradaydataupdatetopicintradayhistoricalqueuedata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
