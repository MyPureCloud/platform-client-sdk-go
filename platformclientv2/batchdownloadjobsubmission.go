package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Batchdownloadjobsubmission
type Batchdownloadjobsubmission struct { 
	// BatchDownloadRequestList - List of up to 100 items requested
	BatchDownloadRequestList *[]Batchdownloadrequest `json:"batchDownloadRequestList,omitempty"`

}

func (o *Batchdownloadjobsubmission) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Batchdownloadjobsubmission
	
	return json.Marshal(&struct { 
		BatchDownloadRequestList *[]Batchdownloadrequest `json:"batchDownloadRequestList,omitempty"`
		*Alias
	}{ 
		BatchDownloadRequestList: o.BatchDownloadRequestList,
		Alias:    (*Alias)(o),
	})
}

func (o *Batchdownloadjobsubmission) UnmarshalJSON(b []byte) error {
	var BatchdownloadjobsubmissionMap map[string]interface{}
	err := json.Unmarshal(b, &BatchdownloadjobsubmissionMap)
	if err != nil {
		return err
	}
	
	if BatchDownloadRequestList, ok := BatchdownloadjobsubmissionMap["batchDownloadRequestList"].([]interface{}); ok {
		BatchDownloadRequestListString, _ := json.Marshal(BatchDownloadRequestList)
		json.Unmarshal(BatchDownloadRequestListString, &o.BatchDownloadRequestList)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Batchdownloadjobsubmission) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
