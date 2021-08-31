package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Locationimage
type Locationimage struct { 
	// Resolution - Height and/or width of image. ex: 640x480 or x128
	Resolution *string `json:"resolution,omitempty"`


	// ImageUri
	ImageUri *string `json:"imageUri,omitempty"`

}

func (o *Locationimage) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Locationimage
	
	return json.Marshal(&struct { 
		Resolution *string `json:"resolution,omitempty"`
		
		ImageUri *string `json:"imageUri,omitempty"`
		*Alias
	}{ 
		Resolution: o.Resolution,
		
		ImageUri: o.ImageUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Locationimage) UnmarshalJSON(b []byte) error {
	var LocationimageMap map[string]interface{}
	err := json.Unmarshal(b, &LocationimageMap)
	if err != nil {
		return err
	}
	
	if Resolution, ok := LocationimageMap["resolution"].(string); ok {
		o.Resolution = &Resolution
	}
	
	if ImageUri, ok := LocationimageMap["imageUri"].(string); ok {
		o.ImageUri = &ImageUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Locationimage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
