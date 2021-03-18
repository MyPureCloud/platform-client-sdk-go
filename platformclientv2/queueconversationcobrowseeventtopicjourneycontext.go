package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
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
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
