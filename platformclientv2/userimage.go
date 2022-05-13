package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Userimage
type Userimage struct { 
	// Resolution - Height and/or width of image. ex: 640x480 or x128
	Resolution *string `json:"resolution,omitempty"`


	// ImageUri
	ImageUri *string `json:"imageUri,omitempty"`

}

func (o *Userimage) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Userimage
	
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

func (o *Userimage) UnmarshalJSON(b []byte) error {
	var UserimageMap map[string]interface{}
	err := json.Unmarshal(b, &UserimageMap)
	if err != nil {
		return err
	}
	
	if Resolution, ok := UserimageMap["resolution"].(string); ok {
		o.Resolution = &Resolution
	}
    
	if ImageUri, ok := UserimageMap["imageUri"].(string); ok {
		o.ImageUri = &ImageUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Userimage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
