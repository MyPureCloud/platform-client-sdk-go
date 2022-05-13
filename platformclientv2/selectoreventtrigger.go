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

func (o *Selectoreventtrigger) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Selectoreventtrigger
	
	return json.Marshal(&struct { 
		Selector *string `json:"selector,omitempty"`
		
		EventName *string `json:"eventName,omitempty"`
		*Alias
	}{ 
		Selector: o.Selector,
		
		EventName: o.EventName,
		Alias:    (*Alias)(o),
	})
}

func (o *Selectoreventtrigger) UnmarshalJSON(b []byte) error {
	var SelectoreventtriggerMap map[string]interface{}
	err := json.Unmarshal(b, &SelectoreventtriggerMap)
	if err != nil {
		return err
	}
	
	if Selector, ok := SelectoreventtriggerMap["selector"].(string); ok {
		o.Selector = &Selector
	}
    
	if EventName, ok := SelectoreventtriggerMap["eventName"].(string); ok {
		o.EventName = &EventName
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Selectoreventtrigger) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
