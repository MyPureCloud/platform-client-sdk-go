package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationemaileventtopicjourneycontext
type Conversationemaileventtopicjourneycontext struct { 
	// Customer
	Customer *Conversationemaileventtopicjourneycustomer `json:"customer,omitempty"`


	// CustomerSession
	CustomerSession *Conversationemaileventtopicjourneycustomersession `json:"customerSession,omitempty"`


	// TriggeringAction
	TriggeringAction *Conversationemaileventtopicjourneyaction `json:"triggeringAction,omitempty"`

}

// String returns a JSON representation of the model
func (o *Conversationemaileventtopicjourneycontext) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
