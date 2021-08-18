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

func (u *Locationimage) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Locationimage

	

	return json.Marshal(&struct { 
		Resolution *string `json:"resolution,omitempty"`
		
		ImageUri *string `json:"imageUri,omitempty"`
		*Alias
	}{ 
		Resolution: u.Resolution,
		
		ImageUri: u.ImageUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Locationimage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
