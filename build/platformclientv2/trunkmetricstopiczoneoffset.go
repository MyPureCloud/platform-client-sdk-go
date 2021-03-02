package platformclientv2
import (
	"encoding/json"
)

// Trunkmetricstopiczoneoffset
type Trunkmetricstopiczoneoffset struct { 
	// TotalSeconds
	TotalSeconds *int `json:"totalSeconds,omitempty"`

}

// String returns a JSON representation of the model
func (o *Trunkmetricstopiczoneoffset) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
