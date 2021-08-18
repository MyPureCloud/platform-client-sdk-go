package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Screenrecordingmetadata
type Screenrecordingmetadata struct { 
	// TrackId
	TrackId *string `json:"trackId,omitempty"`


	// MediaId
	MediaId *string `json:"mediaId,omitempty"`


	// ScreenId
	ScreenId *string `json:"screenId,omitempty"`


	// OriginX
	OriginX *int `json:"originX,omitempty"`


	// OriginY
	OriginY *int `json:"originY,omitempty"`


	// Primary
	Primary *bool `json:"primary,omitempty"`


	// Main
	Main *bool `json:"main,omitempty"`

}

func (u *Screenrecordingmetadata) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Screenrecordingmetadata

	

	return json.Marshal(&struct { 
		TrackId *string `json:"trackId,omitempty"`
		
		MediaId *string `json:"mediaId,omitempty"`
		
		ScreenId *string `json:"screenId,omitempty"`
		
		OriginX *int `json:"originX,omitempty"`
		
		OriginY *int `json:"originY,omitempty"`
		
		Primary *bool `json:"primary,omitempty"`
		
		Main *bool `json:"main,omitempty"`
		*Alias
	}{ 
		TrackId: u.TrackId,
		
		MediaId: u.MediaId,
		
		ScreenId: u.ScreenId,
		
		OriginX: u.OriginX,
		
		OriginY: u.OriginY,
		
		Primary: u.Primary,
		
		Main: u.Main,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Screenrecordingmetadata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
