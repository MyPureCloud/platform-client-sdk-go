package platformclientv2
import (
	"encoding/json"
)

// Queueconversationcallbackeventtopicjourneycontext
type Queueconversationcallbackeventtopicjourneycontext struct { 
	// Customer
	Customer *Queueconversationcallbackeventtopicjourneycustomer `json:"customer,omitempty"`


	// CustomerSession
	CustomerSession *Queueconversationcallbackeventtopicjourneycustomersession `json:"customerSession,omitempty"`


	// TriggeringAction
	TriggeringAction *Queueconversationcallbackeventtopicjourneyaction `json:"triggeringAction,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queueconversationcallbackeventtopicjourneycontext) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
