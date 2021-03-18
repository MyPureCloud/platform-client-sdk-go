package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmbuintradaydataupdatetopicbuintradayscheduledata
type Wfmbuintradaydataupdatetopicbuintradayscheduledata struct { 
	// OnQueueTimeSeconds
	OnQueueTimeSeconds *int `json:"onQueueTimeSeconds,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wfmbuintradaydataupdatetopicbuintradayscheduledata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
