package platformclientv2
import (
	"encoding/json"
)

// Contentnotificationtemplate - Template notification object
type Contentnotificationtemplate struct { 
	// Id - The messaging channel template id. For WhatsApp, 'namespace@name'
	Id *string `json:"id,omitempty"`


	// Language - Template language
	Language *string `json:"language,omitempty"`


	// Header - Template header object
	Header *Notificationtemplateheader `json:"header,omitempty"`


	// Body - Template body object
	Body *Notificationtemplatebody `json:"body,omitempty"`


	// Footer - Template footer object
	Footer *Notificationtemplatefooter `json:"footer,omitempty"`

}

// String returns a JSON representation of the model
func (o *Contentnotificationtemplate) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
