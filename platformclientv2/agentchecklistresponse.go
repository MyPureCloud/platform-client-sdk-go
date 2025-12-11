package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Agentchecklistresponse
type Agentchecklistresponse struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - ID of the checklist.
	Id *string `json:"id,omitempty"`

	// Name - Name of the checklist.
	Name *string `json:"name,omitempty"`

	// ChecklistItems - List of individual Checklist Items.
	ChecklistItems *[]Checklistitem `json:"checklistItems,omitempty"`

	// ActivationTriggers - List of triggers that activated this checklist.
	ActivationTriggers *[]Activationtrigger `json:"activationTriggers,omitempty"`

	// Status - The evaluation status of the checklist.
	Status *string `json:"status,omitempty"`

	// ExitReason - Exit reason provided at the time of finalizing the checklist.
	ExitReason *string `json:"exitReason,omitempty"`

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

	// EvaluationStartDate - Date when the checklist evaluation began. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	EvaluationStartDate *time.Time `json:"evaluationStartDate,omitempty"`

	// EvaluationLastModifiedDate - Date when the checklist was last modified. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	EvaluationLastModifiedDate *time.Time `json:"evaluationLastModifiedDate,omitempty"`

	// EvaluationFinalizedDate - Date when the checklist was finalized. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	EvaluationFinalizedDate *time.Time `json:"evaluationFinalizedDate,omitempty"`

	// EvaluationFinalizedWithAcwDate - Date when the checklist was finalized with ACW. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	EvaluationFinalizedWithAcwDate *time.Time `json:"evaluationFinalizedWithAcwDate,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Agentchecklistresponse) SetField(field string, fieldValue interface{}) {
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

func (o Agentchecklistresponse) MarshalJSON() ([]byte, error) {
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
	type Alias Agentchecklistresponse
	
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
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		ChecklistItems *[]Checklistitem `json:"checklistItems,omitempty"`
		
		ActivationTriggers *[]Activationtrigger `json:"activationTriggers,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		ExitReason *string `json:"exitReason,omitempty"`
		
		Language *string `json:"language,omitempty"`
		
		AgentId *string `json:"agentId,omitempty"`
		
		ParticipantId *string `json:"participantId,omitempty"`
		
		QueueId *string `json:"queueId,omitempty"`
		
		AssistantId *string `json:"assistantId,omitempty"`
		
		MediaType *string `json:"mediaType,omitempty"`
		
		Direction *string `json:"direction,omitempty"`
		
		EvaluationStartDate *string `json:"evaluationStartDate,omitempty"`
		
		EvaluationLastModifiedDate *string `json:"evaluationLastModifiedDate,omitempty"`
		
		EvaluationFinalizedDate *string `json:"evaluationFinalizedDate,omitempty"`
		
		EvaluationFinalizedWithAcwDate *string `json:"evaluationFinalizedWithAcwDate,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		ChecklistItems: o.ChecklistItems,
		
		ActivationTriggers: o.ActivationTriggers,
		
		Status: o.Status,
		
		ExitReason: o.ExitReason,
		
		Language: o.Language,
		
		AgentId: o.AgentId,
		
		ParticipantId: o.ParticipantId,
		
		QueueId: o.QueueId,
		
		AssistantId: o.AssistantId,
		
		MediaType: o.MediaType,
		
		Direction: o.Direction,
		
		EvaluationStartDate: EvaluationStartDate,
		
		EvaluationLastModifiedDate: EvaluationLastModifiedDate,
		
		EvaluationFinalizedDate: EvaluationFinalizedDate,
		
		EvaluationFinalizedWithAcwDate: EvaluationFinalizedWithAcwDate,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Agentchecklistresponse) UnmarshalJSON(b []byte) error {
	var AgentchecklistresponseMap map[string]interface{}
	err := json.Unmarshal(b, &AgentchecklistresponseMap)
	if err != nil {
		return err
	}
	
	if Id, ok := AgentchecklistresponseMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := AgentchecklistresponseMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if ChecklistItems, ok := AgentchecklistresponseMap["checklistItems"].([]interface{}); ok {
		ChecklistItemsString, _ := json.Marshal(ChecklistItems)
		json.Unmarshal(ChecklistItemsString, &o.ChecklistItems)
	}
	
	if ActivationTriggers, ok := AgentchecklistresponseMap["activationTriggers"].([]interface{}); ok {
		ActivationTriggersString, _ := json.Marshal(ActivationTriggers)
		json.Unmarshal(ActivationTriggersString, &o.ActivationTriggers)
	}
	
	if Status, ok := AgentchecklistresponseMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if ExitReason, ok := AgentchecklistresponseMap["exitReason"].(string); ok {
		o.ExitReason = &ExitReason
	}
    
	if Language, ok := AgentchecklistresponseMap["language"].(string); ok {
		o.Language = &Language
	}
    
	if AgentId, ok := AgentchecklistresponseMap["agentId"].(string); ok {
		o.AgentId = &AgentId
	}
    
	if ParticipantId, ok := AgentchecklistresponseMap["participantId"].(string); ok {
		o.ParticipantId = &ParticipantId
	}
    
	if QueueId, ok := AgentchecklistresponseMap["queueId"].(string); ok {
		o.QueueId = &QueueId
	}
    
	if AssistantId, ok := AgentchecklistresponseMap["assistantId"].(string); ok {
		o.AssistantId = &AssistantId
	}
    
	if MediaType, ok := AgentchecklistresponseMap["mediaType"].(string); ok {
		o.MediaType = &MediaType
	}
    
	if Direction, ok := AgentchecklistresponseMap["direction"].(string); ok {
		o.Direction = &Direction
	}
    
	if evaluationStartDateString, ok := AgentchecklistresponseMap["evaluationStartDate"].(string); ok {
		EvaluationStartDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", evaluationStartDateString)
		o.EvaluationStartDate = &EvaluationStartDate
	}
	
	if evaluationLastModifiedDateString, ok := AgentchecklistresponseMap["evaluationLastModifiedDate"].(string); ok {
		EvaluationLastModifiedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", evaluationLastModifiedDateString)
		o.EvaluationLastModifiedDate = &EvaluationLastModifiedDate
	}
	
	if evaluationFinalizedDateString, ok := AgentchecklistresponseMap["evaluationFinalizedDate"].(string); ok {
		EvaluationFinalizedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", evaluationFinalizedDateString)
		o.EvaluationFinalizedDate = &EvaluationFinalizedDate
	}
	
	if evaluationFinalizedWithAcwDateString, ok := AgentchecklistresponseMap["evaluationFinalizedWithAcwDate"].(string); ok {
		EvaluationFinalizedWithAcwDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", evaluationFinalizedWithAcwDateString)
		o.EvaluationFinalizedWithAcwDate = &EvaluationFinalizedWithAcwDate
	}
	
	if SelfUri, ok := AgentchecklistresponseMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Agentchecklistresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
