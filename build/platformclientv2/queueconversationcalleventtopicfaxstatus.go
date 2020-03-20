package platformclientv2
import (
	"encoding/json"
)

// Queueconversationcalleventtopicfaxstatus
type Queueconversationcalleventtopicfaxstatus struct { 
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


	// DataRate
	DataRate *int32 `json:"dataRate,omitempty"`


	// PageErrors
	PageErrors *int32 `json:"pageErrors,omitempty"`


	// LineErrors
	LineErrors *int32 `json:"lineErrors,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queueconversationcalleventtopicfaxstatus) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
