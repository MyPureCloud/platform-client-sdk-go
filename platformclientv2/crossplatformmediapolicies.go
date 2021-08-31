package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Crossplatformmediapolicies
type Crossplatformmediapolicies struct { 
	// CallPolicy - Conditions and actions for calls
	CallPolicy *Crossplatformcallmediapolicy `json:"callPolicy,omitempty"`


	// ChatPolicy - Conditions and actions for chats
	ChatPolicy *Crossplatformchatmediapolicy `json:"chatPolicy,omitempty"`


	// EmailPolicy - Conditions and actions for emails
	EmailPolicy *Crossplatformemailmediapolicy `json:"emailPolicy,omitempty"`


	// MessagePolicy - Conditions and actions for messages
	MessagePolicy *Crossplatformmessagemediapolicy `json:"messagePolicy,omitempty"`

}

func (o *Crossplatformmediapolicies) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Crossplatformmediapolicies
	
	return json.Marshal(&struct { 
		CallPolicy *Crossplatformcallmediapolicy `json:"callPolicy,omitempty"`
		
		ChatPolicy *Crossplatformchatmediapolicy `json:"chatPolicy,omitempty"`
		
		EmailPolicy *Crossplatformemailmediapolicy `json:"emailPolicy,omitempty"`
		
		MessagePolicy *Crossplatformmessagemediapolicy `json:"messagePolicy,omitempty"`
		*Alias
	}{ 
		CallPolicy: o.CallPolicy,
		
		ChatPolicy: o.ChatPolicy,
		
		EmailPolicy: o.EmailPolicy,
		
		MessagePolicy: o.MessagePolicy,
		Alias:    (*Alias)(o),
	})
}

func (o *Crossplatformmediapolicies) UnmarshalJSON(b []byte) error {
	var CrossplatformmediapoliciesMap map[string]interface{}
	err := json.Unmarshal(b, &CrossplatformmediapoliciesMap)
	if err != nil {
		return err
	}
	
	if CallPolicy, ok := CrossplatformmediapoliciesMap["callPolicy"].(map[string]interface{}); ok {
		CallPolicyString, _ := json.Marshal(CallPolicy)
		json.Unmarshal(CallPolicyString, &o.CallPolicy)
	}
	
	if ChatPolicy, ok := CrossplatformmediapoliciesMap["chatPolicy"].(map[string]interface{}); ok {
		ChatPolicyString, _ := json.Marshal(ChatPolicy)
		json.Unmarshal(ChatPolicyString, &o.ChatPolicy)
	}
	
	if EmailPolicy, ok := CrossplatformmediapoliciesMap["emailPolicy"].(map[string]interface{}); ok {
		EmailPolicyString, _ := json.Marshal(EmailPolicy)
		json.Unmarshal(EmailPolicyString, &o.EmailPolicy)
	}
	
	if MessagePolicy, ok := CrossplatformmediapoliciesMap["messagePolicy"].(map[string]interface{}); ok {
		MessagePolicyString, _ := json.Marshal(MessagePolicy)
		json.Unmarshal(MessagePolicyString, &o.MessagePolicy)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Crossplatformmediapolicies) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
