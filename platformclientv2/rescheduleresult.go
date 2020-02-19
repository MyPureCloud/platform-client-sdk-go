package platformclientv2
import (
	"encoding/json"
)

// Rescheduleresult
type Rescheduleresult struct { 
	// DownloadUrl - The url from which to download the resulting WeekSchedule object for the rescheduling range
	DownloadUrl *string `json:"downloadUrl,omitempty"`

}

// String returns a JSON representation of the model
func (o *Rescheduleresult) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
