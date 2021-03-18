package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Eventcondition
type Eventcondition struct { 
	// Key - The event key.
	Key *string `json:"key,omitempty"`


	// Values - The event values.
	Values *[]string `json:"values,omitempty"`


	// Operator - The comparison operator.
	Operator *string `json:"operator,omitempty"`


	// StreamType - The stream type for which this condition can be satisfied.
	StreamType *string `json:"streamType,omitempty"`


	// SessionType - The session type for which this condition can be satisfied.
	SessionType *string `json:"sessionType,omitempty"`


	// EventName - The name of the event for which this condition can be satisfied.
	EventName *string `json:"eventName,omitempty"`

}

// String returns a JSON representation of the model
func (o *Eventcondition) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
