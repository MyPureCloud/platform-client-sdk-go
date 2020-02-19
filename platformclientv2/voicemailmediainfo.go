package platformclientv2
import (
	"encoding/json"
)

// Voicemailmediainfo
type Voicemailmediainfo struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// MediaFileUri
	MediaFileUri *string `json:"mediaFileUri,omitempty"`


	// MediaImageUri
	MediaImageUri *string `json:"mediaImageUri,omitempty"`


	// WaveformData
	WaveformData *[]float32 `json:"waveformData,omitempty"`

}

// String returns a JSON representation of the model
func (o *Voicemailmediainfo) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
