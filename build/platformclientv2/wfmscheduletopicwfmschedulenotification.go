package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmscheduletopicwfmschedulenotification
type Wfmscheduletopicwfmschedulenotification struct { 
	// Status
	Status *string `json:"status,omitempty"`


	// OperationId
	OperationId *string `json:"operationId,omitempty"`


	// DownloadUrl
	DownloadUrl *string `json:"downloadUrl,omitempty"`


	// PercentComplete
	PercentComplete *int `json:"percentComplete,omitempty"`


	// EventType
	EventType *string `json:"eventType,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wfmscheduletopicwfmschedulenotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
