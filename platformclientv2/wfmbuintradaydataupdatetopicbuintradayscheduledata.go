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

func (u *Wfmbuintradaydataupdatetopicbuintradayscheduledata) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmbuintradaydataupdatetopicbuintradayscheduledata

	

	return json.Marshal(&struct { 
		OnQueueTimeSeconds *int `json:"onQueueTimeSeconds,omitempty"`
		*Alias
	}{ 
		OnQueueTimeSeconds: u.OnQueueTimeSeconds,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Wfmbuintradaydataupdatetopicbuintradayscheduledata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
