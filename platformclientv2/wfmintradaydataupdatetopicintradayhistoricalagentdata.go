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

func (u *Wfmintradaydataupdatetopicintradayhistoricalagentdata) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmintradaydataupdatetopicintradayhistoricalagentdata

	

	return json.Marshal(&struct { 
		OnQueueTimeSeconds *float32 `json:"onQueueTimeSeconds,omitempty"`
		
		InteractingTimeSeconds *float32 `json:"interactingTimeSeconds,omitempty"`
		*Alias
	}{ 
		OnQueueTimeSeconds: u.OnQueueTimeSeconds,
		
		InteractingTimeSeconds: u.InteractingTimeSeconds,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Wfmintradaydataupdatetopicintradayhistoricalagentdata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
