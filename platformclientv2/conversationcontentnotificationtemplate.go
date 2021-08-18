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

func (u *Conversationcontentnotificationtemplate) MarshalJSON() ([]byte, error) {
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
		Id: u.Id,
		
		Language: u.Language,
		
		Header: u.Header,
		
		Body: u.Body,
		
		Footer: u.Footer,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Conversationcontentnotificationtemplate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
