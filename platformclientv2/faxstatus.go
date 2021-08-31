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

func (o *Faxstatus) MarshalJSON() ([]byte, error) {
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
		Direction: o.Direction,
		
		ExpectedPages: o.ExpectedPages,
		
		ActivePage: o.ActivePage,
		
		LinesTransmitted: o.LinesTransmitted,
		
		BytesTransmitted: o.BytesTransmitted,
		
		BaudRate: o.BaudRate,
		
		PageErrors: o.PageErrors,
		
		LineErrors: o.LineErrors,
		Alias:    (*Alias)(o),
	})
}

func (o *Faxstatus) UnmarshalJSON(b []byte) error {
	var FaxstatusMap map[string]interface{}
	err := json.Unmarshal(b, &FaxstatusMap)
	if err != nil {
		return err
	}
	
	if Direction, ok := FaxstatusMap["direction"].(string); ok {
		o.Direction = &Direction
	}
	
	if ExpectedPages, ok := FaxstatusMap["expectedPages"].(float64); ok {
		ExpectedPagesInt := int(ExpectedPages)
		o.ExpectedPages = &ExpectedPagesInt
	}
	
	if ActivePage, ok := FaxstatusMap["activePage"].(float64); ok {
		ActivePageInt := int(ActivePage)
		o.ActivePage = &ActivePageInt
	}
	
	if LinesTransmitted, ok := FaxstatusMap["linesTransmitted"].(float64); ok {
		LinesTransmittedInt := int(LinesTransmitted)
		o.LinesTransmitted = &LinesTransmittedInt
	}
	
	if BytesTransmitted, ok := FaxstatusMap["bytesTransmitted"].(float64); ok {
		BytesTransmittedInt := int(BytesTransmitted)
		o.BytesTransmitted = &BytesTransmittedInt
	}
	
	if BaudRate, ok := FaxstatusMap["baudRate"].(float64); ok {
		BaudRateInt := int(BaudRate)
		o.BaudRate = &BaudRateInt
	}
	
	if PageErrors, ok := FaxstatusMap["pageErrors"].(float64); ok {
		PageErrorsInt := int(PageErrors)
		o.PageErrors = &PageErrorsInt
	}
	
	if LineErrors, ok := FaxstatusMap["lineErrors"].(float64); ok {
		LineErrorsInt := int(LineErrors)
		o.LineErrors = &LineErrorsInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Faxstatus) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
