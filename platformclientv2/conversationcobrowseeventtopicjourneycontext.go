package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationcobrowseeventtopicjourneycontext
type Conversationcobrowseeventtopicjourneycontext struct { 
	// Customer
	Customer *Conversationcobrowseeventtopicjourneycustomer `json:"customer,omitempty"`


	// CustomerSession
	CustomerSession *Conversationcobrowseeventtopicjourneycustomersession `json:"customerSession,omitempty"`


	// TriggeringAction
	TriggeringAction *Conversationcobrowseeventtopicjourneyaction `json:"triggeringAction,omitempty"`

}

// String returns a JSON representation of the model
func (o *Conversationcobrowseeventtopicjourneycontext) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
