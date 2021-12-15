package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationvideoeventtopicmessagesticker
type Queueconversationvideoeventtopicmessagesticker struct { 
	// Url - The location of the sticker, useful for retrieving it
	Url *string `json:"url,omitempty"`


	// Id - The unique id of the the sticker object.
	Id *string `json:"id,omitempty"`

}

func (o *Queueconversationvideoeventtopicmessagesticker) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queueconversationvideoeventtopicmessagesticker
	
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

func (o *Queueconversationvideoeventtopicmessagesticker) UnmarshalJSON(b []byte) error {
	var QueueconversationvideoeventtopicmessagestickerMap map[string]interface{}
	err := json.Unmarshal(b, &QueueconversationvideoeventtopicmessagestickerMap)
	if err != nil {
		return err
	}
	
	if Url, ok := QueueconversationvideoeventtopicmessagestickerMap["url"].(string); ok {
		o.Url = &Url
	}
	
	if Id, ok := QueueconversationvideoeventtopicmessagestickerMap["id"].(string); ok {
		o.Id = &Id
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Queueconversationvideoeventtopicmessagesticker) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
