package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Faxstatus
type Faxstatus struct { 
	// Direction - The fax direction, either \"send\" or \"receive\".
	Direction *string `json:"direction,omitempty"`


	// ExpectedPages - Total number of expected pages, if known.
	ExpectedPages *int `json:"expectedPages,omitempty"`


	// ActivePage - Active page of the transmission.
	ActivePage *int `json:"activePage,omitempty"`


	// LinesTransmitted - Number of lines that have completed transmission.
	LinesTransmitted *int `json:"linesTransmitted,omitempty"`


	// BytesTransmitted - Number of bytes that have competed transmission.
	BytesTransmitted *int `json:"bytesTransmitted,omitempty"`


	// BaudRate - Current signaling rate of transmission, baud rate.
	BaudRate *int `json:"baudRate,omitempty"`


	// PageErrors - Number of page errors.
	PageErrors *int `json:"pageErrors,omitempty"`


	// LineErrors - Number of line errors.
	LineErrors *int `json:"lineErrors,omitempty"`

}

func (u *Faxstatus) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Faxstatus

	

	return json.Marshal(&struct { 
		Direction *string `json:"direction,omitempty"`
		
		ExpectedPages *int `json:"expectedPages,omitempty"`
		
		ActivePage *int `json:"activePage,omitempty"`
		
		LinesTransmitted *int `json:"linesTransmitted,omitempty"`
		
		BytesTransmitted *int `json:"bytesTransmitted,omitempty"`
		
		BaudRate *int `json:"baudRate,omitempty"`
		
		PageErrors *int `json:"pageErrors,omitempty"`
		
		LineErrors *int `json:"lineErrors,omitempty"`
		*Alias
	}{ 
		Direction: u.Direction,
		
		ExpectedPages: u.ExpectedPages,
		
		ActivePage: u.ActivePage,
		
		LinesTransmitted: u.LinesTransmitted,
		
		BytesTransmitted: u.BytesTransmitted,
		
		BaudRate: u.BaudRate,
		
		PageErrors: u.PageErrors,
		
		LineErrors: u.LineErrors,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Faxstatus) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
