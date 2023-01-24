package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationassociation
type Conversationassociation struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// ExternalContactId - An external contact ID.  If not supplied, implies the conversation should be disassociated with any external contact.
	ExternalContactId *string `json:"externalContactId,omitempty"`

	// ConversationId - Conversation ID
	ConversationId *string `json:"conversationId,omitempty"`

	// CommunicationId - Communication ID
	CommunicationId *string `json:"communicationId,omitempty"`

	// MediaType - Media type
	MediaType *string `json:"mediaType,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Conversationassociation) SetField(field string, fieldValue interface{}) {
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

func (o Conversationassociation) MarshalJSON() ([]byte, error) {
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
			if contains(dateTimeFields, fieldName) {
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
	type Alias Conversationassociation
	
	return json.Marshal(&struct { 
		ExternalContactId *string `json:"externalContactId,omitempty"`
		
		ConversationId *string `json:"conversationId,omitempty"`
		
		CommunicationId *string `json:"communicationId,omitempty"`
		
		MediaType *string `json:"mediaType,omitempty"`
		Alias
	}{ 
		ExternalContactId: o.ExternalContactId,
		
		ConversationId: o.ConversationId,
		
		CommunicationId: o.CommunicationId,
		
		MediaType: o.MediaType,
		Alias:    (Alias)(o),
	})
}

func (o *Conversationassociation) UnmarshalJSON(b []byte) error {
	var ConversationassociationMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationassociationMap)
	if err != nil {
		return err
	}
	
	if ExternalContactId, ok := ConversationassociationMap["externalContactId"].(string); ok {
		o.ExternalContactId = &ExternalContactId
	}
    
	if ConversationId, ok := ConversationassociationMap["conversationId"].(string); ok {
		o.ConversationId = &ConversationId
	}
    
	if CommunicationId, ok := ConversationassociationMap["communicationId"].(string); ok {
		o.CommunicationId = &CommunicationId
	}
    
	if MediaType, ok := ConversationassociationMap["mediaType"].(string); ok {
		o.MediaType = &MediaType
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationassociation) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
