package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Survey
type Survey struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name
	Name *string `json:"name,omitempty"`

	// Conversation
	Conversation *Conversationreference `json:"conversation,omitempty"`

	// SurveyForm - Survey form used for this survey.
	SurveyForm *Surveyform `json:"surveyForm,omitempty"`

	// Agent
	Agent *Domainentityref `json:"agent,omitempty"`

	// Status
	Status *string `json:"status,omitempty"`

	// Queue
	Queue *Queuereference `json:"queue,omitempty"`

	// Answers
	Answers *Surveyscoringset `json:"answers,omitempty"`

	// CompletedDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	CompletedDate *time.Time `json:"completedDate,omitempty"`

	// SurveyErrorDetails - Additional information about what happened when the survey is in Error status.
	SurveyErrorDetails *Surveyerrordetails `json:"surveyErrorDetails,omitempty"`

	// AgentTeam - The team that the agent belongs to
	AgentTeam *Team `json:"agentTeam,omitempty"`

	// SurveyType - Type of the survey
	SurveyType *string `json:"surveyType,omitempty"`

	// MissingRequiredAnswer - True if any of the required questions for the survey form have not been answered. Null if survey is not finished.
	MissingRequiredAnswer *bool `json:"missingRequiredAnswer,omitempty"`

	// Flow - An Architect flow that executed in order to collect the answers for this survey.
	Flow *Addressableentityref `json:"flow,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Survey) SetField(field string, fieldValue interface{}) {
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

func (o Survey) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "CompletedDate", }
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
	type Alias Survey
	
	CompletedDate := new(string)
	if o.CompletedDate != nil {
		
		*CompletedDate = timeutil.Strftime(o.CompletedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		CompletedDate = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Conversation *Conversationreference `json:"conversation,omitempty"`
		
		SurveyForm *Surveyform `json:"surveyForm,omitempty"`
		
		Agent *Domainentityref `json:"agent,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		Queue *Queuereference `json:"queue,omitempty"`
		
		Answers *Surveyscoringset `json:"answers,omitempty"`
		
		CompletedDate *string `json:"completedDate,omitempty"`
		
		SurveyErrorDetails *Surveyerrordetails `json:"surveyErrorDetails,omitempty"`
		
		AgentTeam *Team `json:"agentTeam,omitempty"`
		
		SurveyType *string `json:"surveyType,omitempty"`
		
		MissingRequiredAnswer *bool `json:"missingRequiredAnswer,omitempty"`
		
		Flow *Addressableentityref `json:"flow,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Conversation: o.Conversation,
		
		SurveyForm: o.SurveyForm,
		
		Agent: o.Agent,
		
		Status: o.Status,
		
		Queue: o.Queue,
		
		Answers: o.Answers,
		
		CompletedDate: CompletedDate,
		
		SurveyErrorDetails: o.SurveyErrorDetails,
		
		AgentTeam: o.AgentTeam,
		
		SurveyType: o.SurveyType,
		
		MissingRequiredAnswer: o.MissingRequiredAnswer,
		
		Flow: o.Flow,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Survey) UnmarshalJSON(b []byte) error {
	var SurveyMap map[string]interface{}
	err := json.Unmarshal(b, &SurveyMap)
	if err != nil {
		return err
	}
	
	if Id, ok := SurveyMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := SurveyMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Conversation, ok := SurveyMap["conversation"].(map[string]interface{}); ok {
		ConversationString, _ := json.Marshal(Conversation)
		json.Unmarshal(ConversationString, &o.Conversation)
	}
	
	if SurveyForm, ok := SurveyMap["surveyForm"].(map[string]interface{}); ok {
		SurveyFormString, _ := json.Marshal(SurveyForm)
		json.Unmarshal(SurveyFormString, &o.SurveyForm)
	}
	
	if Agent, ok := SurveyMap["agent"].(map[string]interface{}); ok {
		AgentString, _ := json.Marshal(Agent)
		json.Unmarshal(AgentString, &o.Agent)
	}
	
	if Status, ok := SurveyMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if Queue, ok := SurveyMap["queue"].(map[string]interface{}); ok {
		QueueString, _ := json.Marshal(Queue)
		json.Unmarshal(QueueString, &o.Queue)
	}
	
	if Answers, ok := SurveyMap["answers"].(map[string]interface{}); ok {
		AnswersString, _ := json.Marshal(Answers)
		json.Unmarshal(AnswersString, &o.Answers)
	}
	
	if completedDateString, ok := SurveyMap["completedDate"].(string); ok {
		CompletedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", completedDateString)
		o.CompletedDate = &CompletedDate
	}
	
	if SurveyErrorDetails, ok := SurveyMap["surveyErrorDetails"].(map[string]interface{}); ok {
		SurveyErrorDetailsString, _ := json.Marshal(SurveyErrorDetails)
		json.Unmarshal(SurveyErrorDetailsString, &o.SurveyErrorDetails)
	}
	
	if AgentTeam, ok := SurveyMap["agentTeam"].(map[string]interface{}); ok {
		AgentTeamString, _ := json.Marshal(AgentTeam)
		json.Unmarshal(AgentTeamString, &o.AgentTeam)
	}
	
	if SurveyType, ok := SurveyMap["surveyType"].(string); ok {
		o.SurveyType = &SurveyType
	}
    
	if MissingRequiredAnswer, ok := SurveyMap["missingRequiredAnswer"].(bool); ok {
		o.MissingRequiredAnswer = &MissingRequiredAnswer
	}
    
	if Flow, ok := SurveyMap["flow"].(map[string]interface{}); ok {
		FlowString, _ := json.Marshal(Flow)
		json.Unmarshal(FlowString, &o.Flow)
	}
	
	if SelfUri, ok := SurveyMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Survey) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
