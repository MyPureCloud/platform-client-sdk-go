package platformclientv2
import (
	"encoding/json"
)

// Usergreetingeventgreetingaudiofile
type Usergreetingeventgreetingaudiofile struct { 
	// DurationMilliseconds
	DurationMilliseconds *int `json:"durationMilliseconds,omitempty"`


	// SizeBytes
	SizeBytes *int `json:"sizeBytes,omitempty"`

}

// String returns a JSON representation of the model
func (o *Usergreetingeventgreetingaudiofile) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
