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

func (u *Mediaresult) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Mediaresult

	

	return json.Marshal(&struct { 
		MediaUri *string `json:"mediaUri,omitempty"`
		
		WaveformData *[]float32 `json:"waveformData,omitempty"`
		*Alias
	}{ 
		MediaUri: u.MediaUri,
		
		WaveformData: u.WaveformData,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Mediaresult) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
