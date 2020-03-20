package platformclientv2
import (
	"time"
	"encoding/json"
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


	// Timestamp - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
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

// String returns a JSON representation of the model
func (o *Recordingmessagingmessage) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
