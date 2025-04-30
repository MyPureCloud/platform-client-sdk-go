package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Knowledgeconversationcontext
type Knowledgeconversationcontext struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// ConversationId - The unique identifier of the conversation.
	ConversationId *string `json:"conversationId,omitempty"`

	// MediaType - The media type of the conversation.
	MediaType *string `json:"mediaType,omitempty"`

	// MessageType - The message type of the conversation.
	MessageType *string `json:"messageType,omitempty"`

	// QueueId - The unique identifier of the queue used to assign the interaction to the user.
	QueueId *string `json:"queueId,omitempty"`

	// ExternalContactId - The external contact identifier of the end-user participant.
	ExternalContactId *string `json:"externalContactId,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Knowledgeconversationcontext) SetField(field string, fieldValue interface{}) {
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

func (o Knowledgeconversationcontext) MarshalJSON() ([]byte, error) {
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
	type Alias Knowledgeconversationcontext
	
	return json.Marshal(&struct { 
		ConversationId *string `json:"conversationId,omitempty"`
		
		MediaType *string `json:"mediaType,omitempty"`
		
		MessageType *string `json:"messageType,omitempty"`
		
		QueueId *string `json:"queueId,omitempty"`
		
		ExternalContactId *string `json:"externalContactId,omitempty"`
		Alias
	}{ 
		ConversationId: o.ConversationId,
		
		MediaType: o.MediaType,
		
		MessageType: o.MessageType,
		
		QueueId: o.QueueId,
		
		ExternalContactId: o.ExternalContactId,
		Alias:    (Alias)(o),
	})
}

func (o *Knowledgeconversationcontext) UnmarshalJSON(b []byte) error {
	var KnowledgeconversationcontextMap map[string]interface{}
	err := json.Unmarshal(b, &KnowledgeconversationcontextMap)
	if err != nil {
		return err
	}
	
	if ConversationId, ok := KnowledgeconversationcontextMap["conversationId"].(string); ok {
		o.ConversationId = &ConversationId
	}
    
	if MediaType, ok := KnowledgeconversationcontextMap["mediaType"].(string); ok {
		o.MediaType = &MediaType
	}
    
	if MessageType, ok := KnowledgeconversationcontextMap["messageType"].(string); ok {
		o.MessageType = &MessageType
	}
    
	if QueueId, ok := KnowledgeconversationcontextMap["queueId"].(string); ok {
		o.QueueId = &QueueId
	}
    
	if ExternalContactId, ok := KnowledgeconversationcontextMap["externalContactId"].(string); ok {
		o.ExternalContactId = &ExternalContactId
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Knowledgeconversationcontext) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
