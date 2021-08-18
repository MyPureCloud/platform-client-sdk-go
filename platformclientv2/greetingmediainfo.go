package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Greetingmediainfo
type Greetingmediainfo struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// MediaFileUri
	MediaFileUri *string `json:"mediaFileUri,omitempty"`


	// MediaImageUri
	MediaImageUri *string `json:"mediaImageUri,omitempty"`

}

func (u *Greetingmediainfo) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Greetingmediainfo

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		MediaFileUri *string `json:"mediaFileUri,omitempty"`
		
		MediaImageUri *string `json:"mediaImageUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		MediaFileUri: u.MediaFileUri,
		
		MediaImageUri: u.MediaImageUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Greetingmediainfo) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
