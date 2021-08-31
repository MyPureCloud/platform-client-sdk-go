package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
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

func (o *Voicemailmediainfo) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Voicemailmediainfo
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		MediaFileUri *string `json:"mediaFileUri,omitempty"`
		
		MediaImageUri *string `json:"mediaImageUri,omitempty"`
		
		WaveformData *[]float32 `json:"waveformData,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		MediaFileUri: o.MediaFileUri,
		
		MediaImageUri: o.MediaImageUri,
		
		WaveformData: o.WaveformData,
		Alias:    (*Alias)(o),
	})
}

func (o *Voicemailmediainfo) UnmarshalJSON(b []byte) error {
	var VoicemailmediainfoMap map[string]interface{}
	err := json.Unmarshal(b, &VoicemailmediainfoMap)
	if err != nil {
		return err
	}
	
	if Id, ok := VoicemailmediainfoMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if MediaFileUri, ok := VoicemailmediainfoMap["mediaFileUri"].(string); ok {
		o.MediaFileUri = &MediaFileUri
	}
	
	if MediaImageUri, ok := VoicemailmediainfoMap["mediaImageUri"].(string); ok {
		o.MediaImageUri = &MediaImageUri
	}
	
	if WaveformData, ok := VoicemailmediainfoMap["waveformData"].([]interface{}); ok {
		WaveformDataString, _ := json.Marshal(WaveformData)
		json.Unmarshal(WaveformDataString, &o.WaveformData)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Voicemailmediainfo) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
