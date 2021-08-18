package platformclientv2
import (
	"github.com/leekchan/timeutil"
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

func (u *Conversationnotificationtemplateheader) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationnotificationtemplateheader

	

	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		Text *string `json:"text,omitempty"`
		
		Media *Conversationcontentattachment `json:"media,omitempty"`
		
		Parameters *[]Conversationnotificationtemplateparameter `json:"parameters,omitempty"`
		*Alias
	}{ 
		VarType: u.VarType,
		
		Text: u.Text,
		
		Media: u.Media,
		
		Parameters: u.Parameters,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Conversationnotificationtemplateheader) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
