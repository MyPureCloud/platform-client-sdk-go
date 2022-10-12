package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Eventerror
type Eventerror struct { 
	// EventId - The eventId (V4 UUID) for the event that encountered an error.
	EventId *string `json:"eventId,omitempty"`


	// Message - A message describing the error.
	Message *string `json:"message,omitempty"`


	// Retryable - The event for this eventId can be resubmitted if this value is true.
	Retryable *bool `json:"retryable,omitempty"`

}

func (o *Eventerror) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Eventerror
	
	return json.Marshal(&struct { 
		EventId *string `json:"eventId,omitempty"`
		
		Message *string `json:"message,omitempty"`
		
		Retryable *bool `json:"retryable,omitempty"`
		*Alias
	}{ 
		EventId: o.EventId,
		
		Message: o.Message,
		
		Retryable: o.Retryable,
		Alias:    (*Alias)(o),
	})
}

func (o *Eventerror) UnmarshalJSON(b []byte) error {
	var EventerrorMap map[string]interface{}
	err := json.Unmarshal(b, &EventerrorMap)
	if err != nil {
		return err
	}
	
	if EventId, ok := EventerrorMap["eventId"].(string); ok {
		o.EventId = &EventId
	}
    
	if Message, ok := EventerrorMap["message"].(string); ok {
		o.Message = &Message
	}
    
	if Retryable, ok := EventerrorMap["retryable"].(bool); ok {
		o.Retryable = &Retryable
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Eventerror) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
