package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Documentthumbnail
type Documentthumbnail struct { 
	// Resolution
	Resolution *string `json:"resolution,omitempty"`


	// ImageUri
	ImageUri *string `json:"imageUri,omitempty"`


	// Height
	Height *int `json:"height,omitempty"`


	// Width
	Width *int `json:"width,omitempty"`

}

func (o *Documentthumbnail) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Documentthumbnail
	
	return json.Marshal(&struct { 
		Resolution *string `json:"resolution,omitempty"`
		
		ImageUri *string `json:"imageUri,omitempty"`
		
		Height *int `json:"height,omitempty"`
		
		Width *int `json:"width,omitempty"`
		*Alias
	}{ 
		Resolution: o.Resolution,
		
		ImageUri: o.ImageUri,
		
		Height: o.Height,
		
		Width: o.Width,
		Alias:    (*Alias)(o),
	})
}

func (o *Documentthumbnail) UnmarshalJSON(b []byte) error {
	var DocumentthumbnailMap map[string]interface{}
	err := json.Unmarshal(b, &DocumentthumbnailMap)
	if err != nil {
		return err
	}
	
	if Resolution, ok := DocumentthumbnailMap["resolution"].(string); ok {
		o.Resolution = &Resolution
	}
    
	if ImageUri, ok := DocumentthumbnailMap["imageUri"].(string); ok {
		o.ImageUri = &ImageUri
	}
    
	if Height, ok := DocumentthumbnailMap["height"].(float64); ok {
		HeightInt := int(Height)
		o.Height = &HeightInt
	}
	
	if Width, ok := DocumentthumbnailMap["width"].(float64); ok {
		WidthInt := int(Width)
		o.Width = &WidthInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Documentthumbnail) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
