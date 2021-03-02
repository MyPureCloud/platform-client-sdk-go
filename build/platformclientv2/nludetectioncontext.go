package platformclientv2
import (
	"encoding/json"
)

// Nludetectioncontext
type Nludetectioncontext struct { 
	// Intent - Restrict detection to this intent.
	Intent *Contextintent `json:"intent,omitempty"`


	// Entity - Use this entity to restrict detection.
	Entity *Contextentity `json:"entity,omitempty"`

}

// String returns a JSON representation of the model
func (o *Nludetectioncontext) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
