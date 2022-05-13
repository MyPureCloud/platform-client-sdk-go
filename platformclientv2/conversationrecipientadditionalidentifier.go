package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationrecipientadditionalidentifier - Additional identifiers for describing messaging recipient.
type Conversationrecipientadditionalidentifier struct { 
	// VarType - Type of the Identifier
	VarType *string `json:"type,omitempty"`


	// Value - The Identifier value.
	Value *string `json:"value,omitempty"`

}

func (o *Conversationrecipientadditionalidentifier) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationrecipientadditionalidentifier
	
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

func (o *Conversationrecipientadditionalidentifier) UnmarshalJSON(b []byte) error {
	var ConversationrecipientadditionalidentifierMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationrecipientadditionalidentifierMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := ConversationrecipientadditionalidentifierMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Value, ok := ConversationrecipientadditionalidentifierMap["value"].(string); ok {
		o.Value = &Value
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationrecipientadditionalidentifier) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
