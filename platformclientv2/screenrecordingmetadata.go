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

func (o *Screenrecordingmetadata) MarshalJSON() ([]byte, error) {
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
		TrackId: o.TrackId,
		
		MediaId: o.MediaId,
		
		ScreenId: o.ScreenId,
		
		OriginX: o.OriginX,
		
		OriginY: o.OriginY,
		
		Primary: o.Primary,
		
		Main: o.Main,
		Alias:    (*Alias)(o),
	})
}

func (o *Screenrecordingmetadata) UnmarshalJSON(b []byte) error {
	var ScreenrecordingmetadataMap map[string]interface{}
	err := json.Unmarshal(b, &ScreenrecordingmetadataMap)
	if err != nil {
		return err
	}
	
	if TrackId, ok := ScreenrecordingmetadataMap["trackId"].(string); ok {
		o.TrackId = &TrackId
	}
	
	if MediaId, ok := ScreenrecordingmetadataMap["mediaId"].(string); ok {
		o.MediaId = &MediaId
	}
	
	if ScreenId, ok := ScreenrecordingmetadataMap["screenId"].(string); ok {
		o.ScreenId = &ScreenId
	}
	
	if OriginX, ok := ScreenrecordingmetadataMap["originX"].(float64); ok {
		OriginXInt := int(OriginX)
		o.OriginX = &OriginXInt
	}
	
	if OriginY, ok := ScreenrecordingmetadataMap["originY"].(float64); ok {
		OriginYInt := int(OriginY)
		o.OriginY = &OriginYInt
	}
	
	if Primary, ok := ScreenrecordingmetadataMap["primary"].(bool); ok {
		o.Primary = &Primary
	}
	
	if Main, ok := ScreenrecordingmetadataMap["main"].(bool); ok {
		o.Main = &Main
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Screenrecordingmetadata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
