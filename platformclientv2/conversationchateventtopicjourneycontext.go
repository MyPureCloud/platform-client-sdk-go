package platformclientv2
import (
	"encoding/json"
)

// Conversationchateventtopicjourneycontext
type Conversationchateventtopicjourneycontext struct { 
	// Customer
	Customer *Conversationchateventtopicjourneycustomer `json:"customer,omitempty"`


	// CustomerSession
	CustomerSession *Conversationchateventtopicjourneycustomersession `json:"customerSession,omitempty"`


	// TriggeringAction
	TriggeringAction *Conversationchateventtopicjourneyaction `json:"triggeringAction,omitempty"`

}

// String returns a JSON representation of the model
func (o *Conversationchateventtopicjourneycontext) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
