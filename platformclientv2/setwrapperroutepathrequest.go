package platformclientv2
import (
	"encoding/json"
)

// Setwrapperroutepathrequest
type Setwrapperroutepathrequest struct { 
	// Values
	Values *[]Routepathrequest `json:"values,omitempty"`

}

// String returns a JSON representation of the model
func (o *Setwrapperroutepathrequest) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
