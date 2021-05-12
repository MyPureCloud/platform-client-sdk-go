package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Trunkconnectedstatus
type Trunkconnectedstatus struct { 
	// Connected
	Connected *bool `json:"connected,omitempty"`


	// ConnectedStateTime - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ConnectedStateTime *time.Time `json:"connectedStateTime,omitempty"`

}

// String returns a JSON representation of the model
func (o *Trunkconnectedstatus) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
