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

func (o *Edgelogsjobuploadrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Edgelogsjobuploadrequest
	
	return json.Marshal(&struct { 
		FileIds *[]string `json:"fileIds,omitempty"`
		*Alias
	}{ 
		FileIds: o.FileIds,
		Alias:    (*Alias)(o),
	})
}

func (o *Edgelogsjobuploadrequest) UnmarshalJSON(b []byte) error {
	var EdgelogsjobuploadrequestMap map[string]interface{}
	err := json.Unmarshal(b, &EdgelogsjobuploadrequestMap)
	if err != nil {
		return err
	}
	
	if FileIds, ok := EdgelogsjobuploadrequestMap["fileIds"].([]interface{}); ok {
		FileIdsString, _ := json.Marshal(FileIds)
		json.Unmarshal(FileIdsString, &o.FileIds)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Edgelogsjobuploadrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
