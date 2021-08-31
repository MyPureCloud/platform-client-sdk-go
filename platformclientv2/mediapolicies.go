package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Mediapolicies
type Mediapolicies struct { 
	// CallPolicy - Conditions and actions for calls
	CallPolicy *Callmediapolicy `json:"callPolicy,omitempty"`


	// ChatPolicy - Conditions and actions for chats
	ChatPolicy *Chatmediapolicy `json:"chatPolicy,omitempty"`


	// EmailPolicy - Conditions and actions for emails
	EmailPolicy *Emailmediapolicy `json:"emailPolicy,omitempty"`


	// MessagePolicy - Conditions and actions for messages
	MessagePolicy *Messagemediapolicy `json:"messagePolicy,omitempty"`

}

func (o *Mediapolicies) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Mediapolicies
	
	return json.Marshal(&struct { 
		CallPolicy *Callmediapolicy `json:"callPolicy,omitempty"`
		
		ChatPolicy *Chatmediapolicy `json:"chatPolicy,omitempty"`
		
		EmailPolicy *Emailmediapolicy `json:"emailPolicy,omitempty"`
		
		MessagePolicy *Messagemediapolicy `json:"messagePolicy,omitempty"`
		*Alias
	}{ 
		CallPolicy: o.CallPolicy,
		
		ChatPolicy: o.ChatPolicy,
		
		EmailPolicy: o.EmailPolicy,
		
		MessagePolicy: o.MessagePolicy,
		Alias:    (*Alias)(o),
	})
}

func (o *Mediapolicies) UnmarshalJSON(b []byte) error {
	var MediapoliciesMap map[string]interface{}
	err := json.Unmarshal(b, &MediapoliciesMap)
	if err != nil {
		return err
	}
	
	if CallPolicy, ok := MediapoliciesMap["callPolicy"].(map[string]interface{}); ok {
		CallPolicyString, _ := json.Marshal(CallPolicy)
		json.Unmarshal(CallPolicyString, &o.CallPolicy)
	}
	
	if ChatPolicy, ok := MediapoliciesMap["chatPolicy"].(map[string]interface{}); ok {
		ChatPolicyString, _ := json.Marshal(ChatPolicy)
		json.Unmarshal(ChatPolicyString, &o.ChatPolicy)
	}
	
	if EmailPolicy, ok := MediapoliciesMap["emailPolicy"].(map[string]interface{}); ok {
		EmailPolicyString, _ := json.Marshal(EmailPolicy)
		json.Unmarshal(EmailPolicyString, &o.EmailPolicy)
	}
	
	if MessagePolicy, ok := MediapoliciesMap["messagePolicy"].(map[string]interface{}); ok {
		MessagePolicyString, _ := json.Marshal(MessagePolicy)
		json.Unmarshal(MessagePolicyString, &o.MessagePolicy)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Mediapolicies) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
