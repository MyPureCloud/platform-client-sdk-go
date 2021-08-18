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

func (u *Buintradayscheduledata) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Buintradayscheduledata

	

	return json.Marshal(&struct { 
		OnQueueTimeSeconds *int `json:"onQueueTimeSeconds,omitempty"`
		*Alias
	}{ 
		OnQueueTimeSeconds: u.OnQueueTimeSeconds,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Buintradayscheduledata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
