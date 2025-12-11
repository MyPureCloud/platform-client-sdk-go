package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Checklistinferencejobresponse
type Checklistinferencejobresponse struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - ID of the inference job.
	Id *string `json:"id,omitempty"`

	// Status - Status of the checklist inference job.
	Status *string `json:"status,omitempty"`

	// VarError - Error information associated with a job in case of any errors.
	VarError *Errorinfo `json:"error,omitempty"`

	// AgentChecklistInfo - Agent checklist info.
	AgentChecklistInfo *Agentchecklistinfo `json:"agentChecklistInfo,omitempty"`

	// JobStartTime - Date when the inference job started. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	JobStartTime *time.Time `json:"jobStartTime,omitempty"`

	// JobEndTime - Date when the inference job finished. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	JobEndTime *time.Time `json:"jobEndTime,omitempty"`

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

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Checklistinferencejobresponse) SetField(field string, fieldValue interface{}) {
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

func (o Checklistinferencejobresponse) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "JobStartTime","JobEndTime", }
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
	type Alias Checklistinferencejobresponse
	
	JobStartTime := new(string)
	if o.JobStartTime != nil {
		
		*JobStartTime = timeutil.Strftime(o.JobStartTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		JobStartTime = nil
	}
	
	JobEndTime := new(string)
	if o.JobEndTime != nil {
		
		*JobEndTime = timeutil.Strftime(o.JobEndTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		JobEndTime = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		VarError *Errorinfo `json:"error,omitempty"`
		
		AgentChecklistInfo *Agentchecklistinfo `json:"agentChecklistInfo,omitempty"`
		
		JobStartTime *string `json:"jobStartTime,omitempty"`
		
		JobEndTime *string `json:"jobEndTime,omitempty"`
		
		Language *string `json:"language,omitempty"`
		
		AgentId *string `json:"agentId,omitempty"`
		
		ParticipantId *string `json:"participantId,omitempty"`
		
		QueueId *string `json:"queueId,omitempty"`
		
		AssistantId *string `json:"assistantId,omitempty"`
		
		MediaType *string `json:"mediaType,omitempty"`
		
		Direction *string `json:"direction,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Status: o.Status,
		
		VarError: o.VarError,
		
		AgentChecklistInfo: o.AgentChecklistInfo,
		
		JobStartTime: JobStartTime,
		
		JobEndTime: JobEndTime,
		
		Language: o.Language,
		
		AgentId: o.AgentId,
		
		ParticipantId: o.ParticipantId,
		
		QueueId: o.QueueId,
		
		AssistantId: o.AssistantId,
		
		MediaType: o.MediaType,
		
		Direction: o.Direction,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Checklistinferencejobresponse) UnmarshalJSON(b []byte) error {
	var ChecklistinferencejobresponseMap map[string]interface{}
	err := json.Unmarshal(b, &ChecklistinferencejobresponseMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ChecklistinferencejobresponseMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Status, ok := ChecklistinferencejobresponseMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if VarError, ok := ChecklistinferencejobresponseMap["error"].(map[string]interface{}); ok {
		VarErrorString, _ := json.Marshal(VarError)
		json.Unmarshal(VarErrorString, &o.VarError)
	}
	
	if AgentChecklistInfo, ok := ChecklistinferencejobresponseMap["agentChecklistInfo"].(map[string]interface{}); ok {
		AgentChecklistInfoString, _ := json.Marshal(AgentChecklistInfo)
		json.Unmarshal(AgentChecklistInfoString, &o.AgentChecklistInfo)
	}
	
	if jobStartTimeString, ok := ChecklistinferencejobresponseMap["jobStartTime"].(string); ok {
		JobStartTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", jobStartTimeString)
		o.JobStartTime = &JobStartTime
	}
	
	if jobEndTimeString, ok := ChecklistinferencejobresponseMap["jobEndTime"].(string); ok {
		JobEndTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", jobEndTimeString)
		o.JobEndTime = &JobEndTime
	}
	
	if Language, ok := ChecklistinferencejobresponseMap["language"].(string); ok {
		o.Language = &Language
	}
    
	if AgentId, ok := ChecklistinferencejobresponseMap["agentId"].(string); ok {
		o.AgentId = &AgentId
	}
    
	if ParticipantId, ok := ChecklistinferencejobresponseMap["participantId"].(string); ok {
		o.ParticipantId = &ParticipantId
	}
    
	if QueueId, ok := ChecklistinferencejobresponseMap["queueId"].(string); ok {
		o.QueueId = &QueueId
	}
    
	if AssistantId, ok := ChecklistinferencejobresponseMap["assistantId"].(string); ok {
		o.AssistantId = &AssistantId
	}
    
	if MediaType, ok := ChecklistinferencejobresponseMap["mediaType"].(string); ok {
		o.MediaType = &MediaType
	}
    
	if Direction, ok := ChecklistinferencejobresponseMap["direction"].(string); ok {
		o.Direction = &Direction
	}
    
	if SelfUri, ok := ChecklistinferencejobresponseMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Checklistinferencejobresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
