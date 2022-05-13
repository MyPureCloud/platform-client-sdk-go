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

func (o *Recordingtranscodecompletetopicmediaresult) MarshalJSON() ([]byte, error) {
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
		ChannelId: o.ChannelId,
		
		WaveUri: o.WaveUri,
		
		MediaUri: o.MediaUri,
		
		WaveformData: o.WaveformData,
		Alias:    (*Alias)(o),
	})
}

func (o *Recordingtranscodecompletetopicmediaresult) UnmarshalJSON(b []byte) error {
	var RecordingtranscodecompletetopicmediaresultMap map[string]interface{}
	err := json.Unmarshal(b, &RecordingtranscodecompletetopicmediaresultMap)
	if err != nil {
		return err
	}
	
	if ChannelId, ok := RecordingtranscodecompletetopicmediaresultMap["channelId"].(string); ok {
		o.ChannelId = &ChannelId
	}
    
	if WaveUri, ok := RecordingtranscodecompletetopicmediaresultMap["waveUri"].(string); ok {
		o.WaveUri = &WaveUri
	}
    
	if MediaUri, ok := RecordingtranscodecompletetopicmediaresultMap["mediaUri"].(string); ok {
		o.MediaUri = &MediaUri
	}
    
	if WaveformData, ok := RecordingtranscodecompletetopicmediaresultMap["waveformData"].([]interface{}); ok {
		WaveformDataString, _ := json.Marshal(WaveformData)
		json.Unmarshal(WaveformDataString, &o.WaveformData)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Recordingtranscodecompletetopicmediaresult) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
