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

func (o *Recordingmessagingmessage) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Recordingmessagingmessage
	
	Timestamp := new(string)
	if o.Timestamp != nil {
		
		*Timestamp = timeutil.Strftime(o.Timestamp, "%Y-%m-%dT%H:%M:%S.%fZ")
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
		From: o.From,
		
		FromUser: o.FromUser,
		
		FromExternalContact: o.FromExternalContact,
		
		To: o.To,
		
		Timestamp: Timestamp,
		
		Id: o.Id,
		
		MessageText: o.MessageText,
		
		MessageMediaAttachments: o.MessageMediaAttachments,
		
		MessageStickerAttachments: o.MessageStickerAttachments,
		Alias:    (*Alias)(o),
	})
}

func (o *Recordingmessagingmessage) UnmarshalJSON(b []byte) error {
	var RecordingmessagingmessageMap map[string]interface{}
	err := json.Unmarshal(b, &RecordingmessagingmessageMap)
	if err != nil {
		return err
	}
	
	if From, ok := RecordingmessagingmessageMap["from"].(string); ok {
		o.From = &From
	}
	
	if FromUser, ok := RecordingmessagingmessageMap["fromUser"].(map[string]interface{}); ok {
		FromUserString, _ := json.Marshal(FromUser)
		json.Unmarshal(FromUserString, &o.FromUser)
	}
	
	if FromExternalContact, ok := RecordingmessagingmessageMap["fromExternalContact"].(map[string]interface{}); ok {
		FromExternalContactString, _ := json.Marshal(FromExternalContact)
		json.Unmarshal(FromExternalContactString, &o.FromExternalContact)
	}
	
	if To, ok := RecordingmessagingmessageMap["to"].(string); ok {
		o.To = &To
	}
	
	if timestampString, ok := RecordingmessagingmessageMap["timestamp"].(string); ok {
		Timestamp, _ := time.Parse("2006-01-02T15:04:05.999999Z", timestampString)
		o.Timestamp = &Timestamp
	}
	
	if Id, ok := RecordingmessagingmessageMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if MessageText, ok := RecordingmessagingmessageMap["messageText"].(string); ok {
		o.MessageText = &MessageText
	}
	
	if MessageMediaAttachments, ok := RecordingmessagingmessageMap["messageMediaAttachments"].([]interface{}); ok {
		MessageMediaAttachmentsString, _ := json.Marshal(MessageMediaAttachments)
		json.Unmarshal(MessageMediaAttachmentsString, &o.MessageMediaAttachments)
	}
	
	if MessageStickerAttachments, ok := RecordingmessagingmessageMap["messageStickerAttachments"].([]interface{}); ok {
		MessageStickerAttachmentsString, _ := json.Marshal(MessageStickerAttachments)
		json.Unmarshal(MessageStickerAttachmentsString, &o.MessageStickerAttachments)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Recordingmessagingmessage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
