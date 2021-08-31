package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationeventtopicmessagemedia
type Conversationeventtopicmessagemedia struct { 
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

func (o *Conversationeventtopicmessagemedia) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationeventtopicmessagemedia
	
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

func (o *Conversationeventtopicmessagemedia) UnmarshalJSON(b []byte) error {
	var ConversationeventtopicmessagemediaMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationeventtopicmessagemediaMap)
	if err != nil {
		return err
	}
	
	if Url, ok := ConversationeventtopicmessagemediaMap["url"].(string); ok {
		o.Url = &Url
	}
	
	if MediaType, ok := ConversationeventtopicmessagemediaMap["mediaType"].(string); ok {
		o.MediaType = &MediaType
	}
	
	if ContentLengthBytes, ok := ConversationeventtopicmessagemediaMap["contentLengthBytes"].(float64); ok {
		ContentLengthBytesInt := int(ContentLengthBytes)
		o.ContentLengthBytes = &ContentLengthBytesInt
	}
	
	if Name, ok := ConversationeventtopicmessagemediaMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if Id, ok := ConversationeventtopicmessagemediaMap["id"].(string); ok {
		o.Id = &Id
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationeventtopicmessagemedia) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
