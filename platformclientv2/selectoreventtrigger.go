package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Selectoreventtrigger - Details about a selector event trigger
type Selectoreventtrigger struct { 
	// Selector - Element that triggers event.
	Selector *string `json:"selector,omitempty"`


	// EventName - Name of event triggered when element matching selector is interacted with.
	EventName *string `json:"eventName,omitempty"`

}

func (u *Selectoreventtrigger) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Selectoreventtrigger

	

	return json.Marshal(&struct { 
		Selector *string `json:"selector,omitempty"`
		
		EventName *string `json:"eventName,omitempty"`
		*Alias
	}{ 
		Selector: u.Selector,
		
		EventName: u.EventName,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Selectoreventtrigger) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
