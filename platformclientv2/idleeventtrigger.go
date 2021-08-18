package platformclientv2
import (
	"github.com/leekchan/timeutil"
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

func (u *Idleeventtrigger) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Idleeventtrigger

	

	return json.Marshal(&struct { 
		EventName *string `json:"eventName,omitempty"`
		
		IdleAfterSeconds *int `json:"idleAfterSeconds,omitempty"`
		*Alias
	}{ 
		EventName: u.EventName,
		
		IdleAfterSeconds: u.IdleAfterSeconds,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Idleeventtrigger) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
