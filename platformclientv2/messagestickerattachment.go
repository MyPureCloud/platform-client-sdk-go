package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Messagestickerattachment
type Messagestickerattachment struct { 
	// Url - The location of the media, useful for retrieving it
	Url *string `json:"url,omitempty"`


	// Id - A globally unique identifier for the media object.
	Id *string `json:"id,omitempty"`

}

func (o *Messagestickerattachment) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Messagestickerattachment
	
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

func (o *Messagestickerattachment) UnmarshalJSON(b []byte) error {
	var MessagestickerattachmentMap map[string]interface{}
	err := json.Unmarshal(b, &MessagestickerattachmentMap)
	if err != nil {
		return err
	}
	
	if Url, ok := MessagestickerattachmentMap["url"].(string); ok {
		o.Url = &Url
	}
    
	if Id, ok := MessagestickerattachmentMap["id"].(string); ok {
		o.Id = &Id
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Messagestickerattachment) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
