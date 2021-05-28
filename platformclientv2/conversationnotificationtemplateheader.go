package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationnotificationtemplateheader - Template header object.
type Conversationnotificationtemplateheader struct { 
	// VarType - Template header type.
	VarType *string `json:"type,omitempty"`


	// Text - Header text. For WhatsApp, ignored.
	Text *string `json:"text,omitempty"`


	// Media - Media template header image.
	Media *Conversationcontentattachment `json:"media,omitempty"`


	// Parameters - Template parameters for placeholders in template.
	Parameters *[]Conversationnotificationtemplateparameter `json:"parameters,omitempty"`

}

// String returns a JSON representation of the model
func (o *Conversationnotificationtemplateheader) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
