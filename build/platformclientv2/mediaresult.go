package platformclientv2
import (
	"encoding/json"
)

// Mediaresult
type Mediaresult struct { 
	// MediaUri
	MediaUri *string `json:"mediaUri,omitempty"`


	// WaveformData
	WaveformData *[]float32 `json:"waveformData,omitempty"`

}

// String returns a JSON representation of the model
func (o *Mediaresult) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
