package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Downloadresponse
type Downloadresponse struct { 
	// ContentLocationUri
	ContentLocationUri *string `json:"contentLocationUri,omitempty"`


	// ImageUri
	ImageUri *string `json:"imageUri,omitempty"`


	// Thumbnails
	Thumbnails *[]Documentthumbnail `json:"thumbnails,omitempty"`

}

func (o *Downloadresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Downloadresponse
	
	return json.Marshal(&struct { 
		ContentLocationUri *string `json:"contentLocationUri,omitempty"`
		
		ImageUri *string `json:"imageUri,omitempty"`
		
		Thumbnails *[]Documentthumbnail `json:"thumbnails,omitempty"`
		*Alias
	}{ 
		ContentLocationUri: o.ContentLocationUri,
		
		ImageUri: o.ImageUri,
		
		Thumbnails: o.Thumbnails,
		Alias:    (*Alias)(o),
	})
}

func (o *Downloadresponse) UnmarshalJSON(b []byte) error {
	var DownloadresponseMap map[string]interface{}
	err := json.Unmarshal(b, &DownloadresponseMap)
	if err != nil {
		return err
	}
	
	if ContentLocationUri, ok := DownloadresponseMap["contentLocationUri"].(string); ok {
		o.ContentLocationUri = &ContentLocationUri
	}
	
	if ImageUri, ok := DownloadresponseMap["imageUri"].(string); ok {
		o.ImageUri = &ImageUri
	}
	
	if Thumbnails, ok := DownloadresponseMap["thumbnails"].([]interface{}); ok {
		ThumbnailsString, _ := json.Marshal(Thumbnails)
		json.Unmarshal(ThumbnailsString, &o.Thumbnails)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Downloadresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
