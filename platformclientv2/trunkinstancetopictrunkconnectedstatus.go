package platformclientv2
import (
	"time"
	"encoding/json"
)

// Trunkinstancetopictrunkconnectedstatus
type Trunkinstancetopictrunkconnectedstatus struct { 
	// Connected
	Connected *bool `json:"connected,omitempty"`


	// ConnectedStateTime
	ConnectedStateTime *time.Time `json:"connectedStateTime,omitempty"`

}

// String returns a JSON representation of the model
func (o *Trunkinstancetopictrunkconnectedstatus) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
