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

	// ButtonResponses - List of Button Response selected by user for this message.
	ButtonResponses *[]Buttonresponse `json:"buttonResponses,omitempty"`

	// Story - Ephemeral story content.
	Story *Recordingcontentstory `json:"story,omitempty"`

	// Cards - List of cards offered for this message
	Cards *[]Card `json:"cards,omitempty"`

	// NotificationTemplate - Template notification content.
	NotificationTemplate *Recordingnotificationtemplate `json:"notificationTemplate,omitempty"`

	// DatePicker - DatePicker content object.
	DatePicker *Datepicker `json:"datePicker,omitempty"`

	// ListPicker - ListPicker content object.
	ListPicker *Listpicker `json:"listPicker,omitempty"`

	// ContentType - Indicates the content type for this message
	ContentType *string `json:"contentType,omitempty"`

	// SocialVisibility - For social media messages, the visibility of the message in the originating social platform
	SocialVisibility *string `json:"socialVisibility,omitempty"`

	// Events - List of event elements
	Events *[]Conversationmessageevent `json:"events,omitempty"`

	// InteractiveApplication - InteractiveApplication content.
	InteractiveApplication *Interactiveapplication `json:"interactiveApplication,omitempty"`

	// PaymentRequest - Payment request content.
	PaymentRequest *Paymentrequest `json:"paymentRequest,omitempty"`

	// PaymentResponse - Payment response content.
	PaymentResponse *Paymentresponse `json:"paymentResponse,omitempty"`

	// Form - Form content.
	Form *Recordingform `json:"form,omitempty"`

	// RoadsideAssistance - Roadside Assistance content.
	RoadsideAssistance *Recordingroadsideassistance `json:"roadsideAssistance,omitempty"`
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
		
		ButtonResponses *[]Buttonresponse `json:"buttonResponses,omitempty"`
		
		Story *Recordingcontentstory `json:"story,omitempty"`
		
		Cards *[]Card `json:"cards,omitempty"`
		
		NotificationTemplate *Recordingnotificationtemplate `json:"notificationTemplate,omitempty"`
		
		DatePicker *Datepicker `json:"datePicker,omitempty"`
		
		ListPicker *Listpicker `json:"listPicker,omitempty"`
		
		ContentType *string `json:"contentType,omitempty"`
		
		SocialVisibility *string `json:"socialVisibility,omitempty"`
		
		Events *[]Conversationmessageevent `json:"events,omitempty"`
		
		InteractiveApplication *Interactiveapplication `json:"interactiveApplication,omitempty"`
		
		PaymentRequest *Paymentrequest `json:"paymentRequest,omitempty"`
		
		PaymentResponse *Paymentresponse `json:"paymentResponse,omitempty"`
		
		Form *Recordingform `json:"form,omitempty"`
		
		RoadsideAssistance *Recordingroadsideassistance `json:"roadsideAssistance,omitempty"`
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
		
		ButtonResponses: o.ButtonResponses,
		
		Story: o.Story,
		
		Cards: o.Cards,
		
		NotificationTemplate: o.NotificationTemplate,
		
		DatePicker: o.DatePicker,
		
		ListPicker: o.ListPicker,
		
		ContentType: o.ContentType,
		
		SocialVisibility: o.SocialVisibility,
		
		Events: o.Events,
		
		InteractiveApplication: o.InteractiveApplication,
		
		PaymentRequest: o.PaymentRequest,
		
		PaymentResponse: o.PaymentResponse,
		
		Form: o.Form,
		
		RoadsideAssistance: o.RoadsideAssistance,
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
	
	if ButtonResponses, ok := RecordingmessagingmessageMap["buttonResponses"].([]interface{}); ok {
		ButtonResponsesString, _ := json.Marshal(ButtonResponses)
		json.Unmarshal(ButtonResponsesString, &o.ButtonResponses)
	}
	
	if Story, ok := RecordingmessagingmessageMap["story"].(map[string]interface{}); ok {
		StoryString, _ := json.Marshal(Story)
		json.Unmarshal(StoryString, &o.Story)
	}
	
	if Cards, ok := RecordingmessagingmessageMap["cards"].([]interface{}); ok {
		CardsString, _ := json.Marshal(Cards)
		json.Unmarshal(CardsString, &o.Cards)
	}
	
	if NotificationTemplate, ok := RecordingmessagingmessageMap["notificationTemplate"].(map[string]interface{}); ok {
		NotificationTemplateString, _ := json.Marshal(NotificationTemplate)
		json.Unmarshal(NotificationTemplateString, &o.NotificationTemplate)
	}
	
	if DatePicker, ok := RecordingmessagingmessageMap["datePicker"].(map[string]interface{}); ok {
		DatePickerString, _ := json.Marshal(DatePicker)
		json.Unmarshal(DatePickerString, &o.DatePicker)
	}
	
	if ListPicker, ok := RecordingmessagingmessageMap["listPicker"].(map[string]interface{}); ok {
		ListPickerString, _ := json.Marshal(ListPicker)
		json.Unmarshal(ListPickerString, &o.ListPicker)
	}
	
	if ContentType, ok := RecordingmessagingmessageMap["contentType"].(string); ok {
		o.ContentType = &ContentType
	}
    
	if SocialVisibility, ok := RecordingmessagingmessageMap["socialVisibility"].(string); ok {
		o.SocialVisibility = &SocialVisibility
	}
    
	if Events, ok := RecordingmessagingmessageMap["events"].([]interface{}); ok {
		EventsString, _ := json.Marshal(Events)
		json.Unmarshal(EventsString, &o.Events)
	}
	
	if InteractiveApplication, ok := RecordingmessagingmessageMap["interactiveApplication"].(map[string]interface{}); ok {
		InteractiveApplicationString, _ := json.Marshal(InteractiveApplication)
		json.Unmarshal(InteractiveApplicationString, &o.InteractiveApplication)
	}
	
	if PaymentRequest, ok := RecordingmessagingmessageMap["paymentRequest"].(map[string]interface{}); ok {
		PaymentRequestString, _ := json.Marshal(PaymentRequest)
		json.Unmarshal(PaymentRequestString, &o.PaymentRequest)
	}
	
	if PaymentResponse, ok := RecordingmessagingmessageMap["paymentResponse"].(map[string]interface{}); ok {
		PaymentResponseString, _ := json.Marshal(PaymentResponse)
		json.Unmarshal(PaymentResponseString, &o.PaymentResponse)
	}
	
	if Form, ok := RecordingmessagingmessageMap["form"].(map[string]interface{}); ok {
		FormString, _ := json.Marshal(Form)
		json.Unmarshal(FormString, &o.Form)
	}
	
	if RoadsideAssistance, ok := RecordingmessagingmessageMap["roadsideAssistance"].(map[string]interface{}); ok {
		RoadsideAssistanceString, _ := json.Marshal(RoadsideAssistance)
		json.Unmarshal(RoadsideAssistanceString, &o.RoadsideAssistance)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Recordingmessagingmessage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
