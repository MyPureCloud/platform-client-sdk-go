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

func (u *Batchdownloadjobsubmission) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Batchdownloadjobsubmission

	

	return json.Marshal(&struct { 
		BatchDownloadRequestList *[]Batchdownloadrequest `json:"batchDownloadRequestList,omitempty"`
		*Alias
	}{ 
		BatchDownloadRequestList: u.BatchDownloadRequestList,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Batchdownloadjobsubmission) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
