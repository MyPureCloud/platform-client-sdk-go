package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationcontentform - Form content object.
type Conversationcontentform struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Introduction - The intro component, used to give an intro into what the form entails
	Introduction *Conversationcontentintroduction `json:"introduction,omitempty"`

	// FormPages - Form pages
	FormPages *[]Conversationformpage `json:"formPages,omitempty"`

	// ReceivedMessage - The message prompt to fill out the form received.
	ReceivedMessage *Conversationcontentreceivedreplymessage `json:"receivedMessage,omitempty"`

	// ReplyMessage - The reply message after the user has filled out the form received.
	ReplyMessage *Conversationcontentreceivedreplymessage `json:"replyMessage,omitempty"`

	// ShowSummary - Show summary at end of form submission.
	ShowSummary *bool `json:"showSummary,omitempty"`

	// Response - Content of the payload included in the Form response.
	Response *[]Conversationformresponsecomponent `json:"response,omitempty"`

	// OriginatingMessageId - Reference to the ID of the original message.
	OriginatingMessageId *string `json:"originatingMessageId,omitempty"`

	// CannedResponseId - The id of the canned response which was used to create the form.
	CannedResponseId *string `json:"cannedResponseId,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Conversationcontentform) SetField(field string, fieldValue interface{}) {
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

func (o Conversationcontentform) MarshalJSON() ([]byte, error) {
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
	type Alias Conversationcontentform
	
	return json.Marshal(&struct { 
		Introduction *Conversationcontentintroduction `json:"introduction,omitempty"`
		
		FormPages *[]Conversationformpage `json:"formPages,omitempty"`
		
		ReceivedMessage *Conversationcontentreceivedreplymessage `json:"receivedMessage,omitempty"`
		
		ReplyMessage *Conversationcontentreceivedreplymessage `json:"replyMessage,omitempty"`
		
		ShowSummary *bool `json:"showSummary,omitempty"`
		
		Response *[]Conversationformresponsecomponent `json:"response,omitempty"`
		
		OriginatingMessageId *string `json:"originatingMessageId,omitempty"`
		
		CannedResponseId *string `json:"cannedResponseId,omitempty"`
		Alias
	}{ 
		Introduction: o.Introduction,
		
		FormPages: o.FormPages,
		
		ReceivedMessage: o.ReceivedMessage,
		
		ReplyMessage: o.ReplyMessage,
		
		ShowSummary: o.ShowSummary,
		
		Response: o.Response,
		
		OriginatingMessageId: o.OriginatingMessageId,
		
		CannedResponseId: o.CannedResponseId,
		Alias:    (Alias)(o),
	})
}

func (o *Conversationcontentform) UnmarshalJSON(b []byte) error {
	var ConversationcontentformMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationcontentformMap)
	if err != nil {
		return err
	}
	
	if Introduction, ok := ConversationcontentformMap["introduction"].(map[string]interface{}); ok {
		IntroductionString, _ := json.Marshal(Introduction)
		json.Unmarshal(IntroductionString, &o.Introduction)
	}
	
	if FormPages, ok := ConversationcontentformMap["formPages"].([]interface{}); ok {
		FormPagesString, _ := json.Marshal(FormPages)
		json.Unmarshal(FormPagesString, &o.FormPages)
	}
	
	if ReceivedMessage, ok := ConversationcontentformMap["receivedMessage"].(map[string]interface{}); ok {
		ReceivedMessageString, _ := json.Marshal(ReceivedMessage)
		json.Unmarshal(ReceivedMessageString, &o.ReceivedMessage)
	}
	
	if ReplyMessage, ok := ConversationcontentformMap["replyMessage"].(map[string]interface{}); ok {
		ReplyMessageString, _ := json.Marshal(ReplyMessage)
		json.Unmarshal(ReplyMessageString, &o.ReplyMessage)
	}
	
	if ShowSummary, ok := ConversationcontentformMap["showSummary"].(bool); ok {
		o.ShowSummary = &ShowSummary
	}
    
	if Response, ok := ConversationcontentformMap["response"].([]interface{}); ok {
		ResponseString, _ := json.Marshal(Response)
		json.Unmarshal(ResponseString, &o.Response)
	}
	
	if OriginatingMessageId, ok := ConversationcontentformMap["originatingMessageId"].(string); ok {
		o.OriginatingMessageId = &OriginatingMessageId
	}
    
	if CannedResponseId, ok := ConversationcontentformMap["cannedResponseId"].(string); ok {
		o.CannedResponseId = &CannedResponseId
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationcontentform) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
