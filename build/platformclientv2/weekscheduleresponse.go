package platformclientv2
import (
	"encoding/json"
)

// Weekscheduleresponse - Response for query for week schedule for a given week in management unit
type Weekscheduleresponse struct { 
	// Result - The result of the request. The value will be null if response is large
	Result *Weekschedule `json:"result,omitempty"`


	// DownloadUrl - The url to fetch the result for large responses. The value is null if result contains the data
	DownloadUrl *string `json:"downloadUrl,omitempty"`

}

// String returns a JSON representation of the model
func (o *Weekscheduleresponse) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
