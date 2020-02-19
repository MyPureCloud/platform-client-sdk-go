package platformclientv2
import (
	"encoding/json"
)

// Queueconversationeventtopicjourneycontext
type Queueconversationeventtopicjourneycontext struct { 
	// Customer
	Customer *Queueconversationeventtopicjourneycustomer `json:"customer,omitempty"`


	// CustomerSession
	CustomerSession *Queueconversationeventtopicjourneycustomersession `json:"customerSession,omitempty"`


	// TriggeringAction
	TriggeringAction *Queueconversationeventtopicjourneyaction `json:"triggeringAction,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queueconversationeventtopicjourneycontext) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
