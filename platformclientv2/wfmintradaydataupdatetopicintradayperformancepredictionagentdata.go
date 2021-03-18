package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmintradaydataupdatetopicintradayperformancepredictionagentdata
type Wfmintradaydataupdatetopicintradayperformancepredictionagentdata struct { 
	// InteractingTimeSeconds
	InteractingTimeSeconds *float32 `json:"interactingTimeSeconds,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wfmintradaydataupdatetopicintradayperformancepredictionagentdata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
