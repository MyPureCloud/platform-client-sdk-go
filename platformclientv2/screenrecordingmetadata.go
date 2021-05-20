package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Screenrecordingmetadata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
