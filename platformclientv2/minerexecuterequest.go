package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Minerexecuterequest
type Minerexecuterequest struct { 
	// DateStart - Start date for the date range to mine. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	DateStart *time.Time `json:"dateStart,omitempty"`


	// DateEnd - End date for the date range to mine. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	DateEnd *time.Time `json:"dateEnd,omitempty"`


	// UploadKey - Location of input conversations.
	UploadKey *string `json:"uploadKey,omitempty"`


	// MediaType - Media type for filtering conversations.
	MediaType *string `json:"mediaType,omitempty"`


	// QueueIds - List of queue IDs for filtering conversations.
	QueueIds *[]string `json:"queueIds,omitempty"`

}

// String returns a JSON representation of the model
func (o *Minerexecuterequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
