package platformclientv2
import (
	"encoding/json"
)

// Callforwardingeventcall
type Callforwardingeventcall struct { 
	// Targets
	Targets *[]Callforwardingeventtarget `json:"targets,omitempty"`

}

// String returns a JSON representation of the model
func (o *Callforwardingeventcall) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
