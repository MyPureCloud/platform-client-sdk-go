package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmintradaydataupdatetopicintradayhistoricalagentdata
type Wfmintradaydataupdatetopicintradayhistoricalagentdata struct { 
	// OnQueueTimeSeconds
	OnQueueTimeSeconds *float32 `json:"onQueueTimeSeconds,omitempty"`


	// InteractingTimeSeconds
	InteractingTimeSeconds *float32 `json:"interactingTimeSeconds,omitempty"`

}

func (o *Wfmintradaydataupdatetopicintradayhistoricalagentdata) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmintradaydataupdatetopicintradayhistoricalagentdata
	
	return json.Marshal(&struct { 
		OnQueueTimeSeconds *float32 `json:"onQueueTimeSeconds,omitempty"`
		
		InteractingTimeSeconds *float32 `json:"interactingTimeSeconds,omitempty"`
		*Alias
	}{ 
		OnQueueTimeSeconds: o.OnQueueTimeSeconds,
		
		InteractingTimeSeconds: o.InteractingTimeSeconds,
		Alias:    (*Alias)(o),
	})
}

func (o *Wfmintradaydataupdatetopicintradayhistoricalagentdata) UnmarshalJSON(b []byte) error {
	var WfmintradaydataupdatetopicintradayhistoricalagentdataMap map[string]interface{}
	err := json.Unmarshal(b, &WfmintradaydataupdatetopicintradayhistoricalagentdataMap)
	if err != nil {
		return err
	}
	
	if OnQueueTimeSeconds, ok := WfmintradaydataupdatetopicintradayhistoricalagentdataMap["onQueueTimeSeconds"].(float64); ok {
		OnQueueTimeSecondsFloat32 := float32(OnQueueTimeSeconds)
		o.OnQueueTimeSeconds = &OnQueueTimeSecondsFloat32
	}
    
	if InteractingTimeSeconds, ok := WfmintradaydataupdatetopicintradayhistoricalagentdataMap["interactingTimeSeconds"].(float64); ok {
		InteractingTimeSecondsFloat32 := float32(InteractingTimeSeconds)
		o.InteractingTimeSeconds = &InteractingTimeSecondsFloat32
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmintradaydataupdatetopicintradayhistoricalagentdata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
