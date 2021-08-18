package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Surveyformandscoringset
type Surveyformandscoringset struct { 
	// SurveyForm
	SurveyForm *Surveyform `json:"surveyForm,omitempty"`


	// Answers
	Answers *Surveyscoringset `json:"answers,omitempty"`

}

func (u *Surveyformandscoringset) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Surveyformandscoringset

	

	return json.Marshal(&struct { 
		SurveyForm *Surveyform `json:"surveyForm,omitempty"`
		
		Answers *Surveyscoringset `json:"answers,omitempty"`
		*Alias
	}{ 
		SurveyForm: u.SurveyForm,
		
		Answers: u.Answers,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Surveyformandscoringset) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
