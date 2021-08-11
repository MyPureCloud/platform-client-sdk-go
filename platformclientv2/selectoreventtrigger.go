package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Selectoreventtrigger) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
