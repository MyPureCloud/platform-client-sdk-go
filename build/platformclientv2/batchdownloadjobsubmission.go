package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Batchdownloadjobsubmission
type Batchdownloadjobsubmission struct { 
	// BatchDownloadRequestList - List of up to 100 items requested
	BatchDownloadRequestList *[]Batchdownloadrequest `json:"batchDownloadRequestList,omitempty"`

}

// String returns a JSON representation of the model
func (o *Batchdownloadjobsubmission) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
