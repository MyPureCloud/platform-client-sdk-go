package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Webmessagingquickreply - Quick reply object
type Webmessagingquickreply struct { 
	// Text - Text to show inside the quick reply. This is also used as the response text after clicking on the quick reply.
	Text *string `json:"text,omitempty"`


	// Payload - Content of the payload included in the quick reply response. Could be an ID identifying the quick reply response.
	Payload *string `json:"payload,omitempty"`


	// Image - URL of an image associated with the quick reply.
	Image *string `json:"image,omitempty"`


	// Action - Specifies the type of action that is triggered upon clicking the quick reply.
	Action *string `json:"action,omitempty"`

}

func (o *Webmessagingquickreply) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Webmessagingquickreply
	
	return json.Marshal(&struct { 
		Text *string `json:"text,omitempty"`
		
		Payload *string `json:"payload,omitempty"`
		
		Image *string `json:"image,omitempty"`
		
		Action *string `json:"action,omitempty"`
		*Alias
	}{ 
		Text: o.Text,
		
		Payload: o.Payload,
		
		Image: o.Image,
		
		Action: o.Action,
		Alias:    (*Alias)(o),
	})
}

func (o *Webmessagingquickreply) UnmarshalJSON(b []byte) error {
	var WebmessagingquickreplyMap map[string]interface{}
	err := json.Unmarshal(b, &WebmessagingquickreplyMap)
	if err != nil {
		return err
	}
	
	if Text, ok := WebmessagingquickreplyMap["text"].(string); ok {
		o.Text = &Text
	}
    
	if Payload, ok := WebmessagingquickreplyMap["payload"].(string); ok {
		o.Payload = &Payload
	}
    
	if Image, ok := WebmessagingquickreplyMap["image"].(string); ok {
		o.Image = &Image
	}
    
	if Action, ok := WebmessagingquickreplyMap["action"].(string); ok {
		o.Action = &Action
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Webmessagingquickreply) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}