package platformclientv2
import (
	"encoding/json"
)

// Surveyformandscoringset
type Surveyformandscoringset struct { 
	// SurveyForm
	SurveyForm *Surveyform `json:"surveyForm,omitempty"`


	// Answers
	Answers *Surveyscoringset `json:"answers,omitempty"`

}

// String returns a JSON representation of the model
func (o *Surveyformandscoringset) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
