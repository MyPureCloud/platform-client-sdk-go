package platformclientv2
import (
	"encoding/json"
)

// Recordingtranscodecompletetopicmediaresult
type Recordingtranscodecompletetopicmediaresult struct { 
	// ChannelId
	ChannelId *string `json:"channelId,omitempty"`


	// WaveUri
	WaveUri *string `json:"waveUri,omitempty"`


	// MediaUri
	MediaUri *string `json:"mediaUri,omitempty"`


	// WaveformData
	WaveformData *[]float32 `json:"waveformData,omitempty"`

}

// String returns a JSON representation of the model
func (o *Recordingtranscodecompletetopicmediaresult) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
