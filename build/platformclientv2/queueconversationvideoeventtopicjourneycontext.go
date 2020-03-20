package platformclientv2
import (
	"encoding/json"
)

// Queueconversationvideoeventtopicjourneycontext
type Queueconversationvideoeventtopicjourneycontext struct { 
	// Customer
	Customer *Queueconversationvideoeventtopicjourneycustomer `json:"customer,omitempty"`


	// CustomerSession
	CustomerSession *Queueconversationvideoeventtopicjourneycustomersession `json:"customerSession,omitempty"`


	// TriggeringAction
	TriggeringAction *Queueconversationvideoeventtopicjourneyaction `json:"triggeringAction,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queueconversationvideoeventtopicjourneycontext) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
