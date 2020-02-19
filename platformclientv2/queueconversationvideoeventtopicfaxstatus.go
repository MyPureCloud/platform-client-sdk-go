package platformclientv2
import (
	"encoding/json"
)

// Queueconversationvideoeventtopicfaxstatus
type Queueconversationvideoeventtopicfaxstatus struct { 
	// Direction
	Direction *string `json:"direction,omitempty"`


	// ExpectedPages
	ExpectedPages *int32 `json:"expectedPages,omitempty"`


	// ActivePage
	ActivePage *int32 `json:"activePage,omitempty"`


	// LinesTransmitted
	LinesTransmitted *int32 `json:"linesTransmitted,omitempty"`


	// BytesTransmitted
	BytesTransmitted *int32 `json:"bytesTransmitted,omitempty"`


	// BaudRate
	BaudRate *int32 `json:"baudRate,omitempty"`


	// PageErrors
	PageErrors *int32 `json:"pageErrors,omitempty"`


	// LineErrors
	LineErrors *int32 `json:"lineErrors,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queueconversationvideoeventtopicfaxstatus) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
