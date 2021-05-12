package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Buintradayscheduledata
type Buintradayscheduledata struct { 
	// OnQueueTimeSeconds - The total on-queue time in seconds for all agents in this group
	OnQueueTimeSeconds *int `json:"onQueueTimeSeconds,omitempty"`

}

// String returns a JSON representation of the model
func (o *Buintradayscheduledata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
