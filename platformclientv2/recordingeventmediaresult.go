package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Recordingeventmediaresult
type Recordingeventmediaresult struct { 
	// ChannelId
	ChannelId *string `json:"channelId,omitempty"`


	// WaveUri
	WaveUri *string `json:"waveUri,omitempty"`


	// MediaUri
	MediaUri *string `json:"mediaUri,omitempty"`


	// WaveformData
	WaveformData *[]float32 `json:"waveformData,omitempty"`

}

func (o *Recordingeventmediaresult) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Recordingeventmediaresult
	
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

func (o *Recordingeventmediaresult) UnmarshalJSON(b []byte) error {
	var RecordingeventmediaresultMap map[string]interface{}
	err := json.Unmarshal(b, &RecordingeventmediaresultMap)
	if err != nil {
		return err
	}
	
	if ChannelId, ok := RecordingeventmediaresultMap["channelId"].(string); ok {
		o.ChannelId = &ChannelId
	}
	
	if WaveUri, ok := RecordingeventmediaresultMap["waveUri"].(string); ok {
		o.WaveUri = &WaveUri
	}
	
	if MediaUri, ok := RecordingeventmediaresultMap["mediaUri"].(string); ok {
		o.MediaUri = &MediaUri
	}
	
	if WaveformData, ok := RecordingeventmediaresultMap["waveformData"].([]interface{}); ok {
		WaveformDataString, _ := json.Marshal(WaveformData)
		json.Unmarshal(WaveformDataString, &o.WaveformData)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Recordingeventmediaresult) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
