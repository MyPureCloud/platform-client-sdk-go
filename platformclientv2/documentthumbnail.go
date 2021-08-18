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

func (u *Documentthumbnail) MarshalJSON() ([]byte, error) {
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
		Resolution: u.Resolution,
		
		ImageUri: u.ImageUri,
		
		Height: u.Height,
		
		Width: u.Width,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Documentthumbnail) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
