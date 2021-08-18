package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Trunkinstancetopictrunkconnectedstatus
type Trunkinstancetopictrunkconnectedstatus struct { 
	// Connected
	Connected *bool `json:"connected,omitempty"`


	// ConnectedStateTime
	ConnectedStateTime *time.Time `json:"connectedStateTime,omitempty"`

}

func (u *Trunkinstancetopictrunkconnectedstatus) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Trunkinstancetopictrunkconnectedstatus

	
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
func (o *Trunkinstancetopictrunkconnectedstatus) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
