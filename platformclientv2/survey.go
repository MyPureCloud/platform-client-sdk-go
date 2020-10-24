package platformclientv2
import (
	"time"
	"encoding/json"
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

// String returns a JSON representation of the model
func (o *Survey) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
