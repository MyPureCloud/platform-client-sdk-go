package platformclientv2
import (
	"encoding/json"
)

// Notificationtemplateheader - Template header object
type Notificationtemplateheader struct { 
	// VarType - Template header type
	VarType *string `json:"type,omitempty"`


	// Text - Header text. For WhatsApp, ignored
	Text *string `json:"text,omitempty"`


	// Media - Attachment object
	Media *Contentattachment `json:"media,omitempty"`


	// Parameters - Template parameters for placeholders in template
	Parameters *[]Notificationtemplateparameter `json:"parameters,omitempty"`

}

// String returns a JSON representation of the model
func (o *Notificationtemplateheader) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
