package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// V2webmessagingundeliveredmessagetopicwebmessagingundeliveredmessageeventbody
type V2webmessagingundeliveredmessagetopicwebmessagingundeliveredmessageeventbody struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// ConversationId
	ConversationId *string `json:"conversationId,omitempty"`

	// DeploymentId
	DeploymentId *string `json:"deploymentId,omitempty"`

	// ParticipantId
	ParticipantId *string `json:"participantId,omitempty"`

	// ExternalContactId
	ExternalContactId *string `json:"externalContactId,omitempty"`

	// CommunicationId
	CommunicationId *string `json:"communicationId,omitempty"`

	// SessionExpiry
	SessionExpiry *int `json:"sessionExpiry,omitempty"`

	// Messages
	Messages *[]V2webmessagingundeliveredmessagetopicmessage `json:"messages,omitempty"`

	// EventTimeMs
	EventTimeMs *int `json:"eventTimeMs,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *V2webmessagingundeliveredmessagetopicwebmessagingundeliveredmessageeventbody) SetField(field string, fieldValue interface{}) {
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

func (o V2webmessagingundeliveredmessagetopicwebmessagingundeliveredmessageeventbody) MarshalJSON() ([]byte, error) {
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
	type Alias V2webmessagingundeliveredmessagetopicwebmessagingundeliveredmessageeventbody
	
	return json.Marshal(&struct { 
		ConversationId *string `json:"conversationId,omitempty"`
		
		DeploymentId *string `json:"deploymentId,omitempty"`
		
		ParticipantId *string `json:"participantId,omitempty"`
		
		ExternalContactId *string `json:"externalContactId,omitempty"`
		
		CommunicationId *string `json:"communicationId,omitempty"`
		
		SessionExpiry *int `json:"sessionExpiry,omitempty"`
		
		Messages *[]V2webmessagingundeliveredmessagetopicmessage `json:"messages,omitempty"`
		
		EventTimeMs *int `json:"eventTimeMs,omitempty"`
		Alias
	}{ 
		ConversationId: o.ConversationId,
		
		DeploymentId: o.DeploymentId,
		
		ParticipantId: o.ParticipantId,
		
		ExternalContactId: o.ExternalContactId,
		
		CommunicationId: o.CommunicationId,
		
		SessionExpiry: o.SessionExpiry,
		
		Messages: o.Messages,
		
		EventTimeMs: o.EventTimeMs,
		Alias:    (Alias)(o),
	})
}

func (o *V2webmessagingundeliveredmessagetopicwebmessagingundeliveredmessageeventbody) UnmarshalJSON(b []byte) error {
	var V2webmessagingundeliveredmessagetopicwebmessagingundeliveredmessageeventbodyMap map[string]interface{}
	err := json.Unmarshal(b, &V2webmessagingundeliveredmessagetopicwebmessagingundeliveredmessageeventbodyMap)
	if err != nil {
		return err
	}
	
	if ConversationId, ok := V2webmessagingundeliveredmessagetopicwebmessagingundeliveredmessageeventbodyMap["conversationId"].(string); ok {
		o.ConversationId = &ConversationId
	}
    
	if DeploymentId, ok := V2webmessagingundeliveredmessagetopicwebmessagingundeliveredmessageeventbodyMap["deploymentId"].(string); ok {
		o.DeploymentId = &DeploymentId
	}
    
	if ParticipantId, ok := V2webmessagingundeliveredmessagetopicwebmessagingundeliveredmessageeventbodyMap["participantId"].(string); ok {
		o.ParticipantId = &ParticipantId
	}
    
	if ExternalContactId, ok := V2webmessagingundeliveredmessagetopicwebmessagingundeliveredmessageeventbodyMap["externalContactId"].(string); ok {
		o.ExternalContactId = &ExternalContactId
	}
    
	if CommunicationId, ok := V2webmessagingundeliveredmessagetopicwebmessagingundeliveredmessageeventbodyMap["communicationId"].(string); ok {
		o.CommunicationId = &CommunicationId
	}
    
	if SessionExpiry, ok := V2webmessagingundeliveredmessagetopicwebmessagingundeliveredmessageeventbodyMap["sessionExpiry"].(float64); ok {
		SessionExpiryInt := int(SessionExpiry)
		o.SessionExpiry = &SessionExpiryInt
	}
	
	if Messages, ok := V2webmessagingundeliveredmessagetopicwebmessagingundeliveredmessageeventbodyMap["messages"].([]interface{}); ok {
		MessagesString, _ := json.Marshal(Messages)
		json.Unmarshal(MessagesString, &o.Messages)
	}
	
	if EventTimeMs, ok := V2webmessagingundeliveredmessagetopicwebmessagingundeliveredmessageeventbodyMap["eventTimeMs"].(float64); ok {
		EventTimeMsInt := int(EventTimeMs)
		o.EventTimeMs = &EventTimeMsInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *V2webmessagingundeliveredmessagetopicwebmessagingundeliveredmessageeventbody) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
