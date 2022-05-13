package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationcontentnotificationtemplate - Template notification object.
type Conversationcontentnotificationtemplate struct { 
	// Id - The messaging provider template ID. For WhatsApp, 'namespace@name'.
	Id *string `json:"id,omitempty"`


	// Language - Template language.
	Language *string `json:"language,omitempty"`


	// Header - The template header.
	Header *Conversationnotificationtemplateheader `json:"header,omitempty"`


	// Body - The template body.
	Body *Conversationnotificationtemplatebody `json:"body,omitempty"`


	// Footer - The template footer.
	Footer *Conversationnotificationtemplatefooter `json:"footer,omitempty"`

}

func (o *Conversationcontentnotificationtemplate) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationcontentnotificationtemplate
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Language *string `json:"language,omitempty"`
		
		Header *Conversationnotificationtemplateheader `json:"header,omitempty"`
		
		Body *Conversationnotificationtemplatebody `json:"body,omitempty"`
		
		Footer *Conversationnotificationtemplatefooter `json:"footer,omitempty"`
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

func (o *Conversationcontentnotificationtemplate) UnmarshalJSON(b []byte) error {
	var ConversationcontentnotificationtemplateMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationcontentnotificationtemplateMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ConversationcontentnotificationtemplateMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Language, ok := ConversationcontentnotificationtemplateMap["language"].(string); ok {
		o.Language = &Language
	}
    
	if Header, ok := ConversationcontentnotificationtemplateMap["header"].(map[string]interface{}); ok {
		HeaderString, _ := json.Marshal(Header)
		json.Unmarshal(HeaderString, &o.Header)
	}
	
	if Body, ok := ConversationcontentnotificationtemplateMap["body"].(map[string]interface{}); ok {
		BodyString, _ := json.Marshal(Body)
		json.Unmarshal(BodyString, &o.Body)
	}
	
	if Footer, ok := ConversationcontentnotificationtemplateMap["footer"].(map[string]interface{}); ok {
		FooterString, _ := json.Marshal(Footer)
		json.Unmarshal(FooterString, &o.Footer)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationcontentnotificationtemplate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
