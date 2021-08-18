package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationmessagecontent - Message content element.
type Conversationmessagecontent struct { 
	// ContentType - Type of this content element. If contentType = \"Attachment\" only one item is allowed.
	ContentType *string `json:"contentType,omitempty"`


	// Attachment - Attachment content.
	Attachment *Conversationcontentattachment `json:"attachment,omitempty"`


	// QuickReply - Quick reply content.
	QuickReply *Conversationcontentquickreply `json:"quickReply,omitempty"`


	// Template - Template notification content.
	Template *Conversationcontentnotificationtemplate `json:"template,omitempty"`


	// ButtonResponse - Button response content.
	ButtonResponse *Conversationcontentbuttonresponse `json:"buttonResponse,omitempty"`

}

func (u *Conversationmessagecontent) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationmessagecontent

	

	return json.Marshal(&struct { 
		ContentType *string `json:"contentType,omitempty"`
		
		Attachment *Conversationcontentattachment `json:"attachment,omitempty"`
		
		QuickReply *Conversationcontentquickreply `json:"quickReply,omitempty"`
		
		Template *Conversationcontentnotificationtemplate `json:"template,omitempty"`
		
		ButtonResponse *Conversationcontentbuttonresponse `json:"buttonResponse,omitempty"`
		*Alias
	}{ 
		ContentType: u.ContentType,
		
		Attachment: u.Attachment,
		
		QuickReply: u.QuickReply,
		
		Template: u.Template,
		
		ButtonResponse: u.ButtonResponse,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Conversationmessagecontent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
