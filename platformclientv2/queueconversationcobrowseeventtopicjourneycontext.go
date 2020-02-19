package platformclientv2
import (
	"encoding/json"
)

// Queueconversationcobrowseeventtopicjourneycontext
type Queueconversationcobrowseeventtopicjourneycontext struct { 
	// Customer
	Customer *Queueconversationcobrowseeventtopicjourneycustomer `json:"customer,omitempty"`


	// CustomerSession
	CustomerSession *Queueconversationcobrowseeventtopicjourneycustomersession `json:"customerSession,omitempty"`


	// TriggeringAction
	TriggeringAction *Queueconversationcobrowseeventtopicjourneyaction `json:"triggeringAction,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queueconversationcobrowseeventtopicjourneycontext) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
