package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Posttextmessage
type Posttextmessage struct { 
	// VarType - Message type
	VarType *string `json:"type,omitempty"`


	// Text - Message text. If type is structured, used as fallback for clients that do not support particular structured content
	Text *string `json:"text,omitempty"`


	// Content - A list of content elements in message
	Content *[]Conversationmessagecontent `json:"content,omitempty"`

}

func (o *Posttextmessage) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Posttextmessage
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		Text *string `json:"text,omitempty"`
		
		Content *[]Conversationmessagecontent `json:"content,omitempty"`
		*Alias
	}{ 
		VarType: o.VarType,
		
		Text: o.Text,
		
		Content: o.Content,
		Alias:    (*Alias)(o),
	})
}

func (o *Posttextmessage) UnmarshalJSON(b []byte) error {
	var PosttextmessageMap map[string]interface{}
	err := json.Unmarshal(b, &PosttextmessageMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := PosttextmessageMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Text, ok := PosttextmessageMap["text"].(string); ok {
		o.Text = &Text
	}
    
	if Content, ok := PosttextmessageMap["content"].([]interface{}); ok {
		ContentString, _ := json.Marshal(Content)
		json.Unmarshal(ContentString, &o.Content)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Posttextmessage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
