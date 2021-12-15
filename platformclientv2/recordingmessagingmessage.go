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
	// From - The message sender session id.
	From *string `json:"from,omitempty"`


	// FromUser - The user who sent this message.
	FromUser *User `json:"fromUser,omitempty"`


	// FromExternalContact - The PureCloud external contact sender details.
	FromExternalContact *Externalcontact `json:"fromExternalContact,omitempty"`


	// To - The message recipient.
	To *string `json:"to,omitempty"`


	// Timestamp - The time when the message was sent. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	Timestamp *time.Time `json:"timestamp,omitempty"`


	// Id - A globally unique identifier for this communication.
	Id *string `json:"id,omitempty"`


	// MessageText - The content of this message.
	MessageText *string `json:"messageText,omitempty"`


	// MessageMediaAttachments - List of media objects attached  with this message.
	MessageMediaAttachments *[]Messagemediaattachment `json:"messageMediaAttachments,omitempty"`


	// MessageStickerAttachments - List of message stickers attached with this message.
	MessageStickerAttachments *[]Messagestickerattachment `json:"messageStickerAttachments,omitempty"`


	// QuickReplies - List of quick reply options offered with this message.
	QuickReplies *[]Quickreply `json:"quickReplies,omitempty"`


	// ButtonResponse - Button Response selected by user for this message.
	ButtonResponse *Buttonresponse `json:"buttonResponse,omitempty"`


	// Story - Ephemeral story content.
	Story *Recordingcontentstory `json:"story,omitempty"`

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
		
		QuickReplies *[]Quickreply `json:"quickReplies,omitempty"`
		
		ButtonResponse *Buttonresponse `json:"buttonResponse,omitempty"`
		
		Story *Recordingcontentstory `json:"story,omitempty"`
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
		
		QuickReplies: o.QuickReplies,
		
		ButtonResponse: o.ButtonResponse,
		
		Story: o.Story,
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
	
	if QuickReplies, ok := RecordingmessagingmessageMap["quickReplies"].([]interface{}); ok {
		QuickRepliesString, _ := json.Marshal(QuickReplies)
		json.Unmarshal(QuickRepliesString, &o.QuickReplies)
	}
	
	if ButtonResponse, ok := RecordingmessagingmessageMap["buttonResponse"].(map[string]interface{}); ok {
		ButtonResponseString, _ := json.Marshal(ButtonResponse)
		json.Unmarshal(ButtonResponseString, &o.ButtonResponse)
	}
	
	if Story, ok := RecordingmessagingmessageMap["story"].(map[string]interface{}); ok {
		StoryString, _ := json.Marshal(Story)
		json.Unmarshal(StoryString, &o.Story)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Recordingmessagingmessage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
