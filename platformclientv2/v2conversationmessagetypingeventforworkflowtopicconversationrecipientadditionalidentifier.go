package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// V2conversationmessagetypingeventforworkflowtopicconversationrecipientadditionalidentifier
type V2conversationmessagetypingeventforworkflowtopicconversationrecipientadditionalidentifier struct { 
	// VarType
	VarType *string `json:"type,omitempty"`


	// Value
	Value *string `json:"value,omitempty"`

}

func (o *V2conversationmessagetypingeventforworkflowtopicconversationrecipientadditionalidentifier) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias V2conversationmessagetypingeventforworkflowtopicconversationrecipientadditionalidentifier
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		Value *string `json:"value,omitempty"`
		*Alias
	}{ 
		VarType: o.VarType,
		
		Value: o.Value,
		Alias:    (*Alias)(o),
	})
}

func (o *V2conversationmessagetypingeventforworkflowtopicconversationrecipientadditionalidentifier) UnmarshalJSON(b []byte) error {
	var V2conversationmessagetypingeventforworkflowtopicconversationrecipientadditionalidentifierMap map[string]interface{}
	err := json.Unmarshal(b, &V2conversationmessagetypingeventforworkflowtopicconversationrecipientadditionalidentifierMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := V2conversationmessagetypingeventforworkflowtopicconversationrecipientadditionalidentifierMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Value, ok := V2conversationmessagetypingeventforworkflowtopicconversationrecipientadditionalidentifierMap["value"].(string); ok {
		o.Value = &Value
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *V2conversationmessagetypingeventforworkflowtopicconversationrecipientadditionalidentifier) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
