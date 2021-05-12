package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Edgelogsjobuploadrequest
type Edgelogsjobuploadrequest struct { 
	// FileIds - A list of file ids to upload.
	FileIds *[]string `json:"fileIds,omitempty"`

}

// String returns a JSON representation of the model
func (o *Edgelogsjobuploadrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
