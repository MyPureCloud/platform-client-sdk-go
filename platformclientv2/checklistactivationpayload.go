package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Checklistactivationpayload
type Checklistactivationpayload struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// ActivationTriggerType - Trigger type that activated this checklist.
	ActivationTriggerType *string `json:"activationTriggerType,omitempty"`

	// IntentId - The intent ID if checklist was triggered by an intent.
	IntentId *string `json:"intentId,omitempty"`

	// IntentName - The intent name if checklist was triggered by an intent.
	IntentName *string `json:"intentName,omitempty"`

	// Language - Language associated with the checklist.
	Language *string `json:"language,omitempty"`

	// AgentId - Agent ID.
	AgentId *string `json:"agentId,omitempty"`

	// ParticipantId - Participant ID.
	ParticipantId *string `json:"participantId,omitempty"`

	// QueueId - Queue ID.
	QueueId *string `json:"queueId,omitempty"`

	// AssistantId - Assistant ID.
	AssistantId *string `json:"assistantId,omitempty"`

	// MediaType - Media type.
	MediaType *string `json:"mediaType,omitempty"`

	// Direction - Direction of the conversation.
	Direction *string `json:"direction,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Checklistactivationpayload) SetField(field string, fieldValue interface{}) {
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

func (o Checklistactivationpayload) MarshalJSON() ([]byte, error) {
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
	type Alias Checklistactivationpayload
	
	return json.Marshal(&struct { 
		ActivationTriggerType *string `json:"activationTriggerType,omitempty"`
		
		IntentId *string `json:"intentId,omitempty"`
		
		IntentName *string `json:"intentName,omitempty"`
		
		Language *string `json:"language,omitempty"`
		
		AgentId *string `json:"agentId,omitempty"`
		
		ParticipantId *string `json:"participantId,omitempty"`
		
		QueueId *string `json:"queueId,omitempty"`
		
		AssistantId *string `json:"assistantId,omitempty"`
		
		MediaType *string `json:"mediaType,omitempty"`
		
		Direction *string `json:"direction,omitempty"`
		Alias
	}{ 
		ActivationTriggerType: o.ActivationTriggerType,
		
		IntentId: o.IntentId,
		
		IntentName: o.IntentName,
		
		Language: o.Language,
		
		AgentId: o.AgentId,
		
		ParticipantId: o.ParticipantId,
		
		QueueId: o.QueueId,
		
		AssistantId: o.AssistantId,
		
		MediaType: o.MediaType,
		
		Direction: o.Direction,
		Alias:    (Alias)(o),
	})
}

func (o *Checklistactivationpayload) UnmarshalJSON(b []byte) error {
	var ChecklistactivationpayloadMap map[string]interface{}
	err := json.Unmarshal(b, &ChecklistactivationpayloadMap)
	if err != nil {
		return err
	}
	
	if ActivationTriggerType, ok := ChecklistactivationpayloadMap["activationTriggerType"].(string); ok {
		o.ActivationTriggerType = &ActivationTriggerType
	}
    
	if IntentId, ok := ChecklistactivationpayloadMap["intentId"].(string); ok {
		o.IntentId = &IntentId
	}
    
	if IntentName, ok := ChecklistactivationpayloadMap["intentName"].(string); ok {
		o.IntentName = &IntentName
	}
    
	if Language, ok := ChecklistactivationpayloadMap["language"].(string); ok {
		o.Language = &Language
	}
    
	if AgentId, ok := ChecklistactivationpayloadMap["agentId"].(string); ok {
		o.AgentId = &AgentId
	}
    
	if ParticipantId, ok := ChecklistactivationpayloadMap["participantId"].(string); ok {
		o.ParticipantId = &ParticipantId
	}
    
	if QueueId, ok := ChecklistactivationpayloadMap["queueId"].(string); ok {
		o.QueueId = &QueueId
	}
    
	if AssistantId, ok := ChecklistactivationpayloadMap["assistantId"].(string); ok {
		o.AssistantId = &AssistantId
	}
    
	if MediaType, ok := ChecklistactivationpayloadMap["mediaType"].(string); ok {
		o.MediaType = &MediaType
	}
    
	if Direction, ok := ChecklistactivationpayloadMap["direction"].(string); ok {
		o.Direction = &Direction
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Checklistactivationpayload) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
