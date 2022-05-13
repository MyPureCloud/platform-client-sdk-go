package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Formstracktrigger - Details about a forms tracking event trigger
type Formstracktrigger struct { 
	// Selector - Form element that triggers the form submitted or abandoned event.
	Selector *string `json:"selector,omitempty"`


	// FormName - Prefix for the form submitted or abandoned event name.
	FormName *string `json:"formName,omitempty"`


	// CaptureDataOnFormAbandon - Whether to capture the form data in the form abandoned event.
	CaptureDataOnFormAbandon *bool `json:"captureDataOnFormAbandon,omitempty"`


	// CaptureDataOnFormSubmit - Whether to capture the form data in the form submitted event.
	CaptureDataOnFormSubmit *bool `json:"captureDataOnFormSubmit,omitempty"`

}

func (o *Formstracktrigger) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Formstracktrigger
	
	return json.Marshal(&struct { 
		Selector *string `json:"selector,omitempty"`
		
		FormName *string `json:"formName,omitempty"`
		
		CaptureDataOnFormAbandon *bool `json:"captureDataOnFormAbandon,omitempty"`
		
		CaptureDataOnFormSubmit *bool `json:"captureDataOnFormSubmit,omitempty"`
		*Alias
	}{ 
		Selector: o.Selector,
		
		FormName: o.FormName,
		
		CaptureDataOnFormAbandon: o.CaptureDataOnFormAbandon,
		
		CaptureDataOnFormSubmit: o.CaptureDataOnFormSubmit,
		Alias:    (*Alias)(o),
	})
}

func (o *Formstracktrigger) UnmarshalJSON(b []byte) error {
	var FormstracktriggerMap map[string]interface{}
	err := json.Unmarshal(b, &FormstracktriggerMap)
	if err != nil {
		return err
	}
	
	if Selector, ok := FormstracktriggerMap["selector"].(string); ok {
		o.Selector = &Selector
	}
    
	if FormName, ok := FormstracktriggerMap["formName"].(string); ok {
		o.FormName = &FormName
	}
    
	if CaptureDataOnFormAbandon, ok := FormstracktriggerMap["captureDataOnFormAbandon"].(bool); ok {
		o.CaptureDataOnFormAbandon = &CaptureDataOnFormAbandon
	}
    
	if CaptureDataOnFormSubmit, ok := FormstracktriggerMap["captureDataOnFormSubmit"].(bool); ok {
		o.CaptureDataOnFormSubmit = &CaptureDataOnFormSubmit
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Formstracktrigger) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
