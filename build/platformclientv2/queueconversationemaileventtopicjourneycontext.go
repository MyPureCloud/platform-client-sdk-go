package platformclientv2
import (
	"encoding/json"
)

// Queueconversationemaileventtopicjourneycontext
type Queueconversationemaileventtopicjourneycontext struct { 
	// Customer
	Customer *Queueconversationemaileventtopicjourneycustomer `json:"customer,omitempty"`


	// CustomerSession
	CustomerSession *Queueconversationemaileventtopicjourneycustomersession `json:"customerSession,omitempty"`


	// TriggeringAction
	TriggeringAction *Queueconversationemaileventtopicjourneyaction `json:"triggeringAction,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queueconversationemaileventtopicjourneycontext) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
