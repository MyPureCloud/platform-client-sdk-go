package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Agentstateeventtopicagentstateeventnotification
type Agentstateeventtopicagentstateeventnotification struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// ConversationId
	ConversationId *string `json:"conversationId,omitempty"`

	// SessionId
	SessionId *string `json:"sessionId,omitempty"`

	// SessionStart
	SessionStart *int `json:"sessionStart,omitempty"`

	// CurrentStateStart
	CurrentStateStart *int `json:"currentStateStart,omitempty"`

	// CurrentState
	CurrentState *string `json:"currentState,omitempty"`

	// UserId
	UserId *string `json:"userId,omitempty"`

	// OriginatingDirection
	OriginatingDirection *string `json:"originatingDirection,omitempty"`

	// MediaType
	MediaType *string `json:"mediaType,omitempty"`

	// QueueId
	QueueId *string `json:"queueId,omitempty"`

	// RequestedLanguageId
	RequestedLanguageId *string `json:"requestedLanguageId,omitempty"`

	// RequestedSkillIds
	RequestedSkillIds *[]string `json:"requestedSkillIds,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Agentstateeventtopicagentstateeventnotification) SetField(field string, fieldValue interface{}) {
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

func (o Agentstateeventtopicagentstateeventnotification) MarshalJSON() ([]byte, error) {
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
	type Alias Agentstateeventtopicagentstateeventnotification
	
	return json.Marshal(&struct { 
		ConversationId *string `json:"conversationId,omitempty"`
		
		SessionId *string `json:"sessionId,omitempty"`
		
		SessionStart *int `json:"sessionStart,omitempty"`
		
		CurrentStateStart *int `json:"currentStateStart,omitempty"`
		
		CurrentState *string `json:"currentState,omitempty"`
		
		UserId *string `json:"userId,omitempty"`
		
		OriginatingDirection *string `json:"originatingDirection,omitempty"`
		
		MediaType *string `json:"mediaType,omitempty"`
		
		QueueId *string `json:"queueId,omitempty"`
		
		RequestedLanguageId *string `json:"requestedLanguageId,omitempty"`
		
		RequestedSkillIds *[]string `json:"requestedSkillIds,omitempty"`
		Alias
	}{ 
		ConversationId: o.ConversationId,
		
		SessionId: o.SessionId,
		
		SessionStart: o.SessionStart,
		
		CurrentStateStart: o.CurrentStateStart,
		
		CurrentState: o.CurrentState,
		
		UserId: o.UserId,
		
		OriginatingDirection: o.OriginatingDirection,
		
		MediaType: o.MediaType,
		
		QueueId: o.QueueId,
		
		RequestedLanguageId: o.RequestedLanguageId,
		
		RequestedSkillIds: o.RequestedSkillIds,
		Alias:    (Alias)(o),
	})
}

func (o *Agentstateeventtopicagentstateeventnotification) UnmarshalJSON(b []byte) error {
	var AgentstateeventtopicagentstateeventnotificationMap map[string]interface{}
	err := json.Unmarshal(b, &AgentstateeventtopicagentstateeventnotificationMap)
	if err != nil {
		return err
	}
	
	if ConversationId, ok := AgentstateeventtopicagentstateeventnotificationMap["conversationId"].(string); ok {
		o.ConversationId = &ConversationId
	}
    
	if SessionId, ok := AgentstateeventtopicagentstateeventnotificationMap["sessionId"].(string); ok {
		o.SessionId = &SessionId
	}
    
	if SessionStart, ok := AgentstateeventtopicagentstateeventnotificationMap["sessionStart"].(float64); ok {
		SessionStartInt := int(SessionStart)
		o.SessionStart = &SessionStartInt
	}
	
	if CurrentStateStart, ok := AgentstateeventtopicagentstateeventnotificationMap["currentStateStart"].(float64); ok {
		CurrentStateStartInt := int(CurrentStateStart)
		o.CurrentStateStart = &CurrentStateStartInt
	}
	
	if CurrentState, ok := AgentstateeventtopicagentstateeventnotificationMap["currentState"].(string); ok {
		o.CurrentState = &CurrentState
	}
    
	if UserId, ok := AgentstateeventtopicagentstateeventnotificationMap["userId"].(string); ok {
		o.UserId = &UserId
	}
    
	if OriginatingDirection, ok := AgentstateeventtopicagentstateeventnotificationMap["originatingDirection"].(string); ok {
		o.OriginatingDirection = &OriginatingDirection
	}
    
	if MediaType, ok := AgentstateeventtopicagentstateeventnotificationMap["mediaType"].(string); ok {
		o.MediaType = &MediaType
	}
    
	if QueueId, ok := AgentstateeventtopicagentstateeventnotificationMap["queueId"].(string); ok {
		o.QueueId = &QueueId
	}
    
	if RequestedLanguageId, ok := AgentstateeventtopicagentstateeventnotificationMap["requestedLanguageId"].(string); ok {
		o.RequestedLanguageId = &RequestedLanguageId
	}
    
	if RequestedSkillIds, ok := AgentstateeventtopicagentstateeventnotificationMap["requestedSkillIds"].([]interface{}); ok {
		RequestedSkillIdsString, _ := json.Marshal(RequestedSkillIds)
		json.Unmarshal(RequestedSkillIdsString, &o.RequestedSkillIds)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Agentstateeventtopicagentstateeventnotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
