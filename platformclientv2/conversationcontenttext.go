package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationcontenttext - Message content element containing text only.
type Conversationcontenttext struct { 
	// VarType - Type of text content (Deprecated).
	VarType *string `json:"type,omitempty"`


	// Body - Text to be shown for this content element.
	Body *string `json:"body,omitempty"`

}

func (o *Conversationcontenttext) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationcontenttext
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		Body *string `json:"body,omitempty"`
		*Alias
	}{ 
		VarType: o.VarType,
		
		Body: o.Body,
		Alias:    (*Alias)(o),
	})
}

func (o *Conversationcontenttext) UnmarshalJSON(b []byte) error {
	var ConversationcontenttextMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationcontenttextMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := ConversationcontenttextMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Body, ok := ConversationcontenttextMap["body"].(string); ok {
		o.Body = &Body
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationcontenttext) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
