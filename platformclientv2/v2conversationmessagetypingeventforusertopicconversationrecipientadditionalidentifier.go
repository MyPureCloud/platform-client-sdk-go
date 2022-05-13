package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// V2conversationmessagetypingeventforusertopicconversationrecipientadditionalidentifier
type V2conversationmessagetypingeventforusertopicconversationrecipientadditionalidentifier struct { 
	// VarType
	VarType *string `json:"type,omitempty"`


	// Value
	Value *string `json:"value,omitempty"`

}

func (o *V2conversationmessagetypingeventforusertopicconversationrecipientadditionalidentifier) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias V2conversationmessagetypingeventforusertopicconversationrecipientadditionalidentifier
	
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

func (o *V2conversationmessagetypingeventforusertopicconversationrecipientadditionalidentifier) UnmarshalJSON(b []byte) error {
	var V2conversationmessagetypingeventforusertopicconversationrecipientadditionalidentifierMap map[string]interface{}
	err := json.Unmarshal(b, &V2conversationmessagetypingeventforusertopicconversationrecipientadditionalidentifierMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := V2conversationmessagetypingeventforusertopicconversationrecipientadditionalidentifierMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Value, ok := V2conversationmessagetypingeventforusertopicconversationrecipientadditionalidentifierMap["value"].(string); ok {
		o.Value = &Value
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *V2conversationmessagetypingeventforusertopicconversationrecipientadditionalidentifier) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
