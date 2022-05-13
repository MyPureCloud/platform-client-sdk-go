package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationcontentbuttonresponse - Button response object representing the click of a structured message button, such as a quick reply.
type Conversationcontentbuttonresponse struct { 
	// VarType - Describes the button that resulted in the Button Response.
	VarType *string `json:"type,omitempty"`


	// Text - The response text from the button click.
	Text *string `json:"text,omitempty"`


	// Payload - The response payload associated with the clicked button.
	Payload *string `json:"payload,omitempty"`

}

func (o *Conversationcontentbuttonresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationcontentbuttonresponse
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		Text *string `json:"text,omitempty"`
		
		Payload *string `json:"payload,omitempty"`
		*Alias
	}{ 
		VarType: o.VarType,
		
		Text: o.Text,
		
		Payload: o.Payload,
		Alias:    (*Alias)(o),
	})
}

func (o *Conversationcontentbuttonresponse) UnmarshalJSON(b []byte) error {
	var ConversationcontentbuttonresponseMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationcontentbuttonresponseMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := ConversationcontentbuttonresponseMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Text, ok := ConversationcontentbuttonresponseMap["text"].(string); ok {
		o.Text = &Text
	}
    
	if Payload, ok := ConversationcontentbuttonresponseMap["payload"].(string); ok {
		o.Payload = &Payload
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationcontentbuttonresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
