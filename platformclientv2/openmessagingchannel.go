package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
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

func (u *Openmessagingchannel) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Openmessagingchannel

	
	Time := new(string)
	if u.Time != nil {
		
		*Time = timeutil.Strftime(u.Time, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		Time = nil
	}
	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Platform *string `json:"platform,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		MessageId *string `json:"messageId,omitempty"`
		
		To *Openmessagingtorecipient `json:"to,omitempty"`
		
		From *Openmessagingfromrecipient `json:"from,omitempty"`
		
		Time *string `json:"time,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Platform: u.Platform,
		
		VarType: u.VarType,
		
		MessageId: u.MessageId,
		
		To: u.To,
		
		From: u.From,
		
		Time: Time,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Openmessagingchannel) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
