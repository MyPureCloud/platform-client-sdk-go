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

func (o *Trunkinstancetopictrunkconnectedstatus) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Trunkinstancetopictrunkconnectedstatus
	
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

func (o *Trunkinstancetopictrunkconnectedstatus) UnmarshalJSON(b []byte) error {
	var TrunkinstancetopictrunkconnectedstatusMap map[string]interface{}
	err := json.Unmarshal(b, &TrunkinstancetopictrunkconnectedstatusMap)
	if err != nil {
		return err
	}
	
	if Connected, ok := TrunkinstancetopictrunkconnectedstatusMap["connected"].(bool); ok {
		o.Connected = &Connected
	}
	
	if connectedStateTimeString, ok := TrunkinstancetopictrunkconnectedstatusMap["connectedStateTime"].(string); ok {
		ConnectedStateTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", connectedStateTimeString)
		o.ConnectedStateTime = &ConnectedStateTime
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Trunkinstancetopictrunkconnectedstatus) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
