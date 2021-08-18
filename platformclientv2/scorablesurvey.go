package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Scorablesurvey
type Scorablesurvey struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// SurveyForm - Survey form used for this survey.
	SurveyForm *Surveyform `json:"surveyForm,omitempty"`


	// Status
	Status *string `json:"status,omitempty"`


	// Answers
	Answers *Surveyscoringset `json:"answers,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (u *Scorablesurvey) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Scorablesurvey

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		SurveyForm *Surveyform `json:"surveyForm,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		Answers *Surveyscoringset `json:"answers,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		SurveyForm: u.SurveyForm,
		
		Status: u.Status,
		
		Answers: u.Answers,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Scorablesurvey) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
