package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationmessageeventtopicjourneycontext
type Conversationmessageeventtopicjourneycontext struct { 
	// Customer
	Customer *Conversationmessageeventtopicjourneycustomer `json:"customer,omitempty"`


	// CustomerSession
	CustomerSession *Conversationmessageeventtopicjourneycustomersession `json:"customerSession,omitempty"`


	// TriggeringAction
	TriggeringAction *Conversationmessageeventtopicjourneyaction `json:"triggeringAction,omitempty"`

}

func (o *Conversationmessageeventtopicjourneycontext) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationmessageeventtopicjourneycontext
	
	return json.Marshal(&struct { 
		Customer *Conversationmessageeventtopicjourneycustomer `json:"customer,omitempty"`
		
		CustomerSession *Conversationmessageeventtopicjourneycustomersession `json:"customerSession,omitempty"`
		
		TriggeringAction *Conversationmessageeventtopicjourneyaction `json:"triggeringAction,omitempty"`
		*Alias
	}{ 
		Customer: o.Customer,
		
		CustomerSession: o.CustomerSession,
		
		TriggeringAction: o.TriggeringAction,
		Alias:    (*Alias)(o),
	})
}

func (o *Conversationmessageeventtopicjourneycontext) UnmarshalJSON(b []byte) error {
	var ConversationmessageeventtopicjourneycontextMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationmessageeventtopicjourneycontextMap)
	if err != nil {
		return err
	}
	
	if Customer, ok := ConversationmessageeventtopicjourneycontextMap["customer"].(map[string]interface{}); ok {
		CustomerString, _ := json.Marshal(Customer)
		json.Unmarshal(CustomerString, &o.Customer)
	}
	
	if CustomerSession, ok := ConversationmessageeventtopicjourneycontextMap["customerSession"].(map[string]interface{}); ok {
		CustomerSessionString, _ := json.Marshal(CustomerSession)
		json.Unmarshal(CustomerSessionString, &o.CustomerSession)
	}
	
	if TriggeringAction, ok := ConversationmessageeventtopicjourneycontextMap["triggeringAction"].(map[string]interface{}); ok {
		TriggeringActionString, _ := json.Marshal(TriggeringAction)
		json.Unmarshal(TriggeringActionString, &o.TriggeringAction)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationmessageeventtopicjourneycontext) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
