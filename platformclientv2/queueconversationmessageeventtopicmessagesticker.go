package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationmessageeventtopicmessagesticker
type Queueconversationmessageeventtopicmessagesticker struct { 
	// Url
	Url *string `json:"url,omitempty"`


	// Id
	Id *string `json:"id,omitempty"`

}

func (o *Queueconversationmessageeventtopicmessagesticker) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queueconversationmessageeventtopicmessagesticker
	
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

func (o *Queueconversationmessageeventtopicmessagesticker) UnmarshalJSON(b []byte) error {
	var QueueconversationmessageeventtopicmessagestickerMap map[string]interface{}
	err := json.Unmarshal(b, &QueueconversationmessageeventtopicmessagestickerMap)
	if err != nil {
		return err
	}
	
	if Url, ok := QueueconversationmessageeventtopicmessagestickerMap["url"].(string); ok {
		o.Url = &Url
	}
	
	if Id, ok := QueueconversationmessageeventtopicmessagestickerMap["id"].(string); ok {
		o.Id = &Id
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Queueconversationmessageeventtopicmessagesticker) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
