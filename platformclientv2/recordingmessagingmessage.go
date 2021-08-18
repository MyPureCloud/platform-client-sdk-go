package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Recordingmessagingmessage
type Recordingmessagingmessage struct { 
	// From
	From *string `json:"from,omitempty"`


	// FromUser
	FromUser *User `json:"fromUser,omitempty"`


	// FromExternalContact
	FromExternalContact *Externalcontact `json:"fromExternalContact,omitempty"`


	// To
	To *string `json:"to,omitempty"`


	// Timestamp - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	Timestamp *time.Time `json:"timestamp,omitempty"`


	// Id
	Id *string `json:"id,omitempty"`


	// MessageText
	MessageText *string `json:"messageText,omitempty"`


	// MessageMediaAttachments
	MessageMediaAttachments *[]Messagemediaattachment `json:"messageMediaAttachments,omitempty"`


	// MessageStickerAttachments
	MessageStickerAttachments *[]Messagestickerattachment `json:"messageStickerAttachments,omitempty"`

}

func (u *Recordingmessagingmessage) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Recordingmessagingmessage

	
	Timestamp := new(string)
	if u.Timestamp != nil {
		
		*Timestamp = timeutil.Strftime(u.Timestamp, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		Timestamp = nil
	}
	

	return json.Marshal(&struct { 
		From *string `json:"from,omitempty"`
		
		FromUser *User `json:"fromUser,omitempty"`
		
		FromExternalContact *Externalcontact `json:"fromExternalContact,omitempty"`
		
		To *string `json:"to,omitempty"`
		
		Timestamp *string `json:"timestamp,omitempty"`
		
		Id *string `json:"id,omitempty"`
		
		MessageText *string `json:"messageText,omitempty"`
		
		MessageMediaAttachments *[]Messagemediaattachment `json:"messageMediaAttachments,omitempty"`
		
		MessageStickerAttachments *[]Messagestickerattachment `json:"messageStickerAttachments,omitempty"`
		*Alias
	}{ 
		From: u.From,
		
		FromUser: u.FromUser,
		
		FromExternalContact: u.FromExternalContact,
		
		To: u.To,
		
		Timestamp: Timestamp,
		
		Id: u.Id,
		
		MessageText: u.MessageText,
		
		MessageMediaAttachments: u.MessageMediaAttachments,
		
		MessageStickerAttachments: u.MessageStickerAttachments,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Recordingmessagingmessage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
