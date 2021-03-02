package platformclientv2
import (
	"encoding/json"
)

// Queueconversationsocialexpressioneventtopicjourneycontext
type Queueconversationsocialexpressioneventtopicjourneycontext struct { 
	// Customer
	Customer *Queueconversationsocialexpressioneventtopicjourneycustomer `json:"customer,omitempty"`


	// CustomerSession
	CustomerSession *Queueconversationsocialexpressioneventtopicjourneycustomersession `json:"customerSession,omitempty"`


	// TriggeringAction
	TriggeringAction *Queueconversationsocialexpressioneventtopicjourneyaction `json:"triggeringAction,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queueconversationsocialexpressioneventtopicjourneycontext) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
