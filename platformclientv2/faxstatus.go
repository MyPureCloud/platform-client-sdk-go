package platformclientv2
import (
	"encoding/json"
)

// Faxstatus
type Faxstatus struct { 
	// Direction - The fax direction, either \"send\" or \"receive\".
	Direction *string `json:"direction,omitempty"`


	// ExpectedPages - Total number of expected pages, if known.
	ExpectedPages *int64 `json:"expectedPages,omitempty"`


	// ActivePage - Active page of the transmission.
	ActivePage *int64 `json:"activePage,omitempty"`


	// LinesTransmitted - Number of lines that have completed transmission.
	LinesTransmitted *int64 `json:"linesTransmitted,omitempty"`


	// BytesTransmitted - Number of bytes that have competed transmission.
	BytesTransmitted *int64 `json:"bytesTransmitted,omitempty"`


	// BaudRate - Current signaling rate of transmission, baud rate.
	BaudRate *int64 `json:"baudRate,omitempty"`


	// PageErrors - Number of page errors.
	PageErrors *int64 `json:"pageErrors,omitempty"`


	// LineErrors - Number of line errors.
	LineErrors *int64 `json:"lineErrors,omitempty"`

}

// String returns a JSON representation of the model
func (o *Faxstatus) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
