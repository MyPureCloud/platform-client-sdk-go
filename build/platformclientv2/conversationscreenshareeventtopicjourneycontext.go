package platformclientv2
import (
	"encoding/json"
)

// Conversationscreenshareeventtopicjourneycontext
type Conversationscreenshareeventtopicjourneycontext struct { 
	// Customer
	Customer *Conversationscreenshareeventtopicjourneycustomer `json:"customer,omitempty"`


	// CustomerSession
	CustomerSession *Conversationscreenshareeventtopicjourneycustomersession `json:"customerSession,omitempty"`


	// TriggeringAction
	TriggeringAction *Conversationscreenshareeventtopicjourneyaction `json:"triggeringAction,omitempty"`

}

// String returns a JSON representation of the model
func (o *Conversationscreenshareeventtopicjourneycontext) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
