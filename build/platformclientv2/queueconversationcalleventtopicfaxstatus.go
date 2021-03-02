package platformclientv2
import (
	"encoding/json"
)

// Queueconversationcalleventtopicfaxstatus
type Queueconversationcalleventtopicfaxstatus struct { 
	// Direction
	Direction *string `json:"direction,omitempty"`


	// ExpectedPages
	ExpectedPages *int `json:"expectedPages,omitempty"`


	// ActivePage
	ActivePage *int `json:"activePage,omitempty"`


	// LinesTransmitted
	LinesTransmitted *int `json:"linesTransmitted,omitempty"`


	// BytesTransmitted
	BytesTransmitted *int `json:"bytesTransmitted,omitempty"`


	// DataRate
	DataRate *int `json:"dataRate,omitempty"`


	// PageErrors
	PageErrors *int `json:"pageErrors,omitempty"`


	// LineErrors
	LineErrors *int `json:"lineErrors,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queueconversationcalleventtopicfaxstatus) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
