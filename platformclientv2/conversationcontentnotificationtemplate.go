package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Conversationcontentnotificationtemplate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
