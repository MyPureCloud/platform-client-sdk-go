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

func (u *Wfmintradaydataupdatetopicintradayperformancepredictionagentdata) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmintradaydataupdatetopicintradayperformancepredictionagentdata

	

	return json.Marshal(&struct { 
		InteractingTimeSeconds *float32 `json:"interactingTimeSeconds,omitempty"`
		*Alias
	}{ 
		InteractingTimeSeconds: u.InteractingTimeSeconds,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Wfmintradaydataupdatetopicintradayperformancepredictionagentdata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
