package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// V2conversationmessagetypingeventforworkflowtopicconversationnotificationtemplateheader
type V2conversationmessagetypingeventforworkflowtopicconversationnotificationtemplateheader struct { 
	// VarType
	VarType *string `json:"type,omitempty"`


	// Text
	Text *string `json:"text,omitempty"`


	// Media
	Media *V2conversationmessagetypingeventforworkflowtopicconversationcontentattachment `json:"media,omitempty"`


	// Parameters
	Parameters *[]V2conversationmessagetypingeventforworkflowtopicconversationnotificationtemplateparameter `json:"parameters,omitempty"`

}

func (o *V2conversationmessagetypingeventforworkflowtopicconversationnotificationtemplateheader) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias V2conversationmessagetypingeventforworkflowtopicconversationnotificationtemplateheader
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		Text *string `json:"text,omitempty"`
		
		Media *V2conversationmessagetypingeventforworkflowtopicconversationcontentattachment `json:"media,omitempty"`
		
		Parameters *[]V2conversationmessagetypingeventforworkflowtopicconversationnotificationtemplateparameter `json:"parameters,omitempty"`
		*Alias
	}{ 
		VarType: o.VarType,
		
		Text: o.Text,
		
		Media: o.Media,
		
		Parameters: o.Parameters,
		Alias:    (*Alias)(o),
	})
}

func (o *V2conversationmessagetypingeventforworkflowtopicconversationnotificationtemplateheader) UnmarshalJSON(b []byte) error {
	var V2conversationmessagetypingeventforworkflowtopicconversationnotificationtemplateheaderMap map[string]interface{}
	err := json.Unmarshal(b, &V2conversationmessagetypingeventforworkflowtopicconversationnotificationtemplateheaderMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := V2conversationmessagetypingeventforworkflowtopicconversationnotificationtemplateheaderMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Text, ok := V2conversationmessagetypingeventforworkflowtopicconversationnotificationtemplateheaderMap["text"].(string); ok {
		o.Text = &Text
	}
    
	if Media, ok := V2conversationmessagetypingeventforworkflowtopicconversationnotificationtemplateheaderMap["media"].(map[string]interface{}); ok {
		MediaString, _ := json.Marshal(Media)
		json.Unmarshal(MediaString, &o.Media)
	}
	
	if Parameters, ok := V2conversationmessagetypingeventforworkflowtopicconversationnotificationtemplateheaderMap["parameters"].([]interface{}); ok {
		ParametersString, _ := json.Marshal(Parameters)
		json.Unmarshal(ParametersString, &o.Parameters)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *V2conversationmessagetypingeventforworkflowtopicconversationnotificationtemplateheader) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
