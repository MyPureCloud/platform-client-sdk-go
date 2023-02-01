package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Chatmessage
type Chatmessage struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Body - The message body
	Body *string `json:"body,omitempty"`

	// Id
	Id *string `json:"id,omitempty"`

	// To - The message recipient
	To *string `json:"to,omitempty"`

	// From - The message sender
	From *string `json:"from,omitempty"`

	// Utc
	Utc *string `json:"utc,omitempty"`

	// Chat - The interaction id (if available)
	Chat *string `json:"chat,omitempty"`

	// Message - The message id
	Message *string `json:"message,omitempty"`

	// VarType
	VarType *string `json:"type,omitempty"`

	// BodyType - Type of the message body (v2 chats only)
	BodyType *string `json:"bodyType,omitempty"`

	// SenderCommunicationId - Communication of sender (v2 chats only)
	SenderCommunicationId *string `json:"senderCommunicationId,omitempty"`

	// ParticipantPurpose - Participant purpose of sender (v2 chats only)
	ParticipantPurpose *string `json:"participantPurpose,omitempty"`

	// User - The user information for the sender (if available)
	User *Chatmessageuser `json:"user,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Chatmessage) SetField(field string, fieldValue interface{}) {
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

func (o Chatmessage) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{  }
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
	type Alias Chatmessage
	
	return json.Marshal(&struct { 
		Body *string `json:"body,omitempty"`
		
		Id *string `json:"id,omitempty"`
		
		To *string `json:"to,omitempty"`
		
		From *string `json:"from,omitempty"`
		
		Utc *string `json:"utc,omitempty"`
		
		Chat *string `json:"chat,omitempty"`
		
		Message *string `json:"message,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		BodyType *string `json:"bodyType,omitempty"`
		
		SenderCommunicationId *string `json:"senderCommunicationId,omitempty"`
		
		ParticipantPurpose *string `json:"participantPurpose,omitempty"`
		
		User *Chatmessageuser `json:"user,omitempty"`
		Alias
	}{ 
		Body: o.Body,
		
		Id: o.Id,
		
		To: o.To,
		
		From: o.From,
		
		Utc: o.Utc,
		
		Chat: o.Chat,
		
		Message: o.Message,
		
		VarType: o.VarType,
		
		BodyType: o.BodyType,
		
		SenderCommunicationId: o.SenderCommunicationId,
		
		ParticipantPurpose: o.ParticipantPurpose,
		
		User: o.User,
		Alias:    (Alias)(o),
	})
}

func (o *Chatmessage) UnmarshalJSON(b []byte) error {
	var ChatmessageMap map[string]interface{}
	err := json.Unmarshal(b, &ChatmessageMap)
	if err != nil {
		return err
	}
	
	if Body, ok := ChatmessageMap["body"].(string); ok {
		o.Body = &Body
	}
    
	if Id, ok := ChatmessageMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if To, ok := ChatmessageMap["to"].(string); ok {
		o.To = &To
	}
    
	if From, ok := ChatmessageMap["from"].(string); ok {
		o.From = &From
	}
    
	if Utc, ok := ChatmessageMap["utc"].(string); ok {
		o.Utc = &Utc
	}
    
	if Chat, ok := ChatmessageMap["chat"].(string); ok {
		o.Chat = &Chat
	}
    
	if Message, ok := ChatmessageMap["message"].(string); ok {
		o.Message = &Message
	}
    
	if VarType, ok := ChatmessageMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if BodyType, ok := ChatmessageMap["bodyType"].(string); ok {
		o.BodyType = &BodyType
	}
    
	if SenderCommunicationId, ok := ChatmessageMap["senderCommunicationId"].(string); ok {
		o.SenderCommunicationId = &SenderCommunicationId
	}
    
	if ParticipantPurpose, ok := ChatmessageMap["participantPurpose"].(string); ok {
		o.ParticipantPurpose = &ParticipantPurpose
	}
    
	if User, ok := ChatmessageMap["user"].(map[string]interface{}); ok {
		UserString, _ := json.Marshal(User)
		json.Unmarshal(UserString, &o.User)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Chatmessage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
