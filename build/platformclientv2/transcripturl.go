package platformclientv2
import (
	"encoding/json"
)

// Transcripturl
type Transcripturl struct { 
	// Url - The pre-signed S3 URL of the transcript
	Url *string `json:"url,omitempty"`

}

// String returns a JSON representation of the model
func (o *Transcripturl) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
