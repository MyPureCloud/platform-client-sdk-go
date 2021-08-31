package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmbuintradaydataupdatetopicbuintradayscheduledata
type Wfmbuintradaydataupdatetopicbuintradayscheduledata struct { 
	// OnQueueTimeSeconds
	OnQueueTimeSeconds *int `json:"onQueueTimeSeconds,omitempty"`

}

func (o *Wfmbuintradaydataupdatetopicbuintradayscheduledata) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmbuintradaydataupdatetopicbuintradayscheduledata
	
	return json.Marshal(&struct { 
		OnQueueTimeSeconds *int `json:"onQueueTimeSeconds,omitempty"`
		*Alias
	}{ 
		OnQueueTimeSeconds: o.OnQueueTimeSeconds,
		Alias:    (*Alias)(o),
	})
}

func (o *Wfmbuintradaydataupdatetopicbuintradayscheduledata) UnmarshalJSON(b []byte) error {
	var WfmbuintradaydataupdatetopicbuintradayscheduledataMap map[string]interface{}
	err := json.Unmarshal(b, &WfmbuintradaydataupdatetopicbuintradayscheduledataMap)
	if err != nil {
		return err
	}
	
	if OnQueueTimeSeconds, ok := WfmbuintradaydataupdatetopicbuintradayscheduledataMap["onQueueTimeSeconds"].(float64); ok {
		OnQueueTimeSecondsInt := int(OnQueueTimeSeconds)
		o.OnQueueTimeSeconds = &OnQueueTimeSecondsInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmbuintradaydataupdatetopicbuintradayscheduledata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
