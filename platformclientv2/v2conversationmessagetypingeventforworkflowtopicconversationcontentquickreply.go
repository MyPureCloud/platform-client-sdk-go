package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// V2conversationmessagetypingeventforworkflowtopicconversationcontentquickreply
type V2conversationmessagetypingeventforworkflowtopicconversationcontentquickreply struct { 
	// Text
	Text *string `json:"text,omitempty"`


	// Payload
	Payload *string `json:"payload,omitempty"`


	// Image
	Image *string `json:"image,omitempty"`


	// Action
	Action *string `json:"action,omitempty"`

}

func (o *V2conversationmessagetypingeventforworkflowtopicconversationcontentquickreply) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias V2conversationmessagetypingeventforworkflowtopicconversationcontentquickreply
	
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

func (o *V2conversationmessagetypingeventforworkflowtopicconversationcontentquickreply) UnmarshalJSON(b []byte) error {
	var V2conversationmessagetypingeventforworkflowtopicconversationcontentquickreplyMap map[string]interface{}
	err := json.Unmarshal(b, &V2conversationmessagetypingeventforworkflowtopicconversationcontentquickreplyMap)
	if err != nil {
		return err
	}
	
	if Text, ok := V2conversationmessagetypingeventforworkflowtopicconversationcontentquickreplyMap["text"].(string); ok {
		o.Text = &Text
	}
    
	if Payload, ok := V2conversationmessagetypingeventforworkflowtopicconversationcontentquickreplyMap["payload"].(string); ok {
		o.Payload = &Payload
	}
    
	if Image, ok := V2conversationmessagetypingeventforworkflowtopicconversationcontentquickreplyMap["image"].(string); ok {
		o.Image = &Image
	}
    
	if Action, ok := V2conversationmessagetypingeventforworkflowtopicconversationcontentquickreplyMap["action"].(string); ok {
		o.Action = &Action
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *V2conversationmessagetypingeventforworkflowtopicconversationcontentquickreply) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
