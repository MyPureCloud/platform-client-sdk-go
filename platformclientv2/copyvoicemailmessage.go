package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Copyvoicemailmessage - Used to copy a VoicemailMessage to either a User or a Group
type Copyvoicemailmessage struct { 
	// VoicemailMessageId - The id of the VoicemailMessage to copy
	VoicemailMessageId *string `json:"voicemailMessageId,omitempty"`


	// UserId - The id of the User to copy the VoicemailMessage to
	UserId *string `json:"userId,omitempty"`


	// GroupId - The id of the Group to copy the VoicemailMessage to
	GroupId *string `json:"groupId,omitempty"`

}

func (u *Copyvoicemailmessage) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Copyvoicemailmessage

	

	return json.Marshal(&struct { 
		VoicemailMessageId *string `json:"voicemailMessageId,omitempty"`
		
		UserId *string `json:"userId,omitempty"`
		
		GroupId *string `json:"groupId,omitempty"`
		*Alias
	}{ 
		VoicemailMessageId: u.VoicemailMessageId,
		
		UserId: u.UserId,
		
		GroupId: u.GroupId,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Copyvoicemailmessage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
