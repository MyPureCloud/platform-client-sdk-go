package platformclientv2
import (
	"encoding/json"
)

// Generateweekscheduleresponse - Response for query for week schedule for a given week in management unit
type Generateweekscheduleresponse struct { 
	// DownloadUrl - The url to fetch the result for large responses. The value is null if result contains the data
	DownloadUrl *string `json:"downloadUrl,omitempty"`


	// Status - The status of the request
	Status *string `json:"status,omitempty"`


	// OperationId - The operation id to watch for on the notification topic if status == Processing
	OperationId *string `json:"operationId,omitempty"`

}

// String returns a JSON representation of the model
func (o *Generateweekscheduleresponse) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
