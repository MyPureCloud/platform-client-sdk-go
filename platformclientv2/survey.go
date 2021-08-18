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

func (u *Survey) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Survey

	
	CompletedDate := new(string)
	if u.CompletedDate != nil {
		
		*CompletedDate = timeutil.Strftime(u.CompletedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
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
		Id: u.Id,
		
		Name: u.Name,
		
		Conversation: u.Conversation,
		
		SurveyForm: u.SurveyForm,
		
		Agent: u.Agent,
		
		Status: u.Status,
		
		Queue: u.Queue,
		
		Answers: u.Answers,
		
		CompletedDate: CompletedDate,
		
		SurveyErrorDetails: u.SurveyErrorDetails,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Survey) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
