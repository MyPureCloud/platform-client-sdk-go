package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Fileuploadmode
type Fileuploadmode struct { 
	// FileTypes - A list of supported content types for uploading files
	FileTypes *[]string `json:"fileTypes,omitempty"`


	// MaxFileSizeKB - The maximum file size for file uploads in kilobytes. Default is 10240 (10 MB)
	MaxFileSizeKB *int `json:"maxFileSizeKB,omitempty"`

}

func (u *Fileuploadmode) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Fileuploadmode

	

	return json.Marshal(&struct { 
		FileTypes *[]string `json:"fileTypes,omitempty"`
		
		MaxFileSizeKB *int `json:"maxFileSizeKB,omitempty"`
		*Alias
	}{ 
		FileTypes: u.FileTypes,
		
		MaxFileSizeKB: u.MaxFileSizeKB,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Fileuploadmode) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
