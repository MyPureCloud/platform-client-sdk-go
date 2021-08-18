package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Uploadurlresponse
type Uploadurlresponse struct { 
	// Url - Presigned URL to PUT the file to
	Url *string `json:"url,omitempty"`


	// UploadKey - Key that identifies the file in the storage including the file name
	UploadKey *string `json:"uploadKey,omitempty"`


	// Headers - Required headers when uploading a file through PUT request to the URL
	Headers *map[string]string `json:"headers,omitempty"`

}

func (u *Uploadurlresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Uploadurlresponse

	

	return json.Marshal(&struct { 
		Url *string `json:"url,omitempty"`
		
		UploadKey *string `json:"uploadKey,omitempty"`
		
		Headers *map[string]string `json:"headers,omitempty"`
		*Alias
	}{ 
		Url: u.Url,
		
		UploadKey: u.UploadKey,
		
		Headers: u.Headers,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Uploadurlresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
