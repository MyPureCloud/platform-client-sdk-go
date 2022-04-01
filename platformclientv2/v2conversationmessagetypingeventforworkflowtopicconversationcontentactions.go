package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// V2conversationmessagetypingeventforworkflowtopicconversationcontentactions
type V2conversationmessagetypingeventforworkflowtopicconversationcontentactions struct { 
	// Url
	Url *string `json:"url,omitempty"`


	// UrlTarget
	UrlTarget *string `json:"urlTarget,omitempty"`


	// Textback
	Textback *string `json:"textback,omitempty"`

}

func (o *V2conversationmessagetypingeventforworkflowtopicconversationcontentactions) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias V2conversationmessagetypingeventforworkflowtopicconversationcontentactions
	
	return json.Marshal(&struct { 
		Url *string `json:"url,omitempty"`
		
		UrlTarget *string `json:"urlTarget,omitempty"`
		
		Textback *string `json:"textback,omitempty"`
		*Alias
	}{ 
		Url: o.Url,
		
		UrlTarget: o.UrlTarget,
		
		Textback: o.Textback,
		Alias:    (*Alias)(o),
	})
}

func (o *V2conversationmessagetypingeventforworkflowtopicconversationcontentactions) UnmarshalJSON(b []byte) error {
	var V2conversationmessagetypingeventforworkflowtopicconversationcontentactionsMap map[string]interface{}
	err := json.Unmarshal(b, &V2conversationmessagetypingeventforworkflowtopicconversationcontentactionsMap)
	if err != nil {
		return err
	}
	
	if Url, ok := V2conversationmessagetypingeventforworkflowtopicconversationcontentactionsMap["url"].(string); ok {
		o.Url = &Url
	}
	
	if UrlTarget, ok := V2conversationmessagetypingeventforworkflowtopicconversationcontentactionsMap["urlTarget"].(string); ok {
		o.UrlTarget = &UrlTarget
	}
	
	if Textback, ok := V2conversationmessagetypingeventforworkflowtopicconversationcontentactionsMap["textback"].(string); ok {
		o.Textback = &Textback
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *V2conversationmessagetypingeventforworkflowtopicconversationcontentactions) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
