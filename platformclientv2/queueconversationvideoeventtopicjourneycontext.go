package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
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
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
