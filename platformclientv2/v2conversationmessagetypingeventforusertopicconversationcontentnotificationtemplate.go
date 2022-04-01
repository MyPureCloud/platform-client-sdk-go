package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// V2conversationmessagetypingeventforusertopicconversationcontentnotificationtemplate
type V2conversationmessagetypingeventforusertopicconversationcontentnotificationtemplate struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Language
	Language *string `json:"language,omitempty"`


	// Header
	Header *V2conversationmessagetypingeventforusertopicconversationnotificationtemplateheader `json:"header,omitempty"`


	// Body
	Body *V2conversationmessagetypingeventforusertopicconversationnotificationtemplatebody `json:"body,omitempty"`


	// Footer
	Footer *V2conversationmessagetypingeventforusertopicconversationnotificationtemplatefooter `json:"footer,omitempty"`

}

func (o *V2conversationmessagetypingeventforusertopicconversationcontentnotificationtemplate) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias V2conversationmessagetypingeventforusertopicconversationcontentnotificationtemplate
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Language *string `json:"language,omitempty"`
		
		Header *V2conversationmessagetypingeventforusertopicconversationnotificationtemplateheader `json:"header,omitempty"`
		
		Body *V2conversationmessagetypingeventforusertopicconversationnotificationtemplatebody `json:"body,omitempty"`
		
		Footer *V2conversationmessagetypingeventforusertopicconversationnotificationtemplatefooter `json:"footer,omitempty"`
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

func (o *V2conversationmessagetypingeventforusertopicconversationcontentnotificationtemplate) UnmarshalJSON(b []byte) error {
	var V2conversationmessagetypingeventforusertopicconversationcontentnotificationtemplateMap map[string]interface{}
	err := json.Unmarshal(b, &V2conversationmessagetypingeventforusertopicconversationcontentnotificationtemplateMap)
	if err != nil {
		return err
	}
	
	if Id, ok := V2conversationmessagetypingeventforusertopicconversationcontentnotificationtemplateMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Language, ok := V2conversationmessagetypingeventforusertopicconversationcontentnotificationtemplateMap["language"].(string); ok {
		o.Language = &Language
	}
	
	if Header, ok := V2conversationmessagetypingeventforusertopicconversationcontentnotificationtemplateMap["header"].(map[string]interface{}); ok {
		HeaderString, _ := json.Marshal(Header)
		json.Unmarshal(HeaderString, &o.Header)
	}
	
	if Body, ok := V2conversationmessagetypingeventforusertopicconversationcontentnotificationtemplateMap["body"].(map[string]interface{}); ok {
		BodyString, _ := json.Marshal(Body)
		json.Unmarshal(BodyString, &o.Body)
	}
	
	if Footer, ok := V2conversationmessagetypingeventforusertopicconversationcontentnotificationtemplateMap["footer"].(map[string]interface{}); ok {
		FooterString, _ := json.Marshal(Footer)
		json.Unmarshal(FooterString, &o.Footer)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *V2conversationmessagetypingeventforusertopicconversationcontentnotificationtemplate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
