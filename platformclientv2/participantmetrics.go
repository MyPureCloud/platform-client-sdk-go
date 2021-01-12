package platformclientv2
import (
	"encoding/json"
)

// Participantmetrics
type Participantmetrics struct { }

// String returns a JSON representation of the model
func (o *Participantmetrics) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
