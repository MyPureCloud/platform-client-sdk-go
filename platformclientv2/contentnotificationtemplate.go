package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Contentnotificationtemplate - Template notification object.
type Contentnotificationtemplate struct { 
	// Id - The messaging provider template ID. For WhatsApp, 'namespace@name'.
	Id *string `json:"id,omitempty"`


	// Language - Template language.
	Language *string `json:"language,omitempty"`


	// Header - The template header.
	Header *Notificationtemplateheader `json:"header,omitempty"`


	// Body - The template body.
	Body *Notificationtemplatebody `json:"body,omitempty"`


	// Footer - The template footer.
	Footer *Notificationtemplatefooter `json:"footer,omitempty"`

}

func (o *Contentnotificationtemplate) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Contentnotificationtemplate
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Language *string `json:"language,omitempty"`
		
		Header *Notificationtemplateheader `json:"header,omitempty"`
		
		Body *Notificationtemplatebody `json:"body,omitempty"`
		
		Footer *Notificationtemplatefooter `json:"footer,omitempty"`
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

func (o *Contentnotificationtemplate) UnmarshalJSON(b []byte) error {
	var ContentnotificationtemplateMap map[string]interface{}
	err := json.Unmarshal(b, &ContentnotificationtemplateMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ContentnotificationtemplateMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Language, ok := ContentnotificationtemplateMap["language"].(string); ok {
		o.Language = &Language
	}
	
	if Header, ok := ContentnotificationtemplateMap["header"].(map[string]interface{}); ok {
		HeaderString, _ := json.Marshal(Header)
		json.Unmarshal(HeaderString, &o.Header)
	}
	
	if Body, ok := ContentnotificationtemplateMap["body"].(map[string]interface{}); ok {
		BodyString, _ := json.Marshal(Body)
		json.Unmarshal(BodyString, &o.Body)
	}
	
	if Footer, ok := ContentnotificationtemplateMap["footer"].(map[string]interface{}); ok {
		FooterString, _ := json.Marshal(Footer)
		json.Unmarshal(FooterString, &o.Footer)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Contentnotificationtemplate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
