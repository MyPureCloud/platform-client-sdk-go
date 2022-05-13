package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// V2conversationmessagetypingeventforusertopicconversationcontentstory
type V2conversationmessagetypingeventforusertopicconversationcontentstory struct { 
	// VarType
	VarType *string `json:"type,omitempty"`


	// Url
	Url *string `json:"url,omitempty"`


	// ReplyToId
	ReplyToId *string `json:"replyToId,omitempty"`

}

func (o *V2conversationmessagetypingeventforusertopicconversationcontentstory) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias V2conversationmessagetypingeventforusertopicconversationcontentstory
	
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

func (o *V2conversationmessagetypingeventforusertopicconversationcontentstory) UnmarshalJSON(b []byte) error {
	var V2conversationmessagetypingeventforusertopicconversationcontentstoryMap map[string]interface{}
	err := json.Unmarshal(b, &V2conversationmessagetypingeventforusertopicconversationcontentstoryMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := V2conversationmessagetypingeventforusertopicconversationcontentstoryMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Url, ok := V2conversationmessagetypingeventforusertopicconversationcontentstoryMap["url"].(string); ok {
		o.Url = &Url
	}
    
	if ReplyToId, ok := V2conversationmessagetypingeventforusertopicconversationcontentstoryMap["replyToId"].(string); ok {
		o.ReplyToId = &ReplyToId
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *V2conversationmessagetypingeventforusertopicconversationcontentstory) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
