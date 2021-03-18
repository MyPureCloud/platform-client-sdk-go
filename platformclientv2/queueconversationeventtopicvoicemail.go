package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationeventtopicvoicemail
type Queueconversationeventtopicvoicemail struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// UploadStatus
	UploadStatus *string `json:"uploadStatus,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queueconversationeventtopicvoicemail) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
