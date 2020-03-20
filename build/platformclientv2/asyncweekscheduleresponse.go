package platformclientv2
import (
	"encoding/json"
)

// Asyncweekscheduleresponse - Response for query for week schedule for a given week in management unit
type Asyncweekscheduleresponse struct { 
	// Result - Week schedule result. The value will be null if the data is sent through notification or if response is large.
	Result *Weekschedule `json:"result,omitempty"`


	// DownloadUrl - The url to fetch the result for large responses. The value is null if result contains the data
	DownloadUrl *string `json:"downloadUrl,omitempty"`


	// Status - The status of the request
	Status *string `json:"status,omitempty"`


	// OperationId - The operation id to watch for on the notification topic if status == Processing
	OperationId *string `json:"operationId,omitempty"`

}

// String returns a JSON representation of the model
func (o *Asyncweekscheduleresponse) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
