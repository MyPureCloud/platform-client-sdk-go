package platformclientv2
import (
	"encoding/json"
)

// Phonechangetopiczoneoffset
type Phonechangetopiczoneoffset struct { 
	// TotalSeconds
	TotalSeconds *int `json:"totalSeconds,omitempty"`

}

// String returns a JSON representation of the model
func (o *Phonechangetopiczoneoffset) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
