package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmintradaydataupdatetopicintradayscheduledata
type Wfmintradaydataupdatetopicintradayscheduledata struct { 
	// OnQueueTimeSeconds
	OnQueueTimeSeconds *int `json:"onQueueTimeSeconds,omitempty"`


	// ScheduledTimeSeconds
	ScheduledTimeSeconds *int `json:"scheduledTimeSeconds,omitempty"`

}

func (o *Wfmintradaydataupdatetopicintradayscheduledata) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmintradaydataupdatetopicintradayscheduledata
	
	return json.Marshal(&struct { 
		OnQueueTimeSeconds *int `json:"onQueueTimeSeconds,omitempty"`
		
		ScheduledTimeSeconds *int `json:"scheduledTimeSeconds,omitempty"`
		*Alias
	}{ 
		OnQueueTimeSeconds: o.OnQueueTimeSeconds,
		
		ScheduledTimeSeconds: o.ScheduledTimeSeconds,
		Alias:    (*Alias)(o),
	})
}

func (o *Wfmintradaydataupdatetopicintradayscheduledata) UnmarshalJSON(b []byte) error {
	var WfmintradaydataupdatetopicintradayscheduledataMap map[string]interface{}
	err := json.Unmarshal(b, &WfmintradaydataupdatetopicintradayscheduledataMap)
	if err != nil {
		return err
	}
	
	if OnQueueTimeSeconds, ok := WfmintradaydataupdatetopicintradayscheduledataMap["onQueueTimeSeconds"].(float64); ok {
		OnQueueTimeSecondsInt := int(OnQueueTimeSeconds)
		o.OnQueueTimeSeconds = &OnQueueTimeSecondsInt
	}
	
	if ScheduledTimeSeconds, ok := WfmintradaydataupdatetopicintradayscheduledataMap["scheduledTimeSeconds"].(float64); ok {
		ScheduledTimeSecondsInt := int(ScheduledTimeSeconds)
		o.ScheduledTimeSeconds = &ScheduledTimeSecondsInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmintradaydataupdatetopicintradayscheduledata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
