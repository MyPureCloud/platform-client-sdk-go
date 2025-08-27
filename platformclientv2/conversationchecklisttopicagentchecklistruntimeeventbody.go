package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationchecklisttopicagentchecklistruntimeeventbody
type Conversationchecklisttopicagentchecklistruntimeeventbody struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// EvaluationStartDate
	EvaluationStartDate *time.Time `json:"evaluationStartDate,omitempty"`

	// EvaluationLastModifiedDate
	EvaluationLastModifiedDate *time.Time `json:"evaluationLastModifiedDate,omitempty"`

	// EvaluationFinalizedDate
	EvaluationFinalizedDate *time.Time `json:"evaluationFinalizedDate,omitempty"`

	// EvaluationFinalizedWithAcwDate
	EvaluationFinalizedWithAcwDate *time.Time `json:"evaluationFinalizedWithAcwDate,omitempty"`

	// ConversationId
	ConversationId *string `json:"conversationId,omitempty"`

	// CommunicationId
	CommunicationId *string `json:"communicationId,omitempty"`

	// AgentChecklistId
	AgentChecklistId *string `json:"agentChecklistId,omitempty"`

	// Name
	Name *string `json:"name,omitempty"`

	// Language
	Language *string `json:"language,omitempty"`

	// AgentId
	AgentId *string `json:"agentId,omitempty"`

	// ParticipantId
	ParticipantId *string `json:"participantId,omitempty"`

	// QueueId
	QueueId *string `json:"queueId,omitempty"`

	// AssistantId
	AssistantId *string `json:"assistantId,omitempty"`

	// MediaType
	MediaType *string `json:"mediaType,omitempty"`

	// Direction
	Direction *string `json:"direction,omitempty"`

	// Status
	Status *string `json:"status,omitempty"`

	// ExitReason
	ExitReason *string `json:"exitReason,omitempty"`

	// ActivationTriggers
	ActivationTriggers *[]Conversationchecklisttopicagentchecklistactivationtrigger `json:"activationTriggers,omitempty"`

	// ConversationContext
	ConversationContext *[]Conversationchecklisttopicconversationcontextturninfo `json:"conversationContext,omitempty"`

	// AgentChecklistItems
	AgentChecklistItems *[]Conversationchecklisttopicagentchecklistitemstate `json:"agentChecklistItems,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Conversationchecklisttopicagentchecklistruntimeeventbody) SetField(field string, fieldValue interface{}) {
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

func (o Conversationchecklisttopicagentchecklistruntimeeventbody) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "EvaluationStartDate","EvaluationLastModifiedDate","EvaluationFinalizedDate","EvaluationFinalizedWithAcwDate", }
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
	type Alias Conversationchecklisttopicagentchecklistruntimeeventbody
	
	EvaluationStartDate := new(string)
	if o.EvaluationStartDate != nil {
		
		*EvaluationStartDate = timeutil.Strftime(o.EvaluationStartDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		EvaluationStartDate = nil
	}
	
	EvaluationLastModifiedDate := new(string)
	if o.EvaluationLastModifiedDate != nil {
		
		*EvaluationLastModifiedDate = timeutil.Strftime(o.EvaluationLastModifiedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		EvaluationLastModifiedDate = nil
	}
	
	EvaluationFinalizedDate := new(string)
	if o.EvaluationFinalizedDate != nil {
		
		*EvaluationFinalizedDate = timeutil.Strftime(o.EvaluationFinalizedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		EvaluationFinalizedDate = nil
	}
	
	EvaluationFinalizedWithAcwDate := new(string)
	if o.EvaluationFinalizedWithAcwDate != nil {
		
		*EvaluationFinalizedWithAcwDate = timeutil.Strftime(o.EvaluationFinalizedWithAcwDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		EvaluationFinalizedWithAcwDate = nil
	}
	
	return json.Marshal(&struct { 
		EvaluationStartDate *string `json:"evaluationStartDate,omitempty"`
		
		EvaluationLastModifiedDate *string `json:"evaluationLastModifiedDate,omitempty"`
		
		EvaluationFinalizedDate *string `json:"evaluationFinalizedDate,omitempty"`
		
		EvaluationFinalizedWithAcwDate *string `json:"evaluationFinalizedWithAcwDate,omitempty"`
		
		ConversationId *string `json:"conversationId,omitempty"`
		
		CommunicationId *string `json:"communicationId,omitempty"`
		
		AgentChecklistId *string `json:"agentChecklistId,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Language *string `json:"language,omitempty"`
		
		AgentId *string `json:"agentId,omitempty"`
		
		ParticipantId *string `json:"participantId,omitempty"`
		
		QueueId *string `json:"queueId,omitempty"`
		
		AssistantId *string `json:"assistantId,omitempty"`
		
		MediaType *string `json:"mediaType,omitempty"`
		
		Direction *string `json:"direction,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		ExitReason *string `json:"exitReason,omitempty"`
		
		ActivationTriggers *[]Conversationchecklisttopicagentchecklistactivationtrigger `json:"activationTriggers,omitempty"`
		
		ConversationContext *[]Conversationchecklisttopicconversationcontextturninfo `json:"conversationContext,omitempty"`
		
		AgentChecklistItems *[]Conversationchecklisttopicagentchecklistitemstate `json:"agentChecklistItems,omitempty"`
		Alias
	}{ 
		EvaluationStartDate: EvaluationStartDate,
		
		EvaluationLastModifiedDate: EvaluationLastModifiedDate,
		
		EvaluationFinalizedDate: EvaluationFinalizedDate,
		
		EvaluationFinalizedWithAcwDate: EvaluationFinalizedWithAcwDate,
		
		ConversationId: o.ConversationId,
		
		CommunicationId: o.CommunicationId,
		
		AgentChecklistId: o.AgentChecklistId,
		
		Name: o.Name,
		
		Language: o.Language,
		
		AgentId: o.AgentId,
		
		ParticipantId: o.ParticipantId,
		
		QueueId: o.QueueId,
		
		AssistantId: o.AssistantId,
		
		MediaType: o.MediaType,
		
		Direction: o.Direction,
		
		Status: o.Status,
		
		ExitReason: o.ExitReason,
		
		ActivationTriggers: o.ActivationTriggers,
		
		ConversationContext: o.ConversationContext,
		
		AgentChecklistItems: o.AgentChecklistItems,
		Alias:    (Alias)(o),
	})
}

func (o *Conversationchecklisttopicagentchecklistruntimeeventbody) UnmarshalJSON(b []byte) error {
	var ConversationchecklisttopicagentchecklistruntimeeventbodyMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationchecklisttopicagentchecklistruntimeeventbodyMap)
	if err != nil {
		return err
	}
	
	if evaluationStartDateString, ok := ConversationchecklisttopicagentchecklistruntimeeventbodyMap["evaluationStartDate"].(string); ok {
		EvaluationStartDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", evaluationStartDateString)
		o.EvaluationStartDate = &EvaluationStartDate
	}
	
	if evaluationLastModifiedDateString, ok := ConversationchecklisttopicagentchecklistruntimeeventbodyMap["evaluationLastModifiedDate"].(string); ok {
		EvaluationLastModifiedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", evaluationLastModifiedDateString)
		o.EvaluationLastModifiedDate = &EvaluationLastModifiedDate
	}
	
	if evaluationFinalizedDateString, ok := ConversationchecklisttopicagentchecklistruntimeeventbodyMap["evaluationFinalizedDate"].(string); ok {
		EvaluationFinalizedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", evaluationFinalizedDateString)
		o.EvaluationFinalizedDate = &EvaluationFinalizedDate
	}
	
	if evaluationFinalizedWithAcwDateString, ok := ConversationchecklisttopicagentchecklistruntimeeventbodyMap["evaluationFinalizedWithAcwDate"].(string); ok {
		EvaluationFinalizedWithAcwDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", evaluationFinalizedWithAcwDateString)
		o.EvaluationFinalizedWithAcwDate = &EvaluationFinalizedWithAcwDate
	}
	
	if ConversationId, ok := ConversationchecklisttopicagentchecklistruntimeeventbodyMap["conversationId"].(string); ok {
		o.ConversationId = &ConversationId
	}
    
	if CommunicationId, ok := ConversationchecklisttopicagentchecklistruntimeeventbodyMap["communicationId"].(string); ok {
		o.CommunicationId = &CommunicationId
	}
    
	if AgentChecklistId, ok := ConversationchecklisttopicagentchecklistruntimeeventbodyMap["agentChecklistId"].(string); ok {
		o.AgentChecklistId = &AgentChecklistId
	}
    
	if Name, ok := ConversationchecklisttopicagentchecklistruntimeeventbodyMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Language, ok := ConversationchecklisttopicagentchecklistruntimeeventbodyMap["language"].(string); ok {
		o.Language = &Language
	}
    
	if AgentId, ok := ConversationchecklisttopicagentchecklistruntimeeventbodyMap["agentId"].(string); ok {
		o.AgentId = &AgentId
	}
    
	if ParticipantId, ok := ConversationchecklisttopicagentchecklistruntimeeventbodyMap["participantId"].(string); ok {
		o.ParticipantId = &ParticipantId
	}
    
	if QueueId, ok := ConversationchecklisttopicagentchecklistruntimeeventbodyMap["queueId"].(string); ok {
		o.QueueId = &QueueId
	}
    
	if AssistantId, ok := ConversationchecklisttopicagentchecklistruntimeeventbodyMap["assistantId"].(string); ok {
		o.AssistantId = &AssistantId
	}
    
	if MediaType, ok := ConversationchecklisttopicagentchecklistruntimeeventbodyMap["mediaType"].(string); ok {
		o.MediaType = &MediaType
	}
    
	if Direction, ok := ConversationchecklisttopicagentchecklistruntimeeventbodyMap["direction"].(string); ok {
		o.Direction = &Direction
	}
    
	if Status, ok := ConversationchecklisttopicagentchecklistruntimeeventbodyMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if ExitReason, ok := ConversationchecklisttopicagentchecklistruntimeeventbodyMap["exitReason"].(string); ok {
		o.ExitReason = &ExitReason
	}
    
	if ActivationTriggers, ok := ConversationchecklisttopicagentchecklistruntimeeventbodyMap["activationTriggers"].([]interface{}); ok {
		ActivationTriggersString, _ := json.Marshal(ActivationTriggers)
		json.Unmarshal(ActivationTriggersString, &o.ActivationTriggers)
	}
	
	if ConversationContext, ok := ConversationchecklisttopicagentchecklistruntimeeventbodyMap["conversationContext"].([]interface{}); ok {
		ConversationContextString, _ := json.Marshal(ConversationContext)
		json.Unmarshal(ConversationContextString, &o.ConversationContext)
	}
	
	if AgentChecklistItems, ok := ConversationchecklisttopicagentchecklistruntimeeventbodyMap["agentChecklistItems"].([]interface{}); ok {
		AgentChecklistItemsString, _ := json.Marshal(AgentChecklistItems)
		json.Unmarshal(AgentChecklistItemsString, &o.AgentChecklistItems)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationchecklisttopicagentchecklistruntimeeventbody) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
