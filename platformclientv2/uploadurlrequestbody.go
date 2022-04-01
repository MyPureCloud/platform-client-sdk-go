package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Uploadurlrequestbody
type Uploadurlrequestbody struct { 
	// ContentLengthBytes - The expected content length (in bytes) of the gzip-encoded data that will be PUT to the returned signed URL
	ContentLengthBytes *int `json:"contentLengthBytes,omitempty"`

}

func (o *Uploadurlrequestbody) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Uploadurlrequestbody
	
	return json.Marshal(&struct { 
		ContentLengthBytes *int `json:"contentLengthBytes,omitempty"`
		*Alias
	}{ 
		ContentLengthBytes: o.ContentLengthBytes,
		Alias:    (*Alias)(o),
	})
}

func (o *Uploadurlrequestbody) UnmarshalJSON(b []byte) error {
	var UploadurlrequestbodyMap map[string]interface{}
	err := json.Unmarshal(b, &UploadurlrequestbodyMap)
	if err != nil {
		return err
	}
	
	if ContentLengthBytes, ok := UploadurlrequestbodyMap["contentLengthBytes"].(float64); ok {
		ContentLengthBytesInt := int(ContentLengthBytes)
		o.ContentLengthBytes = &ContentLengthBytesInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Uploadurlrequestbody) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
