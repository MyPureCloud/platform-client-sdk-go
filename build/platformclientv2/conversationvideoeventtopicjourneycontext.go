package platformclientv2
import (
	"encoding/json"
)

// Conversationvideoeventtopicjourneycontext
type Conversationvideoeventtopicjourneycontext struct { 
	// Customer
	Customer *Conversationvideoeventtopicjourneycustomer `json:"customer,omitempty"`


	// CustomerSession
	CustomerSession *Conversationvideoeventtopicjourneycustomersession `json:"customerSession,omitempty"`


	// TriggeringAction
	TriggeringAction *Conversationvideoeventtopicjourneyaction `json:"triggeringAction,omitempty"`

}

// String returns a JSON representation of the model
func (o *Conversationvideoeventtopicjourneycontext) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
