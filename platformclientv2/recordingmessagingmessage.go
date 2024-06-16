package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Recordingmessagingmessage
type Recordingmessagingmessage struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
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

	// Purpose - A well known string that specifies the purpose or type of the participant on this communication.
	Purpose *string `json:"purpose,omitempty"`

	// ParticipantId - A globally unique identifier for the participant on this communication.
	ParticipantId *string `json:"participantId,omitempty"`

	// Queue - A globally unique identifier for the queue involved in this communication.
	Queue *Addressableentityref `json:"queue,omitempty"`

	// Workflow - A globally unique identifier for the workflow involved in this communication.
	Workflow *Addressableentityref `json:"workflow,omitempty"`

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

	// Cards - List of cards offered for this message
	Cards *[]Card `json:"cards,omitempty"`

	// ContentType - Indicates the content type for this message
	ContentType *string `json:"contentType,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Recordingmessagingmessage) SetField(field string, fieldValue interface{}) {
	// Get Value object for field
	target := reflect.ValueOf(o)
	targetField := reflect.Indirect(target).FieldByName(field)

	// Set value
	if fieldValue != nil {
		targetField.Set(reflect.ValueOf(fieldValue))
	} else {
		// Must create a new Value (creates **type) then get its element (*type), which will be nil pointer of the appropriate type
		x := reflect.Indirect(reflect.New(targetField.Type()))
		targetField.Set(x)
	}

	// Add field to set field names list
	if o.SetFieldNames == nil {
		o.SetFieldNames = make(map[string]bool)
	}
	o.SetFieldNames[field] = true
}

func (o Recordingmessagingmessage) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "Timestamp", }
		localDateTimeFields := []string{  }
		dateFields := []string{  }

		// Construct object
		newObj := make(map[string]interface{})
		for fieldName := range o.SetFieldNames {
			// Get initial field value
			fieldValue := val.FieldByName(fieldName).Interface()

			// Apply value formatting overrides
			if fieldValue == nil || reflect.ValueOf(fieldValue).IsNil()  {
				// Do nothing. Just catching this case to avoid trying to custom serialize a nil value
			} else if contains(dateTimeFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%dT%H:%M:%S.%fZ")
			} else if contains(localDateTimeFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%dT%H:%M:%S.%f")
			} else if contains(dateFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%d")
			}

			// Assign value to field using JSON tag name
			newObj[getFieldName(reflect.TypeOf(&o), fieldName)] = fieldValue
		}

		// Marshal and return dynamically constructed interface
		return json.Marshal(newObj)
	}

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
		
		Purpose *string `json:"purpose,omitempty"`
		
		ParticipantId *string `json:"participantId,omitempty"`
		
		Queue *Addressableentityref `json:"queue,omitempty"`
		
		Workflow *Addressableentityref `json:"workflow,omitempty"`
		
		MessageText *string `json:"messageText,omitempty"`
		
		MessageMediaAttachments *[]Messagemediaattachment `json:"messageMediaAttachments,omitempty"`
		
		MessageStickerAttachments *[]Messagestickerattachment `json:"messageStickerAttachments,omitempty"`
		
		QuickReplies *[]Quickreply `json:"quickReplies,omitempty"`
		
		ButtonResponse *Buttonresponse `json:"buttonResponse,omitempty"`
		
		Story *Recordingcontentstory `json:"story,omitempty"`
		
		Cards *[]Card `json:"cards,omitempty"`
		
		ContentType *string `json:"contentType,omitempty"`
		Alias
	}{ 
		From: o.From,
		
		FromUser: o.FromUser,
		
		FromExternalContact: o.FromExternalContact,
		
		To: o.To,
		
		Timestamp: Timestamp,
		
		Id: o.Id,
		
		Purpose: o.Purpose,
		
		ParticipantId: o.ParticipantId,
		
		Queue: o.Queue,
		
		Workflow: o.Workflow,
		
		MessageText: o.MessageText,
		
		MessageMediaAttachments: o.MessageMediaAttachments,
		
		MessageStickerAttachments: o.MessageStickerAttachments,
		
		QuickReplies: o.QuickReplies,
		
		ButtonResponse: o.ButtonResponse,
		
		Story: o.Story,
		
		Cards: o.Cards,
		
		ContentType: o.ContentType,
		Alias:    (Alias)(o),
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
    
	if Purpose, ok := RecordingmessagingmessageMap["purpose"].(string); ok {
		o.Purpose = &Purpose
	}
    
	if ParticipantId, ok := RecordingmessagingmessageMap["participantId"].(string); ok {
		o.ParticipantId = &ParticipantId
	}
    
	if Queue, ok := RecordingmessagingmessageMap["queue"].(map[string]interface{}); ok {
		QueueString, _ := json.Marshal(Queue)
		json.Unmarshal(QueueString, &o.Queue)
	}
	
	if Workflow, ok := RecordingmessagingmessageMap["workflow"].(map[string]interface{}); ok {
		WorkflowString, _ := json.Marshal(Workflow)
		json.Unmarshal(WorkflowString, &o.Workflow)
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
	
	if Cards, ok := RecordingmessagingmessageMap["cards"].([]interface{}); ok {
		CardsString, _ := json.Marshal(Cards)
		json.Unmarshal(CardsString, &o.Cards)
	}
	
	if ContentType, ok := RecordingmessagingmessageMap["contentType"].(string); ok {
		o.ContentType = &ContentType
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Recordingmessagingmessage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
