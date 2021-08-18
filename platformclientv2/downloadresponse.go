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

func (u *Downloadresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Downloadresponse

	

	return json.Marshal(&struct { 
		ContentLocationUri *string `json:"contentLocationUri,omitempty"`
		
		ImageUri *string `json:"imageUri,omitempty"`
		
		Thumbnails *[]Documentthumbnail `json:"thumbnails,omitempty"`
		*Alias
	}{ 
		ContentLocationUri: u.ContentLocationUri,
		
		ImageUri: u.ImageUri,
		
		Thumbnails: u.Thumbnails,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Downloadresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
