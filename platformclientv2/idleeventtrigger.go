package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Idleeventtrigger - Details about an idle event trigger
type Idleeventtrigger struct { 
	// EventName - Name of event triggered after period of inactivity.
	EventName *string `json:"eventName,omitempty"`


	// IdleAfterSeconds - Number of seconds of inactivity before an event is triggered.
	IdleAfterSeconds *int `json:"idleAfterSeconds,omitempty"`

}

// String returns a JSON representation of the model
func (o *Idleeventtrigger) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
