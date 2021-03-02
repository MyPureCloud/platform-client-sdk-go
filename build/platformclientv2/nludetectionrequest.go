package platformclientv2
import (
	"encoding/json"
)

// Nludetectionrequest
type Nludetectionrequest struct { 
	// Input - The input subject to NLU detection.
	Input *Nludetectioninput `json:"input,omitempty"`


	// Context - The context for the input to NLU detection.
	Context *Nludetectioncontext `json:"context,omitempty"`

}

// String returns a JSON representation of the model
func (o *Nludetectionrequest) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
