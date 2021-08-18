package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Edgelogsjobuploadrequest
type Edgelogsjobuploadrequest struct { 
	// FileIds - A list of file ids to upload.
	FileIds *[]string `json:"fileIds,omitempty"`

}

func (u *Edgelogsjobuploadrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Edgelogsjobuploadrequest

	

	return json.Marshal(&struct { 
		FileIds *[]string `json:"fileIds,omitempty"`
		*Alias
	}{ 
		FileIds: u.FileIds,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Edgelogsjobuploadrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
