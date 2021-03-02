package platformclientv2
import (
	"encoding/json"
)

// Greetingaudiofile
type Greetingaudiofile struct { 
	// DurationMilliseconds
	DurationMilliseconds *int `json:"durationMilliseconds,omitempty"`


	// SizeBytes
	SizeBytes *int `json:"sizeBytes,omitempty"`


	// SelfUri
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Greetingaudiofile) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
