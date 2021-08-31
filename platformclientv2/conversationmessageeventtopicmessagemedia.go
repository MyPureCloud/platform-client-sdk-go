package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationmessageeventtopicmessagemedia
type Conversationmessageeventtopicmessagemedia struct { 
	// Url
	Url *string `json:"url,omitempty"`


	// MediaType
	MediaType *string `json:"mediaType,omitempty"`


	// ContentLengthBytes
	ContentLengthBytes *int `json:"contentLengthBytes,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Id
	Id *string `json:"id,omitempty"`

}

func (o *Conversationmessageeventtopicmessagemedia) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationmessageeventtopicmessagemedia
	
	return json.Marshal(&struct { 
		Url *string `json:"url,omitempty"`
		
		MediaType *string `json:"mediaType,omitempty"`
		
		ContentLengthBytes *int `json:"contentLengthBytes,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Id *string `json:"id,omitempty"`
		*Alias
	}{ 
		Url: o.Url,
		
		MediaType: o.MediaType,
		
		ContentLengthBytes: o.ContentLengthBytes,
		
		Name: o.Name,
		
		Id: o.Id,
		Alias:    (*Alias)(o),
	})
}

func (o *Conversationmessageeventtopicmessagemedia) UnmarshalJSON(b []byte) error {
	var ConversationmessageeventtopicmessagemediaMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationmessageeventtopicmessagemediaMap)
	if err != nil {
		return err
	}
	
	if Url, ok := ConversationmessageeventtopicmessagemediaMap["url"].(string); ok {
		o.Url = &Url
	}
	
	if MediaType, ok := ConversationmessageeventtopicmessagemediaMap["mediaType"].(string); ok {
		o.MediaType = &MediaType
	}
	
	if ContentLengthBytes, ok := ConversationmessageeventtopicmessagemediaMap["contentLengthBytes"].(float64); ok {
		ContentLengthBytesInt := int(ContentLengthBytes)
		o.ContentLengthBytes = &ContentLengthBytesInt
	}
	
	if Name, ok := ConversationmessageeventtopicmessagemediaMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if Id, ok := ConversationmessageeventtopicmessagemediaMap["id"].(string); ok {
		o.Id = &Id
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationmessageeventtopicmessagemedia) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
