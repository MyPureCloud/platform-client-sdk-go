package platformclientv2
import (
	"encoding/json"
)

// Queueconversationscreenshareeventtopicjourneycontext
type Queueconversationscreenshareeventtopicjourneycontext struct { 
	// Customer
	Customer *Queueconversationscreenshareeventtopicjourneycustomer `json:"customer,omitempty"`


	// CustomerSession
	CustomerSession *Queueconversationscreenshareeventtopicjourneycustomersession `json:"customerSession,omitempty"`


	// TriggeringAction
	TriggeringAction *Queueconversationscreenshareeventtopicjourneyaction `json:"triggeringAction,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queueconversationscreenshareeventtopicjourneycontext) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
