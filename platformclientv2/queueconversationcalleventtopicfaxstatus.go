package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
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

func (o *Queueconversationcalleventtopicfaxstatus) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queueconversationcalleventtopicfaxstatus
	
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

func (o *Queueconversationcalleventtopicfaxstatus) UnmarshalJSON(b []byte) error {
	var QueueconversationcalleventtopicfaxstatusMap map[string]interface{}
	err := json.Unmarshal(b, &QueueconversationcalleventtopicfaxstatusMap)
	if err != nil {
		return err
	}
	
	if Direction, ok := QueueconversationcalleventtopicfaxstatusMap["direction"].(string); ok {
		o.Direction = &Direction
	}
	
	if ExpectedPages, ok := QueueconversationcalleventtopicfaxstatusMap["expectedPages"].(float64); ok {
		ExpectedPagesInt := int(ExpectedPages)
		o.ExpectedPages = &ExpectedPagesInt
	}
	
	if ActivePage, ok := QueueconversationcalleventtopicfaxstatusMap["activePage"].(float64); ok {
		ActivePageInt := int(ActivePage)
		o.ActivePage = &ActivePageInt
	}
	
	if LinesTransmitted, ok := QueueconversationcalleventtopicfaxstatusMap["linesTransmitted"].(float64); ok {
		LinesTransmittedInt := int(LinesTransmitted)
		o.LinesTransmitted = &LinesTransmittedInt
	}
	
	if BytesTransmitted, ok := QueueconversationcalleventtopicfaxstatusMap["bytesTransmitted"].(float64); ok {
		BytesTransmittedInt := int(BytesTransmitted)
		o.BytesTransmitted = &BytesTransmittedInt
	}
	
	if DataRate, ok := QueueconversationcalleventtopicfaxstatusMap["dataRate"].(float64); ok {
		DataRateInt := int(DataRate)
		o.DataRate = &DataRateInt
	}
	
	if PageErrors, ok := QueueconversationcalleventtopicfaxstatusMap["pageErrors"].(float64); ok {
		PageErrorsInt := int(PageErrors)
		o.PageErrors = &PageErrorsInt
	}
	
	if LineErrors, ok := QueueconversationcalleventtopicfaxstatusMap["lineErrors"].(float64); ok {
		LineErrorsInt := int(LineErrors)
		o.LineErrors = &LineErrorsInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Queueconversationcalleventtopicfaxstatus) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
