package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Scrollpercentageeventtrigger - Details about a scroll percentage event trigger
type Scrollpercentageeventtrigger struct { 
	// Percentage - Percentage of a webpage at which an event is triggered.
	Percentage *int `json:"percentage,omitempty"`


	// EventName - Name of event triggered after scrolling to the specified percentage.
	EventName *string `json:"eventName,omitempty"`

}

// String returns a JSON representation of the model
func (o *Scrollpercentageeventtrigger) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
