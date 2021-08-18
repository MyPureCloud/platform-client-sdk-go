package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationsocialexpressioneventtopicfaxstatus
type Queueconversationsocialexpressioneventtopicfaxstatus struct { 
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


	// BaudRate
	BaudRate *int `json:"baudRate,omitempty"`


	// PageErrors
	PageErrors *int `json:"pageErrors,omitempty"`


	// LineErrors
	LineErrors *int `json:"lineErrors,omitempty"`

}

func (u *Queueconversationsocialexpressioneventtopicfaxstatus) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queueconversationsocialexpressioneventtopicfaxstatus

	

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
func (o *Queueconversationsocialexpressioneventtopicfaxstatus) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
