package platformclientv2
import (
	"encoding/json"
)

// Edgemetricstopiczoneoffset
type Edgemetricstopiczoneoffset struct { 
	// TotalSeconds
	TotalSeconds *int `json:"totalSeconds,omitempty"`

}

// String returns a JSON representation of the model
func (o *Edgemetricstopiczoneoffset) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
