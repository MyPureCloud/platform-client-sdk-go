package platformclientv2
import (
	"encoding/json"
)

// Callroute
type Callroute struct { 
	// Targets - A list of CallTargets to be called when the CallRoute is executed
	Targets *[]Calltarget `json:"targets,omitempty"`

}

// String returns a JSON representation of the model
func (o *Callroute) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
