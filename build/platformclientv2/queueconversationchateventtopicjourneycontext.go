package platformclientv2
import (
	"encoding/json"
)

// Queueconversationchateventtopicjourneycontext
type Queueconversationchateventtopicjourneycontext struct { 
	// Customer
	Customer *Queueconversationchateventtopicjourneycustomer `json:"customer,omitempty"`


	// CustomerSession
	CustomerSession *Queueconversationchateventtopicjourneycustomersession `json:"customerSession,omitempty"`


	// TriggeringAction
	TriggeringAction *Queueconversationchateventtopicjourneyaction `json:"triggeringAction,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queueconversationchateventtopicjourneycontext) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
