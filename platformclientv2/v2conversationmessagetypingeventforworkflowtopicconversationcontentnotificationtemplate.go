package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// V2conversationmessagetypingeventforworkflowtopicconversationcontentnotificationtemplate
type V2conversationmessagetypingeventforworkflowtopicconversationcontentnotificationtemplate struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Language
	Language *string `json:"language,omitempty"`


	// Header
	Header *V2conversationmessagetypingeventforworkflowtopicconversationnotificationtemplateheader `json:"header,omitempty"`


	// Body
	Body *V2conversationmessagetypingeventforworkflowtopicconversationnotificationtemplatebody `json:"body,omitempty"`


	// Footer
	Footer *V2conversationmessagetypingeventforworkflowtopicconversationnotificationtemplatefooter `json:"footer,omitempty"`

}

func (o *V2conversationmessagetypingeventforworkflowtopicconversationcontentnotificationtemplate) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias V2conversationmessagetypingeventforworkflowtopicconversationcontentnotificationtemplate
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Language *string `json:"language,omitempty"`
		
		Header *V2conversationmessagetypingeventforworkflowtopicconversationnotificationtemplateheader `json:"header,omitempty"`
		
		Body *V2conversationmessagetypingeventforworkflowtopicconversationnotificationtemplatebody `json:"body,omitempty"`
		
		Footer *V2conversationmessagetypingeventforworkflowtopicconversationnotificationtemplatefooter `json:"footer,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Language: o.Language,
		
		Header: o.Header,
		
		Body: o.Body,
		
		Footer: o.Footer,
		Alias:    (*Alias)(o),
	})
}

func (o *V2conversationmessagetypingeventforworkflowtopicconversationcontentnotificationtemplate) UnmarshalJSON(b []byte) error {
	var V2conversationmessagetypingeventforworkflowtopicconversationcontentnotificationtemplateMap map[string]interface{}
	err := json.Unmarshal(b, &V2conversationmessagetypingeventforworkflowtopicconversationcontentnotificationtemplateMap)
	if err != nil {
		return err
	}
	
	if Id, ok := V2conversationmessagetypingeventforworkflowtopicconversationcontentnotificationtemplateMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Language, ok := V2conversationmessagetypingeventforworkflowtopicconversationcontentnotificationtemplateMap["language"].(string); ok {
		o.Language = &Language
	}
    
	if Header, ok := V2conversationmessagetypingeventforworkflowtopicconversationcontentnotificationtemplateMap["header"].(map[string]interface{}); ok {
		HeaderString, _ := json.Marshal(Header)
		json.Unmarshal(HeaderString, &o.Header)
	}
	
	if Body, ok := V2conversationmessagetypingeventforworkflowtopicconversationcontentnotificationtemplateMap["body"].(map[string]interface{}); ok {
		BodyString, _ := json.Marshal(Body)
		json.Unmarshal(BodyString, &o.Body)
	}
	
	if Footer, ok := V2conversationmessagetypingeventforworkflowtopicconversationcontentnotificationtemplateMap["footer"].(map[string]interface{}); ok {
		FooterString, _ := json.Marshal(Footer)
		json.Unmarshal(FooterString, &o.Footer)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *V2conversationmessagetypingeventforworkflowtopicconversationcontentnotificationtemplate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
