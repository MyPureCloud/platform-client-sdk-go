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

func (u *Voicemailmediainfo) MarshalJSON() ([]byte, error) {
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
		Id: u.Id,
		
		MediaFileUri: u.MediaFileUri,
		
		MediaImageUri: u.MediaImageUri,
		
		WaveformData: u.WaveformData,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Voicemailmediainfo) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
