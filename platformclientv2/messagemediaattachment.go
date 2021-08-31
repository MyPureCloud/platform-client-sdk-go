package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Messagemediaattachment
type Messagemediaattachment struct { 
	// Url - The location of the media, useful for retrieving it
	Url *string `json:"url,omitempty"`


	// MediaType - The optional internet media type of the the media object.If null then the media type should be dictated by the url.
	MediaType *string `json:"mediaType,omitempty"`


	// ContentLength - The optional content length of the the media object, in bytes.
	ContentLength *int `json:"contentLength,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Id
	Id *string `json:"id,omitempty"`

}

func (o *Messagemediaattachment) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Messagemediaattachment
	
	return json.Marshal(&struct { 
		Url *string `json:"url,omitempty"`
		
		MediaType *string `json:"mediaType,omitempty"`
		
		ContentLength *int `json:"contentLength,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Id *string `json:"id,omitempty"`
		*Alias
	}{ 
		Url: o.Url,
		
		MediaType: o.MediaType,
		
		ContentLength: o.ContentLength,
		
		Name: o.Name,
		
		Id: o.Id,
		Alias:    (*Alias)(o),
	})
}

func (o *Messagemediaattachment) UnmarshalJSON(b []byte) error {
	var MessagemediaattachmentMap map[string]interface{}
	err := json.Unmarshal(b, &MessagemediaattachmentMap)
	if err != nil {
		return err
	}
	
	if Url, ok := MessagemediaattachmentMap["url"].(string); ok {
		o.Url = &Url
	}
	
	if MediaType, ok := MessagemediaattachmentMap["mediaType"].(string); ok {
		o.MediaType = &MediaType
	}
	
	if ContentLength, ok := MessagemediaattachmentMap["contentLength"].(float64); ok {
		ContentLengthInt := int(ContentLength)
		o.ContentLength = &ContentLengthInt
	}
	
	if Name, ok := MessagemediaattachmentMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if Id, ok := MessagemediaattachmentMap["id"].(string); ok {
		o.Id = &Id
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Messagemediaattachment) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
