package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationcontentstory - Story object.
type Conversationcontentstory struct { 
	// VarType - Type of ephemeral story attachment.
	VarType *string `json:"type,omitempty"`


	// Url - URL to the ephemeral story.
	Url *string `json:"url,omitempty"`


	// ReplyToId - ID of the ephemeral story being replied to.
	ReplyToId *string `json:"replyToId,omitempty"`

}

func (o *Conversationcontentstory) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationcontentstory
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		Url *string `json:"url,omitempty"`
		
		ReplyToId *string `json:"replyToId,omitempty"`
		*Alias
	}{ 
		VarType: o.VarType,
		
		Url: o.Url,
		
		ReplyToId: o.ReplyToId,
		Alias:    (*Alias)(o),
	})
}

func (o *Conversationcontentstory) UnmarshalJSON(b []byte) error {
	var ConversationcontentstoryMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationcontentstoryMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := ConversationcontentstoryMap["type"].(string); ok {
		o.VarType = &VarType
	}
	
	if Url, ok := ConversationcontentstoryMap["url"].(string); ok {
		o.Url = &Url
	}
	
	if ReplyToId, ok := ConversationcontentstoryMap["replyToId"].(string); ok {
		o.ReplyToId = &ReplyToId
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationcontentstory) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
