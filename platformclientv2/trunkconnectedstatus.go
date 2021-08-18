package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
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

func (u *Trunkconnectedstatus) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Trunkconnectedstatus

	
	ConnectedStateTime := new(string)
	if u.ConnectedStateTime != nil {
		
		*ConnectedStateTime = timeutil.Strftime(u.ConnectedStateTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ConnectedStateTime = nil
	}
	

	return json.Marshal(&struct { 
		Connected *bool `json:"connected,omitempty"`
		
		ConnectedStateTime *string `json:"connectedStateTime,omitempty"`
		*Alias
	}{ 
		Connected: u.Connected,
		
		ConnectedStateTime: ConnectedStateTime,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Trunkconnectedstatus) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
