package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Sendagentlessoutboundmessageresponse
type Sendagentlessoutboundmessageresponse struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// ConversationId - The identifier of the conversation.
	ConversationId *string `json:"conversationId,omitempty"`

	// FromAddress - The sender of the message.
	FromAddress *string `json:"fromAddress,omitempty"`

	// ToAddress - The recipient of the message.
	ToAddress *string `json:"toAddress,omitempty"`

	// MessengerType - Type of messenger.
	MessengerType *string `json:"messengerType,omitempty"`

	// TextBody - The body of the text message. (Deprecated - Instead use message.normalizedMessage.text)
	TextBody *string `json:"textBody,omitempty"`

	// MessagingTemplate - The messaging template sent. (Deprecated - Instead use message.normalizedMessage.content[#].template)
	MessagingTemplate *Sendmessagingtemplaterequest `json:"messagingTemplate,omitempty"`

	// UseExistingActiveConversation - Use an existing active conversation to send the agentless outbound message. Set this parameter to 'true' to use active conversation. Default value: false
	UseExistingActiveConversation *bool `json:"useExistingActiveConversation,omitempty"`

	// Message - Sent agentless outbound message in normalized format
	Message *Messagedata `json:"message,omitempty"`

	// Timestamp - The time when the message was sent. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	Timestamp *time.Time `json:"timestamp,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

	// User - Details of the user created the job
	User *Addressableentityref `json:"user,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Sendagentlessoutboundmessageresponse) SetField(field string, fieldValue interface{}) {
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

func (o Sendagentlessoutboundmessageresponse) MarshalJSON() ([]byte, error) {
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
	type Alias Sendagentlessoutboundmessageresponse
	
	Timestamp := new(string)
	if o.Timestamp != nil {
		
		*Timestamp = timeutil.Strftime(o.Timestamp, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		Timestamp = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		ConversationId *string `json:"conversationId,omitempty"`
		
		FromAddress *string `json:"fromAddress,omitempty"`
		
		ToAddress *string `json:"toAddress,omitempty"`
		
		MessengerType *string `json:"messengerType,omitempty"`
		
		TextBody *string `json:"textBody,omitempty"`
		
		MessagingTemplate *Sendmessagingtemplaterequest `json:"messagingTemplate,omitempty"`
		
		UseExistingActiveConversation *bool `json:"useExistingActiveConversation,omitempty"`
		
		Message *Messagedata `json:"message,omitempty"`
		
		Timestamp *string `json:"timestamp,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		
		User *Addressableentityref `json:"user,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		ConversationId: o.ConversationId,
		
		FromAddress: o.FromAddress,
		
		ToAddress: o.ToAddress,
		
		MessengerType: o.MessengerType,
		
		TextBody: o.TextBody,
		
		MessagingTemplate: o.MessagingTemplate,
		
		UseExistingActiveConversation: o.UseExistingActiveConversation,
		
		Message: o.Message,
		
		Timestamp: Timestamp,
		
		SelfUri: o.SelfUri,
		
		User: o.User,
		Alias:    (Alias)(o),
	})
}

func (o *Sendagentlessoutboundmessageresponse) UnmarshalJSON(b []byte) error {
	var SendagentlessoutboundmessageresponseMap map[string]interface{}
	err := json.Unmarshal(b, &SendagentlessoutboundmessageresponseMap)
	if err != nil {
		return err
	}
	
	if Id, ok := SendagentlessoutboundmessageresponseMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if ConversationId, ok := SendagentlessoutboundmessageresponseMap["conversationId"].(string); ok {
		o.ConversationId = &ConversationId
	}
    
	if FromAddress, ok := SendagentlessoutboundmessageresponseMap["fromAddress"].(string); ok {
		o.FromAddress = &FromAddress
	}
    
	if ToAddress, ok := SendagentlessoutboundmessageresponseMap["toAddress"].(string); ok {
		o.ToAddress = &ToAddress
	}
    
	if MessengerType, ok := SendagentlessoutboundmessageresponseMap["messengerType"].(string); ok {
		o.MessengerType = &MessengerType
	}
    
	if TextBody, ok := SendagentlessoutboundmessageresponseMap["textBody"].(string); ok {
		o.TextBody = &TextBody
	}
    
	if MessagingTemplate, ok := SendagentlessoutboundmessageresponseMap["messagingTemplate"].(map[string]interface{}); ok {
		MessagingTemplateString, _ := json.Marshal(MessagingTemplate)
		json.Unmarshal(MessagingTemplateString, &o.MessagingTemplate)
	}
	
	if UseExistingActiveConversation, ok := SendagentlessoutboundmessageresponseMap["useExistingActiveConversation"].(bool); ok {
		o.UseExistingActiveConversation = &UseExistingActiveConversation
	}
    
	if Message, ok := SendagentlessoutboundmessageresponseMap["message"].(map[string]interface{}); ok {
		MessageString, _ := json.Marshal(Message)
		json.Unmarshal(MessageString, &o.Message)
	}
	
	if timestampString, ok := SendagentlessoutboundmessageresponseMap["timestamp"].(string); ok {
		Timestamp, _ := time.Parse("2006-01-02T15:04:05.999999Z", timestampString)
		o.Timestamp = &Timestamp
	}
	
	if SelfUri, ok := SendagentlessoutboundmessageresponseMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    
	if User, ok := SendagentlessoutboundmessageresponseMap["user"].(map[string]interface{}); ok {
		UserString, _ := json.Marshal(User)
		json.Unmarshal(UserString, &o.User)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Sendagentlessoutboundmessageresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
