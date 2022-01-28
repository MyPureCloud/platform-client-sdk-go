package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Messengerapps - The apps embedded in the messenger
type Messengerapps struct { 
	// Conversations - The conversation settings that handles chats within the messenger
	Conversations *Conversationappsettings `json:"conversations,omitempty"`


	// Knowledge - The knowledge base config for messenger
	Knowledge *Knowledge `json:"knowledge,omitempty"`

}

func (o *Messengerapps) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Messengerapps
	
	return json.Marshal(&struct { 
		Conversations *Conversationappsettings `json:"conversations,omitempty"`
		
		Knowledge *Knowledge `json:"knowledge,omitempty"`
		*Alias
	}{ 
		Conversations: o.Conversations,
		
		Knowledge: o.Knowledge,
		Alias:    (*Alias)(o),
	})
}

func (o *Messengerapps) UnmarshalJSON(b []byte) error {
	var MessengerappsMap map[string]interface{}
	err := json.Unmarshal(b, &MessengerappsMap)
	if err != nil {
		return err
	}
	
	if Conversations, ok := MessengerappsMap["conversations"].(map[string]interface{}); ok {
		ConversationsString, _ := json.Marshal(Conversations)
		json.Unmarshal(ConversationsString, &o.Conversations)
	}
	
	if Knowledge, ok := MessengerappsMap["knowledge"].(map[string]interface{}); ok {
		KnowledgeString, _ := json.Marshal(Knowledge)
		json.Unmarshal(KnowledgeString, &o.Knowledge)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Messengerapps) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
