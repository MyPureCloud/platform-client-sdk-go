package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationresponsesuggestionstopicsuggestioncontext
type Conversationresponsesuggestionstopicsuggestioncontext struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// QueueId
	QueueId *string `json:"queueId,omitempty"`

	// MediaType
	MediaType *string `json:"mediaType,omitempty"`

	// UserId
	UserId *string `json:"userId,omitempty"`

	// ExternalContactId
	ExternalContactId *string `json:"externalContactId,omitempty"`

	// AssistantId
	AssistantId *string `json:"assistantId,omitempty"`

	// UtteranceId
	UtteranceId *string `json:"utteranceId,omitempty"`

	// MessageId
	MessageId *string `json:"messageId,omitempty"`

	// QueryStatement
	QueryStatement *string `json:"queryStatement,omitempty"`

	// Language
	Language *string `json:"language,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Conversationresponsesuggestionstopicsuggestioncontext) SetField(field string, fieldValue interface{}) {
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

func (o Conversationresponsesuggestionstopicsuggestioncontext) MarshalJSON() ([]byte, error) {
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
	type Alias Conversationresponsesuggestionstopicsuggestioncontext
	
	return json.Marshal(&struct { 
		QueueId *string `json:"queueId,omitempty"`
		
		MediaType *string `json:"mediaType,omitempty"`
		
		UserId *string `json:"userId,omitempty"`
		
		ExternalContactId *string `json:"externalContactId,omitempty"`
		
		AssistantId *string `json:"assistantId,omitempty"`
		
		UtteranceId *string `json:"utteranceId,omitempty"`
		
		MessageId *string `json:"messageId,omitempty"`
		
		QueryStatement *string `json:"queryStatement,omitempty"`
		
		Language *string `json:"language,omitempty"`
		Alias
	}{ 
		QueueId: o.QueueId,
		
		MediaType: o.MediaType,
		
		UserId: o.UserId,
		
		ExternalContactId: o.ExternalContactId,
		
		AssistantId: o.AssistantId,
		
		UtteranceId: o.UtteranceId,
		
		MessageId: o.MessageId,
		
		QueryStatement: o.QueryStatement,
		
		Language: o.Language,
		Alias:    (Alias)(o),
	})
}

func (o *Conversationresponsesuggestionstopicsuggestioncontext) UnmarshalJSON(b []byte) error {
	var ConversationresponsesuggestionstopicsuggestioncontextMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationresponsesuggestionstopicsuggestioncontextMap)
	if err != nil {
		return err
	}
	
	if QueueId, ok := ConversationresponsesuggestionstopicsuggestioncontextMap["queueId"].(string); ok {
		o.QueueId = &QueueId
	}
    
	if MediaType, ok := ConversationresponsesuggestionstopicsuggestioncontextMap["mediaType"].(string); ok {
		o.MediaType = &MediaType
	}
    
	if UserId, ok := ConversationresponsesuggestionstopicsuggestioncontextMap["userId"].(string); ok {
		o.UserId = &UserId
	}
    
	if ExternalContactId, ok := ConversationresponsesuggestionstopicsuggestioncontextMap["externalContactId"].(string); ok {
		o.ExternalContactId = &ExternalContactId
	}
    
	if AssistantId, ok := ConversationresponsesuggestionstopicsuggestioncontextMap["assistantId"].(string); ok {
		o.AssistantId = &AssistantId
	}
    
	if UtteranceId, ok := ConversationresponsesuggestionstopicsuggestioncontextMap["utteranceId"].(string); ok {
		o.UtteranceId = &UtteranceId
	}
    
	if MessageId, ok := ConversationresponsesuggestionstopicsuggestioncontextMap["messageId"].(string); ok {
		o.MessageId = &MessageId
	}
    
	if QueryStatement, ok := ConversationresponsesuggestionstopicsuggestioncontextMap["queryStatement"].(string); ok {
		o.QueryStatement = &QueryStatement
	}
    
	if Language, ok := ConversationresponsesuggestionstopicsuggestioncontextMap["language"].(string); ok {
		o.Language = &Language
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationresponsesuggestionstopicsuggestioncontext) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
