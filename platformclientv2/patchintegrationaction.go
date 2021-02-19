package platformclientv2
import (
	"encoding/json"
)

// Patchintegrationaction
type Patchintegrationaction struct { 
	// Id - ID of the integration action to be invoked.
	Id *string `json:"id,omitempty"`

}

// String returns a JSON representation of the model
func (o *Patchintegrationaction) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
