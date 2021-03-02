package platformclientv2
import (
	"encoding/json"
)

// Groupgreetingeventgreetingaudiofile
type Groupgreetingeventgreetingaudiofile struct { 
	// DurationMilliseconds
	DurationMilliseconds *int `json:"durationMilliseconds,omitempty"`


	// SizeBytes
	SizeBytes *int `json:"sizeBytes,omitempty"`

}

// String returns a JSON representation of the model
func (o *Groupgreetingeventgreetingaudiofile) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
