package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Copyvoicemailmessage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
