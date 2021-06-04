package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Openmessagingchannel - Channel-specific information that describes the message and the message channel/provider.
type Openmessagingchannel struct { 
	// Id - The Messaging Platform integration ID.
	Id *string `json:"id,omitempty"`


	// Platform - The provider type.
	Platform *string `json:"platform,omitempty"`


	// VarType - Specifies if this message is part of a private or public conversation.
	VarType *string `json:"type,omitempty"`


	// MessageId - Unique provider ID of the message such as a Facebook message ID.
	MessageId *string `json:"messageId,omitempty"`


	// To - Information about the recipient the message is sent to.
	To *Openmessagingtorecipient `json:"to,omitempty"`


	// From - Information about the recipient the message is received from.
	From *Openmessagingfromrecipient `json:"from,omitempty"`


	// Time - Original time of the event. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	Time *time.Time `json:"time,omitempty"`

}

// String returns a JSON representation of the model
func (o *Openmessagingchannel) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
