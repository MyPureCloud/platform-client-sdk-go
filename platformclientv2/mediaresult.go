package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Mediaresult
type Mediaresult struct { 
	// MediaUri
	MediaUri *string `json:"mediaUri,omitempty"`


	// WaveformData
	WaveformData *[]float32 `json:"waveformData,omitempty"`

}

func (o *Mediaresult) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Mediaresult
	
	return json.Marshal(&struct { 
		MediaUri *string `json:"mediaUri,omitempty"`
		
		WaveformData *[]float32 `json:"waveformData,omitempty"`
		*Alias
	}{ 
		MediaUri: o.MediaUri,
		
		WaveformData: o.WaveformData,
		Alias:    (*Alias)(o),
	})
}

func (o *Mediaresult) UnmarshalJSON(b []byte) error {
	var MediaresultMap map[string]interface{}
	err := json.Unmarshal(b, &MediaresultMap)
	if err != nil {
		return err
	}
	
	if MediaUri, ok := MediaresultMap["mediaUri"].(string); ok {
		o.MediaUri = &MediaUri
	}
    
	if WaveformData, ok := MediaresultMap["waveformData"].([]interface{}); ok {
		WaveformDataString, _ := json.Marshal(WaveformData)
		json.Unmarshal(WaveformDataString, &o.WaveformData)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Mediaresult) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
