package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmintradaydataupdatetopicintradayperformancepredictionagentdata
type Wfmintradaydataupdatetopicintradayperformancepredictionagentdata struct { 
	// InteractingTimeSeconds
	InteractingTimeSeconds *float32 `json:"interactingTimeSeconds,omitempty"`

}

func (o *Wfmintradaydataupdatetopicintradayperformancepredictionagentdata) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmintradaydataupdatetopicintradayperformancepredictionagentdata
	
	return json.Marshal(&struct { 
		InteractingTimeSeconds *float32 `json:"interactingTimeSeconds,omitempty"`
		*Alias
	}{ 
		InteractingTimeSeconds: o.InteractingTimeSeconds,
		Alias:    (*Alias)(o),
	})
}

func (o *Wfmintradaydataupdatetopicintradayperformancepredictionagentdata) UnmarshalJSON(b []byte) error {
	var WfmintradaydataupdatetopicintradayperformancepredictionagentdataMap map[string]interface{}
	err := json.Unmarshal(b, &WfmintradaydataupdatetopicintradayperformancepredictionagentdataMap)
	if err != nil {
		return err
	}
	
	if InteractingTimeSeconds, ok := WfmintradaydataupdatetopicintradayperformancepredictionagentdataMap["interactingTimeSeconds"].(float64); ok {
		InteractingTimeSecondsFloat32 := float32(InteractingTimeSeconds)
		o.InteractingTimeSeconds = &InteractingTimeSecondsFloat32
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmintradaydataupdatetopicintradayperformancepredictionagentdata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
