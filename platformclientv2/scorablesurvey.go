package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Scorablesurvey) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
