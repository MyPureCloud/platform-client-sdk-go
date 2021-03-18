package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Wfmintradaydataupdatetopicintradayscheduledata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
