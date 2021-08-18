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

func (u *Wfmintradaydataupdatetopicintradayscheduledata) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmintradaydataupdatetopicintradayscheduledata

	

	return json.Marshal(&struct { 
		OnQueueTimeSeconds *int `json:"onQueueTimeSeconds,omitempty"`
		
		ScheduledTimeSeconds *int `json:"scheduledTimeSeconds,omitempty"`
		*Alias
	}{ 
		OnQueueTimeSeconds: u.OnQueueTimeSeconds,
		
		ScheduledTimeSeconds: u.ScheduledTimeSeconds,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Wfmintradaydataupdatetopicintradayscheduledata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
