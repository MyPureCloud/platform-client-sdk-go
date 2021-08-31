package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Survey
type Survey struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Conversation
	Conversation *Conversation `json:"conversation,omitempty"`


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


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Survey) MarshalJSON() ([]byte, error) {
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
		
		Conversation *Conversation `json:"conversation,omitempty"`
		
		SurveyForm *Surveyform `json:"surveyForm,omitempty"`
		
		Agent *Domainentityref `json:"agent,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		Queue *Queuereference `json:"queue,omitempty"`
		
		Answers *Surveyscoringset `json:"answers,omitempty"`
		
		CompletedDate *string `json:"completedDate,omitempty"`
		
		SurveyErrorDetails *Surveyerrordetails `json:"surveyErrorDetails,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
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
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
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
