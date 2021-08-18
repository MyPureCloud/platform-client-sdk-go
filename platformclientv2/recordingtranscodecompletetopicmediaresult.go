package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
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

func (u *Recordingtranscodecompletetopicmediaresult) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Recordingtranscodecompletetopicmediaresult

	

	return json.Marshal(&struct { 
		ChannelId *string `json:"channelId,omitempty"`
		
		WaveUri *string `json:"waveUri,omitempty"`
		
		MediaUri *string `json:"mediaUri,omitempty"`
		
		WaveformData *[]float32 `json:"waveformData,omitempty"`
		*Alias
	}{ 
		ChannelId: u.ChannelId,
		
		WaveUri: u.WaveUri,
		
		MediaUri: u.MediaUri,
		
		WaveformData: u.WaveformData,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Recordingtranscodecompletetopicmediaresult) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
