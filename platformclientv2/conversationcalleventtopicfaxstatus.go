package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationcalleventtopicfaxstatus
type Conversationcalleventtopicfaxstatus struct { 
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

func (o *Conversationcalleventtopicfaxstatus) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationcalleventtopicfaxstatus
	
	return json.Marshal(&struct { 
		Direction *string `json:"direction,omitempty"`
		
		ExpectedPages *int `json:"expectedPages,omitempty"`
		
		ActivePage *int `json:"activePage,omitempty"`
		
		LinesTransmitted *int `json:"linesTransmitted,omitempty"`
		
		BytesTransmitted *int `json:"bytesTransmitted,omitempty"`
		
		DataRate *int `json:"dataRate,omitempty"`
		
		PageErrors *int `json:"pageErrors,omitempty"`
		
		LineErrors *int `json:"lineErrors,omitempty"`
		*Alias
	}{ 
		Direction: o.Direction,
		
		ExpectedPages: o.ExpectedPages,
		
		ActivePage: o.ActivePage,
		
		LinesTransmitted: o.LinesTransmitted,
		
		BytesTransmitted: o.BytesTransmitted,
		
		DataRate: o.DataRate,
		
		PageErrors: o.PageErrors,
		
		LineErrors: o.LineErrors,
		Alias:    (*Alias)(o),
	})
}

func (o *Conversationcalleventtopicfaxstatus) UnmarshalJSON(b []byte) error {
	var ConversationcalleventtopicfaxstatusMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationcalleventtopicfaxstatusMap)
	if err != nil {
		return err
	}
	
	if Direction, ok := ConversationcalleventtopicfaxstatusMap["direction"].(string); ok {
		o.Direction = &Direction
	}
	
	if ExpectedPages, ok := ConversationcalleventtopicfaxstatusMap["expectedPages"].(float64); ok {
		ExpectedPagesInt := int(ExpectedPages)
		o.ExpectedPages = &ExpectedPagesInt
	}
	
	if ActivePage, ok := ConversationcalleventtopicfaxstatusMap["activePage"].(float64); ok {
		ActivePageInt := int(ActivePage)
		o.ActivePage = &ActivePageInt
	}
	
	if LinesTransmitted, ok := ConversationcalleventtopicfaxstatusMap["linesTransmitted"].(float64); ok {
		LinesTransmittedInt := int(LinesTransmitted)
		o.LinesTransmitted = &LinesTransmittedInt
	}
	
	if BytesTransmitted, ok := ConversationcalleventtopicfaxstatusMap["bytesTransmitted"].(float64); ok {
		BytesTransmittedInt := int(BytesTransmitted)
		o.BytesTransmitted = &BytesTransmittedInt
	}
	
	if DataRate, ok := ConversationcalleventtopicfaxstatusMap["dataRate"].(float64); ok {
		DataRateInt := int(DataRate)
		o.DataRate = &DataRateInt
	}
	
	if PageErrors, ok := ConversationcalleventtopicfaxstatusMap["pageErrors"].(float64); ok {
		PageErrorsInt := int(PageErrors)
		o.PageErrors = &PageErrorsInt
	}
	
	if LineErrors, ok := ConversationcalleventtopicfaxstatusMap["lineErrors"].(float64); ok {
		LineErrorsInt := int(LineErrors)
		o.LineErrors = &LineErrorsInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationcalleventtopicfaxstatus) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
