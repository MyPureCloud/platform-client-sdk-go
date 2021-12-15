package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationmessageeventtopicmessagesticker
type Conversationmessageeventtopicmessagesticker struct { 
	// Url - The location of the sticker, useful for retrieving it
	Url *string `json:"url,omitempty"`


	// Id - The unique id of the the sticker object.
	Id *string `json:"id,omitempty"`

}

func (o *Conversationmessageeventtopicmessagesticker) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationmessageeventtopicmessagesticker
	
	return json.Marshal(&struct { 
		Url *string `json:"url,omitempty"`
		
		Id *string `json:"id,omitempty"`
		*Alias
	}{ 
		Url: o.Url,
		
		Id: o.Id,
		Alias:    (*Alias)(o),
	})
}

func (o *Conversationmessageeventtopicmessagesticker) UnmarshalJSON(b []byte) error {
	var ConversationmessageeventtopicmessagestickerMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationmessageeventtopicmessagestickerMap)
	if err != nil {
		return err
	}
	
	if Url, ok := ConversationmessageeventtopicmessagestickerMap["url"].(string); ok {
		o.Url = &Url
	}
	
	if Id, ok := ConversationmessageeventtopicmessagestickerMap["id"].(string); ok {
		o.Id = &Id
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationmessageeventtopicmessagesticker) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
