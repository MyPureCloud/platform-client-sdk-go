package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Messagecontent - Message content element.
type Messagecontent struct { 
	// ContentType - Type of this content element. If contentType = \"Attachment\" only one item is allowed.
	ContentType *string `json:"contentType,omitempty"`


	// Location - Location content.
	Location *Contentlocation `json:"location,omitempty"`


	// Attachment - Attachment content.
	Attachment *Contentattachment `json:"attachment,omitempty"`


	// QuickReply - Quick reply content.
	QuickReply *Contentquickreply `json:"quickReply,omitempty"`


	// ButtonResponse - Button response content.
	ButtonResponse *Contentbuttonresponse `json:"buttonResponse,omitempty"`


	// Generic - Generic content.
	Generic *Contentgeneric `json:"generic,omitempty"`


	// List - List content.
	List *Contentlist `json:"list,omitempty"`


	// Template - Template notification content.
	Template *Contentnotificationtemplate `json:"template,omitempty"`


	// Reactions - A set of reactions to a message.
	Reactions *[]Contentreaction `json:"reactions,omitempty"`


	// Mention - Mention content.
	Mention *Messagingrecipient `json:"mention,omitempty"`


	// Postback - Structured message postback (Deprecated).
	Postback *Contentpostback `json:"postback,omitempty"`

}

func (u *Messagecontent) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Messagecontent

	

	return json.Marshal(&struct { 
		ContentType *string `json:"contentType,omitempty"`
		
		Location *Contentlocation `json:"location,omitempty"`
		
		Attachment *Contentattachment `json:"attachment,omitempty"`
		
		QuickReply *Contentquickreply `json:"quickReply,omitempty"`
		
		ButtonResponse *Contentbuttonresponse `json:"buttonResponse,omitempty"`
		
		Generic *Contentgeneric `json:"generic,omitempty"`
		
		List *Contentlist `json:"list,omitempty"`
		
		Template *Contentnotificationtemplate `json:"template,omitempty"`
		
		Reactions *[]Contentreaction `json:"reactions,omitempty"`
		
		Mention *Messagingrecipient `json:"mention,omitempty"`
		
		Postback *Contentpostback `json:"postback,omitempty"`
		*Alias
	}{ 
		ContentType: u.ContentType,
		
		Location: u.Location,
		
		Attachment: u.Attachment,
		
		QuickReply: u.QuickReply,
		
		ButtonResponse: u.ButtonResponse,
		
		Generic: u.Generic,
		
		List: u.List,
		
		Template: u.Template,
		
		Reactions: u.Reactions,
		
		Mention: u.Mention,
		
		Postback: u.Postback,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Messagecontent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
