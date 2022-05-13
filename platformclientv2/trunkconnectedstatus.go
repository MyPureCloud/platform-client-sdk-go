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

func (o *Trunkconnectedstatus) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Trunkconnectedstatus
	
	ConnectedStateTime := new(string)
	if o.ConnectedStateTime != nil {
		
		*ConnectedStateTime = timeutil.Strftime(o.ConnectedStateTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ConnectedStateTime = nil
	}
	
	return json.Marshal(&struct { 
		Connected *bool `json:"connected,omitempty"`
		
		ConnectedStateTime *string `json:"connectedStateTime,omitempty"`
		*Alias
	}{ 
		Connected: o.Connected,
		
		ConnectedStateTime: ConnectedStateTime,
		Alias:    (*Alias)(o),
	})
}

func (o *Trunkconnectedstatus) UnmarshalJSON(b []byte) error {
	var TrunkconnectedstatusMap map[string]interface{}
	err := json.Unmarshal(b, &TrunkconnectedstatusMap)
	if err != nil {
		return err
	}
	
	if Connected, ok := TrunkconnectedstatusMap["connected"].(bool); ok {
		o.Connected = &Connected
	}
    
	if connectedStateTimeString, ok := TrunkconnectedstatusMap["connectedStateTime"].(string); ok {
		ConnectedStateTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", connectedStateTimeString)
		o.ConnectedStateTime = &ConnectedStateTime
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Trunkconnectedstatus) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
