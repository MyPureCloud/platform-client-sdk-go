package platformclientv2
import (
	"encoding/json"
)

// Queueconversationcalleventtopicjourneycontext
type Queueconversationcalleventtopicjourneycontext struct { 
	// Customer
	Customer *Queueconversationcalleventtopicjourneycustomer `json:"customer,omitempty"`


	// CustomerSession
	CustomerSession *Queueconversationcalleventtopicjourneycustomersession `json:"customerSession,omitempty"`


	// TriggeringAction
	TriggeringAction *Queueconversationcalleventtopicjourneyaction `json:"triggeringAction,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queueconversationcalleventtopicjourneycontext) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
