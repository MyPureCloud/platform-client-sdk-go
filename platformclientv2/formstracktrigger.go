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

func (u *Formstracktrigger) MarshalJSON() ([]byte, error) {
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
		Selector: u.Selector,
		
		FormName: u.FormName,
		
		CaptureDataOnFormAbandon: u.CaptureDataOnFormAbandon,
		
		CaptureDataOnFormSubmit: u.CaptureDataOnFormSubmit,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Formstracktrigger) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
