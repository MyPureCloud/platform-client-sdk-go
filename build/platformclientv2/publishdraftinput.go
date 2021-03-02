package platformclientv2
import (
	"encoding/json"
)

// Publishdraftinput - Draft to be published
type Publishdraftinput struct { 
	// Version - The current draft version.
	Version *int `json:"version,omitempty"`

}

// String returns a JSON representation of the model
func (o *Publishdraftinput) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
