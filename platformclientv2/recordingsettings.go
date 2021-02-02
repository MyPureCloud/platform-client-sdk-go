package platformclientv2
import (
	"encoding/json"
)

// Recordingsettings
type Recordingsettings struct { 
	// MaxSimultaneousStreams
	MaxSimultaneousStreams *int `json:"maxSimultaneousStreams,omitempty"`

}

// String returns a JSON representation of the model
func (o *Recordingsettings) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
