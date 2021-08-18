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

func (u *Contentnotificationtemplate) MarshalJSON() ([]byte, error) {
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
		Id: u.Id,
		
		Language: u.Language,
		
		Header: u.Header,
		
		Body: u.Body,
		
		Footer: u.Footer,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Contentnotificationtemplate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
