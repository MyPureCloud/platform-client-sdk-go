package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Buintradayscheduledata
type Buintradayscheduledata struct { 
	// OnQueueTimeSeconds - The total on-queue time in seconds for all agents in this group
	OnQueueTimeSeconds *int `json:"onQueueTimeSeconds,omitempty"`

}

func (o *Buintradayscheduledata) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Buintradayscheduledata
	
	return json.Marshal(&struct { 
		OnQueueTimeSeconds *int `json:"onQueueTimeSeconds,omitempty"`
		*Alias
	}{ 
		OnQueueTimeSeconds: o.OnQueueTimeSeconds,
		Alias:    (*Alias)(o),
	})
}

func (o *Buintradayscheduledata) UnmarshalJSON(b []byte) error {
	var BuintradayscheduledataMap map[string]interface{}
	err := json.Unmarshal(b, &BuintradayscheduledataMap)
	if err != nil {
		return err
	}
	
	if OnQueueTimeSeconds, ok := BuintradayscheduledataMap["onQueueTimeSeconds"].(float64); ok {
		OnQueueTimeSecondsInt := int(OnQueueTimeSeconds)
		o.OnQueueTimeSeconds = &OnQueueTimeSecondsInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Buintradayscheduledata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
