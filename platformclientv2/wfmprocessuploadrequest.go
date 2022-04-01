package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmprocessuploadrequest
type Wfmprocessuploadrequest struct { 
	// UploadKey - The uploadKey provided by the request to get an upload URL
	UploadKey *string `json:"uploadKey,omitempty"`

}

func (o *Wfmprocessuploadrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmprocessuploadrequest
	
	return json.Marshal(&struct { 
		UploadKey *string `json:"uploadKey,omitempty"`
		*Alias
	}{ 
		UploadKey: o.UploadKey,
		Alias:    (*Alias)(o),
	})
}

func (o *Wfmprocessuploadrequest) UnmarshalJSON(b []byte) error {
	var WfmprocessuploadrequestMap map[string]interface{}
	err := json.Unmarshal(b, &WfmprocessuploadrequestMap)
	if err != nil {
		return err
	}
	
	if UploadKey, ok := WfmprocessuploadrequestMap["uploadKey"].(string); ok {
		o.UploadKey = &UploadKey
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmprocessuploadrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
